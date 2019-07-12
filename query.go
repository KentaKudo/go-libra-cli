package main

import (
	cli "github.com/jawher/mow.cli"
	log "github.com/sirupsen/logrus"
)

type query struct{}

func (q query) action(cmd *cli.Cmd) {
	cmd.Command("balance", "Get the current balance of an account", q.balance)
	cmd.Command("sequence", "Get the current sequence number for an account", q.sequence)
	cmd.Command("account_state", "Get the latest state for an account", q.accountState)
	cmd.Command("txn_acc_seq", "Get the committed transaction by account and sequence number", q.txnAccSeq)
	cmd.Command("txn_range", "Get the committed transaction by range", q.txnRange)
	cmd.Command("event", "Get event by account and path", q.event)

	cmd.Action = func() {
		log.Println("query command help is not implemented yet")
	}
}

func (q query) balance(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query balance command is not implemented yet")
	}
}

func (q query) sequence(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query sequence command is not implemented yet")
	}
}

func (q query) accountState(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query account_state command is not implemented yet")
	}
}

func (q query) txnAccSeq(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query txn_acc_seq command is not implemented yet")
	}
}

func (q query) txnRange(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query txn_range command is not implemented yet")
	}
}

func (q query) event(cmd *cli.Cmd) {
	cmd.Action = func() {
		log.Println("query event command is not implemented yet")
	}
}
