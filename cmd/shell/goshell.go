package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"qfluent-go/pkg/gosh"
	"qfluent-go/pkg/gosh/api"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ctx = context.WithValue(ctx, "gosh.prompt", api.DefaultPrompt)
	ctx = context.WithValue(ctx, "gosh.stdout", os.Stdout)
	ctx = context.WithValue(ctx, "gosh.stderr", os.Stderr)
	ctx = context.WithValue(ctx, "gosh.stdin", os.Stdin)

	shell := gosh.New()
	if err := shell.Init(ctx); err != nil {
		fmt.Println("\n\nfailed to initialize:\n", err)
		os.Exit(1)
	}

	// prompt for help
	cmdCount := len(shell.Commands)
	if cmdCount > 0 {
		if _, ok := shell.Commands["help"]; ok {
			fmt.Printf("\nLoaded %d command(s)...", cmdCount)
			fmt.Println("\nType help for available commands")
			fmt.Print("\n")
		}
	} else {
		fmt.Print("\n\nNo commands found")
	}

	go shell.Open(bufio.NewReader(os.Stdin))

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT)
	select {
	case <-sigs:
		cancel()
		<-shell.Closed()
	case <-shell.Closed():
	}
}
