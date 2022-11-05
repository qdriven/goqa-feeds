package gh

import (
	"context"
	"github.com/google/go-github/v48/github"
	"io/ioutil"
	"log"
	"strings"
)

type QGithub struct {
	qClient     *github.Client
	userName    string
	accessToken string
}

func New() *QGithub {
	return &QGithub{
		qClient:  github.NewClient(nil),
		userName: "qdriven",
	}
}

func NewByUserNameAndToken(userName, accessToken string) *QGithub {
	return &QGithub{
		qClient:     github.NewClient(nil),
		userName:    userName,
		accessToken: accessToken,
	}
}

/**
get access token by access_token.txt
*/
func (q *QGithub) GetAccessToken(filePath string) string {
	tokenBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error occurs when getting access token")
	}
	return strings.TrimSpace(string(tokenBytes))
}

func (q *QGithub) GetRepoStats(owner, repoName string) (*github.Repository, *github.Response, error) {
	return q.qClient.Repositories.Get(context.Background(), owner, repoName)
}

func (q *QGithub) GetStarredRepos(page int) []*github.StarredRepository {
	listOpt := &github.ListOptions{
		Page:    page,
		PerPage: 100,
	}
	opts := &github.ActivityListStarredOptions{
		ListOptions: *listOpt,
	}
	result, _, _ := q.qClient.Activity.ListStarred(context.Background(), q.userName, opts)
	return result
}

func (q *QGithub) GetAllStarredRepos(from, end int) []*github.StarredRepository {
	var result []*github.StarredRepository
	for i := from; i < end; i++ {
		starred := q.GetStarredRepos(i)
		result = append(result, starred...)
		WriteStarredRepoToFile(result) // write as wait time
		//append(result, starred...)
		if len(starred) < 100 {
			break
		}
	}
	return result
}

func (q *QGithub) GetFollowing(page int) []*github.User {
	listOpt := &github.ListOptions{
		Page:    page,
		PerPage: 100,
	}
	result, _, _ := q.qClient.Users.ListFollowing(context.Background(), q.userName, listOpt)
	return result
}

func (q *QGithub) GetAllFollowing(from, end int) []*github.User {
	var result []*github.User
	for i := from; i < end; i++ {
		starred := q.GetFollowing(i)
		result = append(result, starred...)
		WriteFollowingUserToFile(result) // write as wait time
		//append(result, starred...)
		if len(starred) < 100 {
			break
		}
	}
	return result
}
