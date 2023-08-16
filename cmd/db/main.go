package db

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/subcommands"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/yunomu/auth/lib/db/productdb"

	"github.com/yunomu/auth/cmd/db/get"
	"github.com/yunomu/auth/cmd/db/list"
	"github.com/yunomu/auth/cmd/db/put"
)

type Command struct {
	endpoint *string
	region   *string
	table    *string
	log      *bool

	commander *subcommands.Commander
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "db" }
func (c *Command) Synopsis() string { return "db test" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.endpoint = f.String("endpoint", "", "Endpoint of DynamoDB")
	c.region = f.String("region", "ap-northeast-1", "AWS Region")
	c.table = f.String("table", "", "Table name")
	c.log = f.Bool("log", false, "output log")

	commander := subcommands.NewCommander(f, "")
	commander.Register(commander.FlagsCommand(), "help")
	commander.Register(commander.CommandsCommand(), "help")
	commander.Register(commander.HelpCommand(), "help")

	commander.Register(get.NewCommand(), "")
	commander.Register(put.NewCommand(), "")
	commander.Register(list.NewCommand(), "")

	c.commander = commander
}

type logger struct{}

func (l *logger) Info(function string, message string) {
	log.Println(function, message)
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if *c.table == "" {
		log.Fatal("table name is required")
	}

	return c.commander.Execute(ctx, func() productdb.DB {
		config := aws.NewConfig().WithRegion(*c.region)

		if *c.endpoint != "" {
			config.WithEndpoint(*c.endpoint)
		}

		return productdb.NewDynamoDB(
			dynamodb.New(session.New(), config),
			*c.table,
		)
	})
}
