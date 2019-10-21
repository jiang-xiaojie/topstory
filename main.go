package main

import (
	"os"

	"github.com/jianggushi/topstory/cmd"
	_ "github.com/jianggushi/topstory/models"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "topstory"
	app.Usage = "a topstory"
	app.Commands = []cli.Command{
		cmd.Web,
		cmd.Spider,
	}
	app.Run(os.Args)
}
