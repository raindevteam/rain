package rain

import (
	"fmt"
	"os"

	"github.com/raindevteam/rain/rbot"
	"github.com/urfave/cli"
)

func config(c *cli.Context) error {
	fmt.Println(" Making an example config.yaml...")

	f, err := os.Create("config.yaml")
	if err != nil {
		fmt.Println(" Oops, something went wrong: " + err.Error())
		return err
	}

	_, err = f.Write([]byte(rbot.ExampleConfig))
	if err != nil {
		fmt.Println(" Couldn't write config: " + err.Error())
		return err
	}

	f.Close()
	fmt.Println(" All done!")
	return nil
}

var CommandConfig = cli.Command{
	Name:        "config",
	Aliases:     []string{"c"},
	Usage:       "Creates a new JSON config",
	Description: "Will create a new config.json file in the current directory.",
	Action:      config,
}
