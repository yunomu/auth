package list

import (
	"context"
	"encoding/csv"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/productdb"
)

type Command struct {
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "list" }
func (c *Command) Synopsis() string { return "list" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(productdb.DB)
	if !ok {
		slog.Info("cast error")
		return subcommands.ExitFailure
	}

	w := csv.NewWriter(os.Stdout)
	w.Comma = '\t'
	w.Write([]string{
		"ClientId",
		"AppCode",
		"Created",
		"FuncArn",
	})

	if err := db.Scan(ctx, func(rec *productdb.Record) {
		t := time.Unix(rec.Created, 0)
		w.Write([]string{
			rec.ClientId,
			rec.AppCode,
			t.Format(time.RFC3339),
			rec.FuncArn,
		})
	}); err != nil {
		slog.Error("db.Scan", "err", err)
		return subcommands.ExitFailure
	}

	w.Flush()

	return subcommands.ExitSuccess
}
