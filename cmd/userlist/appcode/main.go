package appcode

import (
	"context"
	"flag"
	"log/slog"
	"strings"

	"github.com/google/subcommands"

	"github.com/yunomu/auth/lib/db/userlist"
)

type Command struct {
	user *string
	add  *string
	del  *string
}

func NewCommand() *Command {
	return &Command{}
}

func (c *Command) Name() string     { return "appcode" }
func (c *Command) Synopsis() string { return "appcode" }
func (c *Command) Usage() string {
	return `
`
}

func (c *Command) SetFlags(f *flag.FlagSet) {
	c.user = f.String("user", "", "User name")
	c.add = f.String("add", "", "appcode1[,appcode1,...]")
	c.del = f.String("del", "", "appcode1[,appcode1,...]")
}

func addCodes(orig, codes []string) []string {
	if len(codes) == 0 {
		return orig
	}

	m := make(map[string]struct{})
	for _, c := range orig {
		m[c] = struct{}{}
	}
	for _, c := range codes {
		m[c] = struct{}{}
	}

	var ret []string
	for c := range m {
		ret = append(ret, c)
	}

	return ret
}

func delCodes(orig, codes []string) []string {
	if len(codes) == 0 {
		return orig
	}

	m := make(map[string]struct{})
	for _, c := range orig {
		m[c] = struct{}{}
	}
	for _, c := range codes {
		delete(m, c)
	}

	var ret []string
	for c := range m {
		ret = append(ret, c)
	}

	return ret
}

func (c *Command) Execute(ctx context.Context, _ *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	db, ok := args[0].(userlist.DB)
	if !ok {
		slog.Error("cast error")
		return subcommands.ExitFailure
	}

	if *c.user == "" {
		slog.Error("user name is required")
		return subcommands.ExitFailure
	}
	if *c.add == "" && *c.del == "" {
		slog.Error("add/del is required")
		return subcommands.ExitFailure
	}

	user, _, err := db.Get(ctx, *c.user)
	if err != nil {
		slog.Error("db.Get", "err", err)
		return subcommands.ExitFailure
	}

	user.AppCodes = addCodes(user.AppCodes, strings.Split(*c.add, ","))
	user.AppCodes = delCodes(user.AppCodes, strings.Split(*c.del, ","))

	if err := db.Put(ctx, user); err != nil {
		slog.Error("db.Put", "err", err)
		return subcommands.ExitFailure
	}

	return subcommands.ExitSuccess
}
