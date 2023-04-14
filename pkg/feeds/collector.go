package feeds

import (
	"github.com/google/go-github/v48/github"
	"github.com/qdriven/go-for-qa/pkg/gh"
	funk "github.com/thoas/go-funk"
	"sort"
	"strconv"
)

type Collector interface {
	Collect() map[string]*Category
}

type GithubStarredRepoCollector struct {
	GhClient *gh.QGithub
}

var (
	filterKeywords = []string{"chatgpt", "ai", "automation", "low-code", "no-code", "awesome",
		"framework"}
	starredFilter = []int{1000, 500, 100, 50}
)

func CreateGithubStarredRepoCollector() *GithubStarredRepoCollector {
	return &GithubStarredRepoCollector{
		GhClient: gh.NewGithubClient(),
	}
}
func (g *GithubStarredRepoCollector) Collect() map[string]*Category {
	cateSites := map[string]*Category{}
	funk.ForEach(filterKeywords, func(word string) {
		cateSites[word] = &Category{
			Name:  word,
			Sites: []*Site{},
		}
	})
	funk.ForEach(starredFilter, func(star int) {
		cateSites[strconv.Itoa(star)] = &Category{
			Name:  "Github Stars:" + strconv.Itoa(star),
			Sites: []*Site{},
		}
	})
	cateSites["others"] = &Category{
		Name:  "其他",
		Sites: []*Site{},
	}
	repos := g.GhClient.GetAllStarredRepositories(0, 60)
	funk.ForEach(repos, func(repo *github.StarredRepository) {
		site := &Site{
			Name:        *repo.Repository.Name,
			Description: *repo.Repository.Description,
			URL:         *repo.Repository.HTMLURL,
			Icon:        *repo.Repository.HTMLURL,
		}
		funk.ForEach(filterKeywords, func(word string) {
			if funk.Contains(word, repo.Repository.Topics) {
				cateSites[word].Sites = append(cateSites[word].Sites, site)
			} else {
				cateSites["others"].Sites = append(cateSites[word].Sites, site)
			}
		})
		starCount := repo.Repository.GetStargazersCount()
		category := getStartCountCategory(starCount, cateSites)
		category.Sites = append(category.Sites, site)
	})
	return cateSites
}

func getStartCountCategory(starCount int, cateSite map[string]*Category) *Category {
	sort.Ints(starredFilter)
	for i := 0; i < len(starredFilter); i++ {
		if starCount > starredFilter[i] {
			return cateSite[strconv.Itoa(starCount)]
		}
	}
	return cateSite[strconv.Itoa(starredFilter[len(starredFilter)-1])]
}
