package rain

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/raindevteam/rain/hail"
	"github.com/raindevteam/rain/handler"
	"github.com/raindevteam/rain/internal"
	"github.com/raindevteam/rain/rbot"
	cli "gopkg.in/urfave/cli.v1"
)

// Run will start an instance of Rainbot.
func Run(ctx *cli.Context) {
	if ctx.NArg() != 1 {
		fmt.Println(ctx.Command.Usage)
		return
	}
	hail.Defaults()

	// Get Config.
	conffile := ctx.Args().First()
	conf, err := rbot.NewConfigFromFile(conffile)
	if err != nil {
		hail.Err(hail.Frain, err.Error())
		os.Exit(1)
	}

	// Create bot.
	bot, err := rbot.NewBot(os.Getenv("DG_TOKEN"), conf)
	if err != nil {
		hail.Errf(hail.Frain, "Unable to create bot: err: %s", err)
		os.Exit(1)
	}

	// Create and attach handler to bot.
	handler.CreateHandler()
	handler.Attach(bot)

	// Attach internals to bot.
	internal.Attach(bot)

	// Open discord session.
	bot.Connect()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

// NewCmdRun wraps Run() in a command struct.
func NewCmdRun(app *cli.App) {
	Command := cli.Command{
		Name:    "run",
		Aliases: []string{"r"},
		Usage:   "Usage: rain run config-file",
		Action:  Run,
	}
	app.Commands = append(app.Commands, Command)
}
