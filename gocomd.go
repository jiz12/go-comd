package main

import (
	"flag"
	"github.com/jiz12/go-comd/cmd"
	"github.com/jiz12/go-comd/core/logger"
	"strings"
)

var configFileName = flag.String("cfn", "config", "name of configs file")
var configFilePath = flag.String("cfp", "./core/configs/", "path of configs file")

func main() {
	flag.Parse()
	err := logger.InitLogger(*configFileName, strings.Split(*configFilePath, ","))
	if err != nil {
		panic(err)
	}
	cmd.Execute()
}
