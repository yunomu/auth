package unmarshaler

import (
	"context"
	"flag"
	"io"
	"io/ioutil"
	"log/slog"
	"os"

	"github.com/google/subcommands"
	"github.com/yunomu/auth/proto/api"
	"google.golang.org/protobuf/encoding/protojson"
)

type Command struct {
	input *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "unmarshaler" }
func (c *Command) Synopsis() string { return "unmarshaler test" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	f.SetOutput(os.Stderr)

	c.input = f.String("in", "-", "Input file")
}

func (c *Command) Execute(ctx context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	var in io.Reader = os.Stdin
	if *c.input != "-" {
		f, err := os.Open(*c.input)
		if err != nil {
			slog.Error("Open", "err", err, "file", *c.input)
			return subcommands.ExitFailure
		}
		defer f.Close()
		in = f
	}

	bs, err := ioutil.ReadAll(in)
	if err != nil {
		slog.Error("ReadAll", "err", err)
		return subcommands.ExitFailure
	}

	var res api.RestrictionsResponse
	unmarshaler := protojson.UnmarshalOptions{}
	if err := unmarshaler.Unmarshal(bs, &res); err != nil {
		slog.Error("Unmarshal", "err", err, "string", string(bs))
		return subcommands.ExitFailure
	}

	slog.Info("unmarshal", "res", res)

	return subcommands.ExitSuccess
}
