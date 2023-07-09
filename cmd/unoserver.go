package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/libreofficedocker/go-unoserver/unoserver"
	"github.com/urfave/cli"
)

var Name = "unoserver"
var Version = "0.0.0"

func init() {
	log.SetPrefix(Name + ": ")
}

func main() {
	app := cli.NewApp()
	app.Name = Name
	app.Version = Version
	app.Usage = "A Go implementation for unoserver"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "0.0.0.0:2004",
			Usage: "The addr used by the unoserver api server",
		},
		cli.BoolFlag{
			Name:  "daemon",
			Usage: "Run as daemon",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreofficedocker",
			Email: "https://github.com/libreofficedocker/unoserver-rest-api",
		},
	}
	app.Action = mainAction

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func mainAction(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	host, port, err := net.SplitHostPort(c.String("addr"))

	if err != nil {
		log.Fatal(err)
	}

	server := unoserver.Default()
	server.Host = host
	server.Port = port

	cmd := server.CommandContext(ctx)

	if c.Bool("daemon") {
		return cmd.Start()
	}

	return cmd.Run()
}
