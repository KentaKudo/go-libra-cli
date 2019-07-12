package main

import (
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

type account struct{}

func (a account) action(cmd *cli.Cmd) {
	cmd.Command("create", "Create a random account with private/public key pair", a.create)
	cmd.Command("list", "Print all accounts that were created or loaded", a.list)
	cmd.Command("recover", "Recover all accounts that were written to a file via the `account write` command", a.recover)
	cmd.Command("write", "Save Libra wallet mnemonic recovery seed to disk", a.write)
	cmd.Command("mint", "Mint coins to the account", a.mint)

	cmd.Action = func() {
		log.Println("account command help is not implemented yet")
	}
}

func (a account) create(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("account create command is not implemented yet")
	}
}

func (a account) list(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("account list command is not implemented yet")
	}
}

func (a account) recover(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("account recover command is not implemented yet")
	}
}

func (a account) write(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("account write command is not implemented yet")
	}
}

func (a account) mint(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("account mint command is not implemented yet")
	}
}
