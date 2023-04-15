package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/qdriven/go-for-qa/pkg/gh"
	"os"
	"sort"
)

func FetchAndSaveStarredRepo() {
	//starredLists := []string{"starred_repo.json", "starred_repo_15_40.json", "starred_repo_41_50.json"}
	starredLists := []string{"starred_repo.json"}
	client := gh.NewGithubClient()
	//write starred repo to starred_repo.json
	client.GetAllStarredRepos(0, 60)
	var allRepos []*github.StarredRepository
	for i, list := range starredLists {
		fmt.Println(i)
		var repos []*github.StarredRepository
		b, _ := os.ReadFile(list)
		err := json.Unmarshal(b, &repos)
		if err != nil {
			return
		}
		allRepos = append(allRepos, repos...)
	}
	sort.SliceStable(allRepos, func(i, j int) bool {
		return *allRepos[i].Repository.StargazersCount >
			*allRepos[i].Repository.StargazersCount
	})
	gh.SaveRanking(allRepos, "my-starred")
	client.GetAllFollowing(0, 20)
}
