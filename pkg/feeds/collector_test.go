package feeds

import (
	"fmt"
	"testing"
)

func TestGetGithubSite(t *testing.T) {
	c := CreateGithubStarredRepoCollector()
	cates := c.Collect()
	fmt.Println(cates)
}
