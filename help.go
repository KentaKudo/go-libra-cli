package main

import (
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

type help struct{}

func (h help) action(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("help command is not implemented yet")
	}
}
