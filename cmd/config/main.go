package config

import (
	"context"
	"flag"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/cmd/config/console"
	"github.com/yunomu/auth/cmd/config/load"
)

type Command struct {
	commander *subcommands.Commander
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "config" }
func (c *Command) Synopsis() string { return "config utils" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	commander := subcommands.NewCommander(f, "")
	commander.Register(commander.FlagsCommand(), "help")
	commander.Register(commander.CommandsCommand(), "help")
	commander.Register(commander.HelpCommand(), "help")

	commander.Register(load.NewCommand(), "")
	commander.Register(console.NewCommand(), "")

	c.commander = commander
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	return c.commander.Execute(ctx, args...)
}
