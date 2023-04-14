package generators

import (
	"fmt"
	"github.com/qdriven/go-for-qa/pkg/feeds"
	"io"
	"net"
	"net/url"
	"strings"
)

// Generator is interface of template generator
type Generator interface {
	Run(cfg *feeds.Config, writer io.Writer)
}

// favicon return favicon of url
func favicon(rawurl string) string {
	base := "https://f.start.me/%s"
	rawurl = strings.TrimSpace(rawurl)
	u, err := url.Parse(rawurl)
	if err != nil {
		return fmt.Sprintf(base, "o.oo")
	}
	host := u.Host
	if strings.Contains(host, ":") {
		host, _, _ = net.SplitHostPort(host)
	}
	return fmt.Sprintf(base, host)
}
