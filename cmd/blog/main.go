package main

import(
	"blog.com/packages/cmd/internal/server";
	"blog.com/packages/cmd/internal/conf";
	"blog.com/packages/cmd/internal/cli";
)

func main() {
	env := cli.Parse()
	server.Start(conf.NewConfig(env))
}
