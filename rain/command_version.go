package rain

import (
	"fmt"

	"github.com/urfave/cli"
)

func ver(c *cli.Context) error {
	fmt.Println(Version())
	return nil
}

var CommandVersion = cli.Command{
	Name:    "version",
	Aliases: []string{"v"},
	Usage:   "Shows ONLY the version, as opposed to -(-v)ersion",
	Action:  ver,
}
