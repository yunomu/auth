package add

import (
	"context"
	"flag"
	"log/slog"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/productdb"
)

type Command struct {
	clientId *string
	appCode  *string
	funcArn  *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "add" }
func (c *Command) Synopsis() string { return "add" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	c.clientId = f.String("clientId", "", "UserPoolClientId (required)")
	c.appCode = f.String("app", "", "AppCode (required)")
	c.funcArn = f.String("func", "", "Lambda function arn (optional)")
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if *c.clientId == "" {
		slog.Error("clientId is required")
		return subcommands.ExitFailure
	}
	if *c.appCode == "" {
		slog.Error("app is required")
		return subcommands.ExitFailure
	}

	db, ok := args[0].(productdb.DB)
	if !ok {
		slog.Error("cast error")
		return subcommands.ExitFailure
	}

	version, err := db.Put(ctx, &productdb.Product{
		ClientId: *c.clientId,
		AppCode:  *c.appCode,
		FuncArn:  *c.funcArn,
	})
	if err != nil {
		slog.Error("db.Put", "err", err)
		return subcommands.ExitFailure
	}
	slog.Info("SUCCESS", "version", version)

	return subcommands.ExitSuccess
}
