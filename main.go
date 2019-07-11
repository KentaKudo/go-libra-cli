package main

import (
	"os"

	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

const (
	appName = "go-libra-cli"
	appDesc = "The command line tool to communicate with Libra nodes in Go"
)

func main() {
	app := cli.App(appName, appDesc)

	app.Action = func() {
		log.Printf("hello world!")
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatalln("app run")
	}
}
