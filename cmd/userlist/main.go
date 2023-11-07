package userlist

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/yunomu/auth/cmd/userlist/appcode"
	"github.com/yunomu/auth/cmd/userlist/list"
	"github.com/yunomu/auth/cmd/userlist/put"
	"github.com/yunomu/auth/lib/db/userlist"
)

type Command struct {
	region *string
	table  *string

	commander *subcommands.Commander
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "userlist" }
func (c *Command) Synopsis() string { return "userlist util" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.table = f.String("table", "", "RestrictionTable name")
	c.region = f.String("region", "ap-northeast-1", "AWS Region")

	commander := subcommands.NewCommander(f, "")
	commander.Register(commander.FlagsCommand(), "help")
	commander.Register(commander.CommandsCommand(), "help")
	commander.Register(commander.HelpCommand(), "help")

	commander.Register(list.NewCommand(), "")
	commander.Register(put.NewCommand(), "")
	commander.Register(appcode.NewCommand(), "")

	c.commander = commander
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var table string
	cfg, ok := args[0].(map[string]string)
	if ok {
		table = cfg["RestrictionTable"]
	}
	if *c.table != "" {
		table = *c.table
	}

	if table == "" {
		slog.Error("table name is required")
		return subcommands.ExitFailure
	}

	sess, err := session.NewSession(aws.NewConfig().WithRegion(*c.region))
	if err != nil {
		slog.Error("NewSession", "err", err)
		return subcommands.ExitFailure
	}

	return c.commander.Execute(ctx, userlist.NewDynamoDB(
		dynamodb.New(sess),
		table,
	))
}
