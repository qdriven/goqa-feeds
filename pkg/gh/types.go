package gh

import "time"

const (
	GITHUB_URL     = "https://github.com/"
	GITHUB_API_URL = "https://api.github.com/"
	head           = `# My Starred List
A list of popular github projects related to My Starred List (ranked by stars automatically)
Please update **<topics>.txt** (via Pull Request)

| Project Name | Stars | Forks | Open Issues | Description | Last Commit |
| ------------ | ----- | ----- | ----------- | ----------- | ----------- |
`
	tail = "\n*Last Automatic Update: %v*"

	warning = "⚠️ No longer maintained ⚠️  "
)

type Repo struct {
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	DefaultBranch  string    `json:"default_branch"`
	Stars          int       `json:"stargazers_count"`
	Forks          int       `json:"forks_count"`
	Issues         int       `json:"open_issues_count"`
	Created        time.Time `json:"created_at"`
	Updated        time.Time `json:"updated_at"`
	URL            string    `json:"html_url"`
	LastCommitDate time.Time `json:"-"`
}

// HeadCommit describes a head commit of default branch
type HeadCommit struct {
	Sha    string `json:"sha"`
	Commit struct {
		Committer struct {
			Name  string    `json:"name"`
			Email string    `json:"email"`
			Date  time.Time `json:"date"`
		} `json:"committer"`
	} `json:"commit"`
}
