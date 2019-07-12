package main

import (
	"context"
	"os"

	"github.com/KentaKudo/go-libra-cli/pb/admission_control"
	cli "github.com/jawher/mow.cli"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

const (
	appName = "go-libra-cli"
	appDesc = "The command line tool to communicate with Libra nodes in Go"
)

func main() {
	app := cli.App(appName, appDesc)

	destinationURL := app.String(cli.StringOpt{
		Name:   "destination-url",
		Desc:   "Destination URL",
		EnvVar: "DESTINATION_URL",
		Value:  "localhost:30307",
	})

	app.Action = func() {
		log.Printf("hello world!")

		conn, err := initialiseGRPCClientConnection(context.Background(), *destinationURL)
		if err != nil {
			log.WithError(err).Fatalln("init gRPC conn")
		}

		_ = admission_control.NewAdmissionControlClient(conn)
	}

	if err := app.Run(os.Args); err != nil {
		log.WithError(err).Fatalln("app run")
	}
}

func initialiseGRPCClientConnection(ctx context.Context, addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.DialContext(ctx,
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, errors.Wrapf(err, "grpc dial: %s", addr)
	}

	return conn, nil
}
