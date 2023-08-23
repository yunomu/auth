package config

import (
	"context"
	"encoding/json"
	"flag"
	"log"
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

func (c *Command) Name() string     { return "config" }
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
		log.Fatal("stack name is required")
	}

	sess, err := session.NewSession(aws.NewConfig().WithRegion(*c.region))
	if err != nil {
		log.Fatalf("NewSession: %v", err)
	}

	client := cloudformation.New(sess)

	out, err := client.DescribeStacksWithContext(ctx, &cloudformation.DescribeStacksInput{
		StackName: aws.String(*c.stack),
	})
	if err != nil {
		log.Fatalf("DescribeStacks: %v", err)
	}

	if len(out.Stacks) == 0 {
		log.Fatalf("No stacks")
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
		log.Fatalf("json.Encode: %v", err)
	}

	return subcommands.ExitSuccess
}
