package main

import (
	"os"

	"github.com/raindevteam/rain/rain"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rain"
	app.Usage = "Command line tool for the management of everything Rain"
	app.Version = rain.Version

	app.Commands = []cli.Command{
		rain.CommandInstall,
		rain.CommandModule,
		rain.CommandConfig,
		rain.CommandDepends,
	}

	app.Run(os.Args)
}
