package api

import (
	"context"
	"io"
	"os"
)

type ShellModule interface {
	Init(ctx context.Context) error
}

type Command interface {
	Name() string
	Usage() string
	ShortDesc() string
	LongDesc() string
	Exec(context.Context, []string) (context.Context, error)
}

type Commands interface {
	ShellModule
	Registry() map[string]Command
}

const (
	PluginsDir    = "./plugins"
	CmdSymbolName = "Commands"
	DefaultPrompt = "goshell>"
)

func GetStdout(ctx context.Context) io.Writer {
	var out io.Writer = os.Stdout
	if ctx == nil {
		return out
	}
	if outVal := ctx.Value("gosh.stdout"); outVal != nil {
		if stdout, ok := outVal.(io.Writer); ok {
			out = stdout
		}
	}
	return out
}

func GetPrompt(ctx context.Context) string {
	prompt := DefaultPrompt
	if ctx == nil {
		return prompt
	}
	if promptVal := ctx.Value("gosh.prompt"); promptVal != nil {
		if p, ok := promptVal.(string); ok {
			prompt = p
		}
	}
	return prompt
}
