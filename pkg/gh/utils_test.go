package gh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadTxtRepo(t *testing.T) {
	repos := ReadTxtRepos("go-list.md")
	assert.True(t, len(repos) > 1, "read repos successfully")
}

func TestResolveOwnerNameAndRepoName(t *testing.T) {
	owner, repo := ResolveOwnerAndRepoName("https://github.com/gobuffalo/buffalo")
	assert.Equal(t, "gobuffalo", owner, "parsed owner name")
	assert.Equal(t, "buffalo", repo, "parsed repo name")
}

func TestSaveRanks(t *testing.T) {

}
