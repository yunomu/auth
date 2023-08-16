package get

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/productdb"
)

type Command struct {
	clientId *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "get" }
func (c *Command) Synopsis() string { return "get clientId" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	c.clientId = f.String("clientId", "", "Cognito Client ID")
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	f, ok := args[0].(func() productdb.DB)
	if !ok {
		log.Fatalf("cast error")
	}
	db := f()

	rec, err := db.Get(ctx, *c.clientId)
	if err != nil {
		log.Printf("db.Get: %v", err)
		return subcommands.ExitFailure
	}

	enc := json.NewEncoder(os.Stdout)
	if err := enc.Encode(rec); err != nil {
		log.Printf("json.Encode: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
