package put

import (
	"context"
	"flag"
	"log/slog"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/userlist"
)

type Command struct {
	user *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "put" }
func (c *Command) Synopsis() string { return "put" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	c.user = f.String("user", "", "User name")
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(userlist.DB)
	if !ok {
		slog.Error("cast error")
		return subcommands.ExitFailure
	}

	if *c.user == "" {
		slog.Error("user name is required")
		return subcommands.ExitFailure
	}

	if err := db.Put(ctx, &userlist.User{
		Name: *c.user,
	}); err != nil {
		slog.Error("db.Put", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
