package main

import (
	"context"
	"encoding/json"
	"flag"
	"log/slog"
	"os"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/cmd/config"
	"github.com/yunomu/auth/cmd/productdb"
	"github.com/yunomu/auth/cmd/unmarshaler"
	"github.com/yunomu/auth/cmd/userlist"
)

var (
	configPath = flag.String("config", ".config", "config file path")
)

func init() {
	subcommands.Register(productdb.NewCommand(), "")
	subcommands.Register(userlist.NewCommand(), "")
	subcommands.Register(config.NewCommand(), "")
	subcommands.Register(unmarshaler.NewCommand(), "")

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
		slog.Info("loadConfig",
			slog.String("path", *configPath),
			slog.Any("err", err),
		)
	}

	ctx := context.Background()

	os.Exit(int(subcommands.Execute(ctx, cfg)))
}
