package rain

import (
	"fmt"

	"github.com/raindevteam/rain/tmpl"
	"github.com/urfave/cli"
)

var rmls = []string{
	"go",
	"py",
	"js",
}

func module(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if !isIn(rmls, c.Args().First()) {
		fmt.Printf(" Not a valid install type. Valid install types are:\n%s", listPrint(rmls))
	} else {
		if c.Args().Get(1) == "" {
			fmt.Println(" No name specified")
			fmt.Println(" Usage: " + c.Command.Name + " rml-prefix module-name")
			return nil
		}

		tmpl.CreateModTemplate(c.Args().Get(0), c.Args().Get(1))
		fmt.Printf(" Created module template: %s.%s", c.Args().Get(1), c.Args().Get(0))
	}

	return nil
}

var CommandModule = cli.Command{
	Name:        "module",
	Aliases:     []string{"m"},
	Usage:       "Creates boilerplate for a module",
	Description: "Creates boilerplate for a module tailored to the specified RML.",
	ArgsUsage:   "rml-prefix module-name",
	Action:      module,
}
