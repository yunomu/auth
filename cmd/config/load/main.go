package load

import (
	"context"
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

type Command struct {
	stack  *string
	region *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "load-from-stack" }
func (c *Command) Synopsis() string { return "output config" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.stack = f.String("stack", "", "Stack name")
	c.region = f.String("region", "ap-northeast-1", "AWS Region")
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if *c.stack == "" {
		slog.Error("stack name is required")
		return subcommands.ExitFailure
	}

	sess, err := session.NewSession(aws.NewConfig().WithRegion(*c.region))
	if err != nil {
		slog.Error("NewSession", "err", err)
		return subcommands.ExitFailure
	}

	client := cloudformation.New(sess)

	out, err := client.DescribeStacksWithContext(ctx, &cloudformation.DescribeStacksInput{
		StackName: aws.String(*c.stack),
	})
	if err != nil {
		slog.Error("DescribeStacks", "err", err)
		return subcommands.ExitFailure
	}

	if len(out.Stacks) == 0 {
		slog.Error("No stacks")
		return subcommands.ExitFailure
	}
	stack := out.Stacks[0]

	outputs := make(map[string]string)
	for _, p := range stack.Parameters {
		outputs[aws.StringValue(p.ParameterKey)] = aws.StringValue(p.ParameterValue)
	}
	for _, o := range stack.Outputs {
		outputs[aws.StringValue(o.OutputKey)] = aws.StringValue(o.OutputValue)
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(outputs); err != nil {
		slog.Error("json.Encode", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
