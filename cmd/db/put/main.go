package put

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
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "put" }
func (c *Command) Synopsis() string { return "put < record" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(productdb.DB)
	if !ok {
		log.Printf("cast error")
		return subcommands.ExitFailure
	}

	dec := json.NewDecoder(os.Stdin)
	var rec productdb.Record
	if err := dec.Decode(&rec); err != nil {
		log.Printf("json.Decode: %v", err)
		return subcommands.ExitFailure
	}

	if err := db.Put(ctx, &rec); err != nil {
		log.Printf("db.Put: %v", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
