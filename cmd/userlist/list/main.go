package list

import (
	"context"
	"flag"
	"fmt"
	"log/slog"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/userlist"
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
	db, ok := args[0].(userlist.DB)
	if !ok {
		slog.Error("cast error")
		return subcommands.ExitFailure
	}

	if err := db.Scan(ctx, func(u *userlist.User) {
		fmt.Println(u.Name)
	}); err != nil {
		slog.Error("scan error")
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
