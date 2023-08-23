package put

import (
	"context"
	"flag"
	"log"

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
		log.Printf("cast error")
		return subcommands.ExitFailure
	}

	if *c.user == "" {
		log.Printf("user name is required")
		return subcommands.ExitFailure
	}

	if err := db.Put(ctx, &userlist.User{
		Name: *c.user,
	}); err != nil {
		log.Printf("db.Put: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
