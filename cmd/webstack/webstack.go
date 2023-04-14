package main

import (
	"flag"
	feeds "github.com/qdriven/go-for-qa/pkg/feeds"
	"github.com/qdriven/go-for-qa/pkg/generator"
	"log"
	"os"
)

func main() {
	cfgFile := flag.String("c", "config.yml", "Config file")

	flag.Parse()

	cfg, err := feeds.ParseConfig(*cfgFile)
	if err != nil {
		log.Fatal("Parse Config file error: ", err.Error())
	}

	var generator generators.Generator
	generator = &generators.WebStackGenerator{}
	generator.Run(cfg, os.Stdout)
}
