package main

import "github.com/urfave/cli"

// Error codes as defined in the draft, section 6
var (
	Err3  = cli.Exit("No acceptable signatures found (\"sop verify\")", 3)
	Err13 = cli.Exit("Asymmetric algorithm unsupported (\"sop encrypt\")", 13)
	Err17 = cli.Exit(`Certificate not encryption-capable (e.g., expired,
	revoked, unacceptable usage flags) ("sop encrypt")`, 17)
	Err19 = cli.Exit("Missing required argument", 19)
	Err23 = cli.Exit(`Incomplete verification instructions ("sop decrypt")`, 23)
	Err29 = cli.Exit(`Unable to decrypt ("sop decrypt")`, 29)
	Err31 = cli.Exit(`Non-"UTF-8" password ("sop encrypt")`, 31)
	Err37 = cli.Exit("Unsupported option", 37)
	Err41 = cli.Exit(`Invalid data type (no secret key where "KEY" expected, etc)`, 41)
	Err53 = cli.Exit("Non-text input where text expected", 53)
	Err69 = cli.Exit("Unsupported subcommand", 69)
)

// Err99 returns the error message of any error not defined above
func Err99(err error) error {
	return cli.Exit("Unexpected error: "+err.Error(), 99)
}
