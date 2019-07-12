package main

import (
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

type transfer struct{}

func (t transfer) action(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("transfer command is not implemented yet")
	}
}
