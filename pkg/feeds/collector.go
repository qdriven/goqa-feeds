package feeds

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/qdriven/go-for-qa/pkg/gh"
	"github.com/thoas/go-funk"
	"os"
	"strconv"
	"strings"
)

type Collector interface {
	Collect() (map[string]*Category, []*Category)
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
func (g *GithubStarredRepoCollector) Collect() (map[string]*Category, []*Category) {
	categoryMapping, categories := g.setupCategories()
	//repos := g.GhClient.GetAllStarredRepositories(0, 60)
	var repos []*github.StarredRepository
	b, _ := os.ReadFile("./starred_repo.json")
	err := json.Unmarshal(b, &repos)
	if err != nil {
		panic(err)
	}
	funk.ForEach(repos, func(repo *github.StarredRepository) {

		site := &Site{
			Name:        *repo.Repository.Name,
			Description: *repo.Repository.Name,
			URL:         *repo.Repository.HTMLURL,
			Icon:        *repo.Repository.HTMLURL,
		}
		isOther := true
		funk.ForEach(filterKeywords, func(word string) {
			for _, topic := range repo.Repository.Topics {
				if funk.Contains(strings.ToUpper(topic), strings.ToUpper(word)) {
					categoryMapping[word].Sites = append(categoryMapping[word].Sites, site)
					isOther = false
				}
			}
		})
		if isOther {
			categoryMapping["others"].Sites = append(categoryMapping["others"].Sites, site)
		}
		starCount := repo.Repository.GetStargazersCount()
		category := getStartCountCategory(starCount, categoryMapping)
		category.Sites = append(category.Sites, site)
	})
	return categoryMapping, categories
}

func (g *GithubStarredRepoCollector) setupCategories() (map[string]*Category, []*Category) {
	categoryMapping := map[string]*Category{}
	categories := make([]*Category, 0)
	funk.ForEach(filterKeywords, func(word string) {
		categoryMapping[word] = &Category{
			Name:  word,
			Sites: make([]*Site, 0),
		}
	})
	funk.ForEach(starredFilter, func(star int) {
		categoryMapping[strconv.Itoa(star)] = &Category{
			Name:  "Github Stars:" + strconv.Itoa(star),
			Sites: make([]*Site, 0),
		}
	})
	categoryMapping["others"] = &Category{
		Name:  "其他",
		Sites: make([]*Site, 0),
	}
	for s, category := range categoryMapping {
		fmt.Println(`add {s} to list`, s)
		categories = append(categories, category)
	}
	return categoryMapping, categories
}

func getStartCountCategory(starCount int, cateSite map[string]*Category) *Category {
	sorted := starredFilter
	for i := 0; i < len(sorted); i++ {
		if starCount > sorted[i] {
			return cateSite[strconv.Itoa(sorted[i])]
		}
	}
	return cateSite[strconv.Itoa(sorted[len(sorted)-1])]
}
