package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/urfave/cli"
)

func validLibraryOption(opt string) bool {
	switch opt {
	case
		"bot",
		"db",
		"template":
		return true
	}
	return false
}

func validInstallOption(opt string) bool {
	switch opt {
	case
		"default",
		"simple":
		return true
	}
	return false
}

func main() {
	app := cli.NewApp()
	app.Name = "rainbot"
	app.Usage = "Command line tool for the management of everything rainbot"
	app.Version = "0.4.0 (Monterey Jack)"

	app.Commands = []cli.Command{
		{
			Name:      "make",
			Aliases:   []string{"mk"},
			Usage:     "Creates a bot",
			UsageText: "new - Creates a new bot with given name and type. Possible types are: default, simple",
			Action:    createNewBot,
		},
		{
			Name:      "template",
			Aliases:   []string{"t", "templ"},
			Usage:     "template <type>",
			UsageText: "template - Create a template for either a test or module.",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
		{
			Name: "mkconfig",
		},
	}

	app.Run(os.Args)
}

func createNewBot(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if validInstallOption(c.Args().First()) {
		if c.Args().Get(1) == "" {
			fmt.Println("No name specified")
			return nil
		}

		path := os.Getenv("GOPATH") + "/bin/" + c.Args().Get(1) + ".exe"

		output, err := exec.Command("go", "build", "-o", path,
			"github.com/wolfchase/rainbot/install/"+c.Args().First()).CombinedOutput()

		s := string(output[:])
		if err != nil {
			fmt.Println("Internal command error: " + s)
		}
	} else {
		fmt.Println("Not a valid install type")
	}

	return nil
}
