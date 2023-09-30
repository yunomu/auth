package get

import (
	"context"
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/productdb"
)

type Command struct {
	clientId *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "get" }
func (c *Command) Synopsis() string { return "get clientId" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	c.clientId = f.String("clientId", "", "Cognito Client ID")
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(productdb.DB)
	if !ok {
		slog.Info("cast error")
		return subcommands.ExitFailure
	}

	rec, _, err := db.Get(ctx, *c.clientId)
	if err != nil {
		slog.Error("db.Get", "err", err)
		return subcommands.ExitFailure
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(rec); err != nil {
		slog.Error("json.Encode", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
