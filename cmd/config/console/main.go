package console

import (
	"context"
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"
)

type Command struct {
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "gen-for-console" }
func (c *Command) Synopsis() string { return "output config" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	cfg, ok := args[0].(map[string]string)
	if !ok {
		slog.Error("config file is required")
		return subcommands.ExitFailure
	}

	outputs := map[string]string{
		"AuthRedirectURL":   cfg["CallbackURL"],
		"LogoutRedirectURL": cfg["LogoutURL"],
		"IDP":               cfg["DomainName"],
		"UserPoolClientId":  cfg["AdminUserPoolClientId"],
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(outputs); err != nil {
		slog.Error("json.Encode", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
