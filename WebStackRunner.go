package main

import (
	"github.com/qdriven/go-for-qa/pkg"
	"github.com/qdriven/go-for-qa/pkg/feeds"
	generators "github.com/qdriven/go-for-qa/pkg/generator"
	"log"
	"os"
)

func main() {
	pkg.FetchAndSaveStarredRepo()
	_, categories := feeds.CreateGithubStarredRepoCollector().Collect()
	cfg, err := feeds.ParseConfig("config.yaml")
	if err != nil {
		log.Fatal("Parse Config file error: ", err.Error())
	}
	cfg.Content = &feeds.Content{Categories: categories}
	var generator generators.Generator
	generator = &generators.WebStackGenerator{}
	f, _ := os.Create("index.html")
	generator.Run(cfg, f)
}
