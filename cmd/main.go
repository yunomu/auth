package main

import (
	"context"
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/cmd/config"
	"github.com/yunomu/auth/cmd/db"
	"github.com/yunomu/auth/cmd/s3"
)

var (
	configPath = flag.String("config", ".config", "config file path")
)

func init() {
	log.SetOutput(os.Stderr)

	subcommands.Register(db.NewCommand(), "")
	subcommands.Register(s3.NewCommand(), "")
	subcommands.Register(config.NewCommand(), "")

	subcommands.Register(subcommands.CommandsCommand(), "other")
	subcommands.Register(subcommands.FlagsCommand(), "other")
	subcommands.Register(subcommands.HelpCommand(), "other")

	flag.Parse()
}

func loadConfig(file string) (map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	ret := make(map[string]string)
	if err := dec.Decode(&ret); err != nil {
		return nil, err
	}

	return ret, nil
}

func main() {
	cfg, err := loadConfig(*configPath)
	if err != nil {
		log.Printf("loadConfig: path=%v, err=%v", *configPath, err)
	}

	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx, cfg)))
}
