package put

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
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "put" }
func (c *Command) Synopsis() string { return "put < record" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(productdb.DB)
	if !ok {
		slog.Error("cast error")
		return subcommands.ExitFailure
	}

	dec := json.NewDecoder(os.Stdin)
	var rec productdb.Product
	if err := dec.Decode(&rec); err != nil {
		slog.Error("json.Decode", "err", err)
		return subcommands.ExitFailure
	}

	if err := db.Put(ctx, &rec); err != nil {
		slog.Error("db.Put", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
