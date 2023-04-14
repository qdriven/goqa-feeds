package feeds

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
)

type SimpleFeeds struct {
	Host string `mapstructure:"hostname"`
	Port int    `mapstructure:"port"`
	User string `mapstructure:"username"`
	Pass string `mapstructure:"password"`
}

type SiteFeed struct {
	Category    string
	Name        string
	Description string
	URL         string
	Icon        string // 网站图标
}

type CategoryFeeds struct {
	Category  string
	SiteFeeds []*SiteFeed
}

type Config struct {
	Title           string
	Description     string
	Template        string
	URL             string
	Logo            string
	Favicon         string
	Github          string
	Footer          string
	Content         *Content
	GoogleAnalytics string `yaml:"google_analytics"`
	WebStack        *WebStackConf
}

// Content is struct of categories
type Content struct {
	Categories []*Category
}

// Category struct
type Category struct {
	Name  string
	Sites []*Site
}

// Site struct
type Site struct {
	Name        string
	Description string
	URL         string
	Icon        string // 网站图标
}

// WebStackConf is config of webstack
type WebStackConf struct {
	Search *WebStackSearchConf
}

// WebStackSearchConf search engine config
type WebStackSearchConf struct {
	Enabled bool
	Default string
	Engines []string
}

// ParseConfig parse config from file
func ParseConfig(file string) (*Config, error) {
	buf, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	if err := yaml.Unmarshal(buf, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
