package gh

import (
	"fmt"
	"github.com/google/go-github/v48/github"
	"github.com/ogen-go/ogen/json"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func ReadTxtRepos(txtRepoPath string) []string {
	byteContents, err := ioutil.ReadFile(txtRepoPath)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(byteContents), "\n")
}

func SaveRanking(repos []*github.StarredRepository, topic string) {
	readme, err := os.OpenFile("README.md", os.O_RDWR|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer func(readme *os.File) {
		err := readme.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(readme)
	_, _ = readme.WriteString(head)
	for _, starredRepo := range repos {
		//if isDeprecated(repo.URL) {
		//	repo.Description = warning + repo.Description
		//}
		repo := starredRepo.Repository
		if repo != nil {
			var desc string
			if repo.Description != nil {
				desc = *repo.Description
			} else {
				desc = ""
			}
			_, err = readme.WriteString(fmt.Sprintf(
				"| [%s](%s) | %d | %d | %d | %s |\n", *repo.Name, *repo.URL,
				*repo.StargazersCount, *repo.ForksCount, *repo.OpenIssuesCount,
				desc))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	_, _ = readme.WriteString(fmt.Sprintf(tail, time.Now().Format(time.RFC3339)))
}

/**
Github URL Pattern: https://github.com/beego/beego
*/
func ResolveOwnerAndRepoName(gitUrl string) (ownerName, repoName string) {
	result := strings.Replace(gitUrl, GITHUB_URL, "", -1)
	parsedRepo := strings.Split(result, "/")
	return parsedRepo[0], parsedRepo[1]
}

func WriteStarredRepoToFile(repos []*github.StarredRepository) {
	b, _ := json.Marshal(repos)
	_ = os.WriteFile("starred_repo.json", b, fs.ModePerm)
}

func WriteFollowingUserToFile(repos []*github.User) {
	b, _ := json.Marshal(repos)
	_ = os.WriteFile("following_user.json", b, fs.ModePerm)
}
