package rain

import (
	"fmt"

	"github.com/urfave/cli"
)

func deps(c *cli.Context) error {
	fmt.Println(" Getting dependencies... ")
	err := GetDepends(AllDepends)
	if err != nil {
		fmt.Println("An error has occurred, please review the above and report as an issue if you can")
		return err
	}
	fmt.Println(" Dependencies installed!")
	return nil
}

var CommandDepends = cli.Command{
	Name:    "depends",
	Aliases: []string{"d"},
	Usage:   "Gets the dependencies needed for using the subpackages",
	Action:  deps,
}
