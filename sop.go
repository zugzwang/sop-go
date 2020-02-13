package main

import (
	"os"
	"fmt"
	"log"
	"sort"

	"github.com/urfave/cli"
)

var version = "0.1.0"

// SOPGOComment is the comment expected by openpgp.NewEntity
var SOPGOComment = "sop-go version " + version


func main() {

	app := &cli.App{
		Name:  "sop-go",
		Usage: "Stateless OpenPGP implementation for github.com/ProtonMail/crypto",
		Commands: []*cli.Command{
			{
				Name:  "version",
				Usage: "Version Information",
				Action: func(c *cli.Context) error {
					fmt.Printf("sop-go %v\n", version)
					return nil
				},
			},
			{
				Name:  "generate-key",
				Usage: "Generate a Secret Key",
				Flags: []cli.Flag{
					ArmorFlag,
					NoArmorFlag,
				},
				Action: func(c *cli.Context) error {
					return GenerateKey(c.Args().Slice()...)
				},
			},
			{
				Name:  "extract-cert",
				Usage: "Extract a Certificate from a Secret Key",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "sign",
				Usage: "Create a Detached Signature",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "verify",
				Usage: "Verify a Detached Signature",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "encrypt",
				Usage: "Encrypt a Message",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a Message",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "armor",
				Usage: "Add ASCII Armor",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
			{
				Name:  "dearmor",
				Usage: "Remove ASCII Armor",
				Action: func(c *cli.Context) error {
					return Err69
				},
			},
		},
		Action: func(c *cli.Context) error {
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	Err := app.Run(os.Args)
	if Err != nil {
		log.Fatal(Err)
	}
}
