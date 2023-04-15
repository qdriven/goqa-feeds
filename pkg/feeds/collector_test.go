package feeds

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	"testing"
)

func TestGetGithubSite(t *testing.T) {
	c := CreateGithubStarredRepoCollector()
	_, categories := c.Collect()
	fmt.Println(categories)
	content := &Content{Categories: categories}
	data, _ := yaml.Marshal(content)
	ioutil.WriteFile("content.yaml", data, fs.ModePerm)
}
