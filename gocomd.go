package main

import (
	"flag"
	"fmt"
	"github.com/urfave/cli"
	"gocomd/core/logger"
	"gocomd/gen"
	"os"
	"runtime"
	"strings"
)

var commands = []cli.Command{
	{
		Name:  "gen",
		Usage: "generate template file",
		Subcommands: []cli.Command{
			{
				Name:  "mysql",
				Usage: "generate mysql template file",
				Flags: []cli.Flag{
					cli.StringFlag{
						Name:  "dir",
						Usage: "file directory",
					},
					cli.StringFlag{
						Name:  "template",
						Value: "migrate",
						Usage: "template name",
					},
					cli.StringFlag{
						Name:  "file",
						Usage: "file name",
					},
				},
				Action: gen.GoCommand,
			},
		},
	},
}

var configFileName = flag.String("cfn", "config", "name of configs file")
var configFilePath = flag.String("cfp", "./core/configs/", "path of configs file")

func main() {
	flag.Parse()
	err := logger.InitLogger(*configFileName, strings.Split(*configFilePath, ","))
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Usage = "a cli tool to generate code"
	app.Version = fmt.Sprintf("%s %s/%s", "1.3.5", runtime.GOOS, runtime.GOARCH)
	app.Commands = commands

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
