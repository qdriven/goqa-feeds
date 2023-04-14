package feeds

import (
	"github.com/google/go-github/v48/github"
	"github.com/qdriven/go-for-qa/pkg/gh"
	funk "github.com/thoas/go-funk"
)

type FeedsCollector interface {
	Collect()
}

type GithubStarredRepoCollector struct {
	GhClient *gh.QGithub
}

func CreateGithubStarredRepoCollort() *GithubStarredRepoCollector {
	return &GithubStarredRepoCollector{
		GhClient: gh.NewGithubClient(),
	}
}
func (g *GithubStarredRepoCollector) Collect() {
	repos := g.GhClient.GetAllStarredRepositories(1, 10000)
	//convert to yaml file
	//keywords: chatgpt/ai/automation/qa
	filter_keywords := []string{"chatgpt", "ai", "automation", "low-code", "no-code", "awesome",
		"framework"}
	starred_filter := []string{"100", "50", "500", "1000"}
	cateSites := map[string]*Category{}

	funk.ForEach(filter_keywords, func(word string) {
		cateSites[word] = &Category{
			Name:  word,
			Sites: []*Site{},
		}
	})
	funk.ForEach(starred_filter, func(word string) {
		cateSites[word] = &Category{
			Name:  "Github Stars:" + word,
			Sites: []*Site{},
		}
	})
	cateSites["others"] = &Category{
		Name:  "其他",
		Sites: []*Site{},
	}

	funk.ForEach(repos, func(repo *github.StarredRepository) {
		site := &Site{
			Name:        *repo.Repository.Name,
			Description: *repo.Repository.Description,
			URL:         *repo.Repository.HTMLURL,
			Icon:        *repo.Repository.HTMLURL,
		}
		funk.ForEach(filter_keywords, func(word string) {
			if funk.Contains(word, repo.Repository.Topics) {
				cateSites[word].Sites = append(cateSites[word].Sites, site)
			} else {
				cateSites["others"].Sites = append(cateSites[word].Sites, site)
			}
		})
		//TODO: add to stars
		funk.ForEach(starred_filter, func(stars string) {
			repo.Repository.GetStargazersCount()
		})
	})

}
