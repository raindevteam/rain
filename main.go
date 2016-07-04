package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/raindevteam/rain/template"
	"github.com/urfave/cli"
)

const exConfig = `{
	"Host":    "irc.examplehost.net",
	"Port":    "6667",
	"Channel": ["#Your", "#Channels"],

	"Nick":     "MyBot",
	"RealName": "Wolfstein Jr. II",
	"UserName": "wolfstein",
	"Modes":    "+B",

	"CmdPrefix": "."
}`

func validLibraryOption(opt string) bool {
	switch opt {
	case
		"bot",
		"db",
		"rlog",
		"setup",
		"cli",
		"template":
		return true
	}
	return false
}

func validRmlOption(opt string) bool {
	switch opt {
	case
		"go",
		"js",
		"py":
		return true
	}
	return false
}

func validInstallOption(opt string) bool {
	switch opt {
	case
		"default",
		"neverfree",
		"simple":
		return true
	}
	return false
}

func main() {
	app := cli.NewApp()
	app.Name = "rainbot"
	app.Usage = "Command line tool for the management of everything rainbot"
	app.Version = "Beta 0.5.0"

	app.Commands = []cli.Command{
		{
			Name:    "make",
			Aliases: []string{"mk"},
			Usage:   "Creates a new bot with a given install and name",
			Description: "Creates a new bot by buidling it using the \"go build\" command.\n\n" +
				"   Possible build types are:\n    	default\n   	neverfree\n   	simple",
			ArgsUsage: "install name",
			Action:    createNewBot,
		},
		{
			Name:        "test",
			Aliases:     []string{"t"},
			Usage:       "Creates a template for either a test or module.",
			Description: "Creates a test template with the specified library to test",
			ArgsUsage:   "library filename",
			Action:      createTest,
			Hidden:      true,
		},
		{
			Name:        "mod",
			Aliases:     []string{"m"},
			Usage:       "Creates boilerplate for a module",
			Description: "Creates a template for a module tailored to the specified RML.",
			ArgsUsage:   "rml-prefix module-name",
			Action:      createModule,
		},
		{
			Name:        "config",
			Aliases:     []string{"c"},
			Usage:       "Creates a new JSON config",
			Description: "Will create a new config.json file in the current directory.",
			Action:      createConfig,
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
			fmt.Println(" No name specified")
			fmt.Println(" Usage: " + c.Command.Name + " <install> <name>")
			return nil
		}

		path := os.Getenv("GOPATH") + "/bin/" + c.Args().Get(1) + ".exe"

		output, err := exec.Command("go", "build", "-o", path,
			"github.com/raindevteam/rain/install/"+c.Args().First()).CombinedOutput()

		s := string(output[:])
		if err != nil {
			fmt.Println(" Internal command error >>>\n" + s)
		}
	} else {
		fmt.Println(" Not a valid install type. Valid install types are:\n - default\n - neverfree\n - simple")
	}

	return nil
}

func createTest(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if validLibraryOption(c.Args().First()) {
		if c.Args().Get(1) == "" {
			fmt.Println(" No filename specified")
			fmt.Println(" Usage: " + c.Command.Name + " library filename")
			return nil
		}

		tmpl.CreateTestTemplate(c.Args().Get(0), c.Args().Get(1))
	} else {
		// TODO: This is extremely inefficient and needs to be fixed!
		fmt.Println(" Not a valid library. Valid libraries are:\n - rlog\n - db\n - bot\n - template\n - cli\n - setup")
	}

	return nil
}

func createModule(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println(c.App.ArgsUsage)
	}

	if validRmlOption(c.Args().First()) {
		if c.Args().Get(1) == "" {
			fmt.Println(" No name specified")
			fmt.Println(" Usage: " + c.Command.Name + " rml-prefix module-name")
			return nil
		}

		tmpl.CreateModTemplate(c.Args().Get(0), c.Args().Get(1))
	} else {
		// TODO: Ok, seriously
		fmt.Println(" Not a valid install type. Valid install types are:\n   go\n   js\n   py")
	}

	return nil
}

func createConfig(c *cli.Context) error {
	fmt.Println(" Making an example config.json...")

	f, err := os.Create("config.json")
	if err != nil {
		fmt.Println(" Oops, something went wrong: " + err.Error())
		return nil
	}

	_, err = f.Write([]byte(exConfig))
	if err != nil {
		fmt.Println("Couldn't write config: " + err.Error())
		return nil
	}

	f.Close()
	fmt.Println(" All done!")
	return nil
}
