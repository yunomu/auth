package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/cmd/db"
)

var ()

func init() {
	log.SetOutput(os.Stderr)

	subcommands.Register(db.NewCommand(), "")

	subcommands.Register(subcommands.CommandsCommand(), "other")
	subcommands.Register(subcommands.FlagsCommand(), "other")
	subcommands.Register(subcommands.HelpCommand(), "other")

	flag.Parse()
}

func main() {
	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx)))
}
