package gosh

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"plugin"
	"qfluent-go/pkg/gosh/api"
	"regexp"
	"strings"
)

var (
	reCmd = regexp.MustCompile(`\S+`)
)

type GoShell struct {
	ctx        context.Context
	pluginsDir string
	Commands   map[string]api.Command
	closed     chan struct{}
}

func New() *GoShell {
	return &GoShell{
		pluginsDir: api.PluginsDir,
		Commands:   make(map[string]api.Command),
		closed:     make(chan struct{}),
	}
}

func (gosh *GoShell) Init(ctx context.Context) error {
	gosh.ctx = ctx
	return gosh.loadCommands()
}

func (gosh *GoShell) loadCommands() error {
	if _, err := os.Stat(gosh.pluginsDir); err != nil {
		return err
	}
	plugins, err := listFiles(gosh.pluginsDir, `.*_command.so`)
	if err != nil {
		return err
	}
	for _, cmdPlugin := range plugins {
		plug, err := plugin.Open(path.Join(gosh.pluginsDir, cmdPlugin.Name()))
		if err != nil {
			fmt.Printf("failed to open plugin %s: %v\n", cmdPlugin.Name(), err)
			continue
		}
		cmdSymbol, err := plug.Lookup(api.CmdSymbolName)
		if err != nil {
			fmt.Printf("plugin %s does not export symbol \"%s\"\n",
				cmdPlugin.Name(), api.CmdSymbolName)
			continue
		}
		commands, ok := cmdSymbol.(api.Commands)
		if !ok {
			fmt.Printf("Symbol %s (from %s) does not implement Commands interface\n",
				api.CmdSymbolName, cmdPlugin.Name())
			continue
		}
		if err := commands.Init(gosh.ctx); err != nil {
			fmt.Printf("%s initialization failed: %v\n", cmdPlugin.Name(), err)
			continue
		}
		for name, cmd := range commands.Registry() {
			gosh.Commands[name] = cmd
		}
		gosh.ctx = context.WithValue(gosh.ctx, "gosh.commands", gosh.Commands)
	}
	return nil
}
func (gosh *GoShell) Open(r *bufio.Reader) {
	loopCtx := gosh.ctx
	line := make(chan string)
	for {
		// start a goroutine to get input from the user
		go func(ctx context.Context, input chan<- string) {
			for {
				// TODO: future enhancement is to capture input key by key
				// to give command granular notification of key events.
				// This could be used to implement command autocompletion.
				_, _ = fmt.Fprintf(ctx.Value("gosh.stdout").(io.Writer), "%s ", api.GetPrompt(loopCtx))
				line, err := r.ReadString('\n')
				if err != nil {
					_, _ = fmt.Fprintf(ctx.Value("gosh.stderr").(io.Writer), "%v\n", err)
					continue
				}

				input <- line
				return
			}
		}(loopCtx, line)

		// wait for input or cancel
		select {
		case <-gosh.ctx.Done():
			close(gosh.closed)
			return
		case input := <-line:
			var err error
			loopCtx, err = gosh.handle(loopCtx, input)
			if err != nil {
				_, _ = fmt.Fprintf(loopCtx.Value("gosh.stderr").(io.Writer), "%v\n", err)
			}
		}
	}
}
func (gosh *GoShell) Closed() <-chan struct{} {
	return gosh.closed
}

func (gosh *GoShell) handle(ctx context.Context, cmdLine string) (context.Context, error) {
	line := strings.TrimSpace(cmdLine)
	if line == "" {
		return ctx, nil
	}
	args := reCmd.FindAllString(line, -1)
	if args != nil {
		cmdName := args[0]
		cmd, ok := gosh.Commands[cmdName]
		if !ok {
			return ctx, errors.New(fmt.Sprintf("command not found: %s", cmdName))
		}
		return cmd.Exec(ctx, args)
	}
	return ctx, errors.New(fmt.Sprintf("unable to parse command line: %s", line))
}
func listFiles(dir, pattern string) ([]os.FileInfo, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filteredFiles []os.FileInfo
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		matched, err := regexp.MatchString(pattern, file.Name())
		if err != nil {
			return nil, err
		}
		if matched {
			filteredFiles = append(filteredFiles, file)
		}
	}
	return filteredFiles, nil
}
