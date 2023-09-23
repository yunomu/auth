package s3

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"
	"github.com/yunomu/auth/cmd/s3/list"
	"github.com/yunomu/auth/cmd/s3/put"
	"github.com/yunomu/auth/lib/userlist"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Command struct {
	region       *string
	bucket       *string
	whiteListKey *string

	commander *subcommands.Commander
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "s3" }
func (c *Command) Synopsis() string { return "s3 util" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.bucket = f.String("bucket", "", "Bucket name")
	c.whiteListKey = f.String("white-list-key", "", "Whitelist object key")

	commander := subcommands.NewCommander(f, "")
	commander.Register(commander.FlagsCommand(), "help")
	commander.Register(commander.CommandsCommand(), "help")
	commander.Register(commander.HelpCommand(), "help")

	commander.Register(list.NewCommand(), "")
	commander.Register(put.NewCommand(), "")

	c.commander = commander
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var bucket, whiteListKey string
	cfg, ok := args[0].(map[string]string)
	if ok {
		bucket = cfg["Bucket"]
		whiteListKey = cfg["WhiteListPath"]
	}
	if *c.bucket != "" {
		bucket = *c.bucket
	}
	if *c.whiteListKey != "" {
		whiteListKey = *c.whiteListKey
	}

	if bucket == "" {
		slog.Error("bucket name is required")
		return subcommands.ExitFailure
	}
	if whiteListKey == "" {
		slog.Error("white-list-key is required")
		return subcommands.ExitFailure
	}

	sess, err := session.NewSession()
	if err != nil {
		slog.Error("NewSession", "err", err)
		return subcommands.ExitFailure
	}

	return c.commander.Execute(ctx, userlist.NewS3(
		s3.New(sess),
		bucket, whiteListKey,
	))
}
