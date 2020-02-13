package main

import "github.com/urfave/cli"

// Variables defined by flags
var (
	NoArmor bool
)

// All possible flags for subcommands
var (
	ArmorFlag = &cli.BoolFlag{
		Name:  "armor",
		Value: true,
	}
	NoArmorFlag = &cli.BoolFlag{
		Name: "no-armor",
		Value: false,
		Destination: &NoArmor,
	}
	AsFlag = &cli.StringFlag{
		Name:  "as",
		Value: "binary",
		Usage: "--as={binary|text}",
	}
)
