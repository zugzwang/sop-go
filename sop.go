package main

import (
	"log"
	"os"
	"fmt"
	"sort"

	"github.com/urfave/cli"
)

var (
	version = "0.1.0"
	commandNotImplemented = cli.Exit("Command not implemented", 69)
	subCommandNotImplemented = cli.Exit("Subcommand not implemented", 37)
)

func main() {

	app := &cli.App{
		Name:  "sop-go",
		Usage: "Stateless OpenPGP implementation for ProtonMail/crypto",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "armor",
				Value: true,
			},
			&cli.BoolFlag{
				Name: "no-armor",
			},
			&cli.StringFlag{
				Name: "as",
				Value: "binary",
				Usage: "--as={binary|text}",
			},
		},
		Commands: []*cli.Command{
			{
				Name: "version",
				Usage: "Version Information",
				Action: func(c *cli.Context) error {
					fmt.Printf("sop-go %v\n", version)
					return nil
				},
			},
			{
				Name: "generate-key",
				Usage: "Generate a Secret Key",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name: "extract-cert",
				Usage: "Extract a Certificate from a Secret Key",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name: "sign",
				Usage: "Create a Detached Signature",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name: "verify",
				Usage: "Verify a Detached Signature",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name:  "encrypt",
				Usage: "Encrypt a Message",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a Message",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name:  "armor",
				Usage: "Add ASCII Armor",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
			{
				Name:  "dearmor",
				Usage: "Remove ASCII Armor",
				Action: func(c *cli.Context) error {
					return commandNotImplemented
				},
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
