package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/libreofficedocker/go-unoserver/unoserver"
	"github.com/urfave/cli"
)

var Version = "0.0.0"

func main() {
	app := cli.NewApp()
	app.Version = Version
	app.Usage = "A Go implementation for unoserver"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "addr",
			Value: "127.0.0.1:2002",
			Usage: "The addr used by the unoserver api server",
		},
		cli.StringFlag{
			Name:  "user-installation",
			Usage: "The path to the user installation directory. If not specified, the default user installation will be used.",
		},
		cli.BoolFlag{
			Name:  "daemon",
			Usage: "Run as daemon",
		},
	}
	app.Authors = []cli.Author{
		{
			Name:  "libreofficedocker",
			Email: "https://github.com/libreofficedocker",
		},
	}
	app.Action = action

	// Set log prefix
	log.SetPrefix(app.Name + ": ")

	if err := app.Run(os.Args); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	host, port, err := net.SplitHostPort(c.String("addr"))
	if err != nil {
		return err
	}

	server := unoserver.Default()
	server.Host = host
	server.Port = port

	userInstallation := c.String("user-installation")
	if userInstallation != "" {
		server.SetUserInstallation(userInstallation)
	}

	cmd := server.CommandContext(ctx)

	if c.Bool("daemon") {
		return cmd.Start()
	}

	return cmd.Run()
}
