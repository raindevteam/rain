package rain

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/raindevteam/rain/internal/hail"
	"github.com/raindevteam/rain/internal/handler"
	"github.com/raindevteam/rain/internal/internal"
	"github.com/raindevteam/rain/internal/rbot"
	cli "gopkg.in/urfave/cli.v1"
)

// Run will start an instance of Rainbot.
func Run(ctx *cli.Context) {
	if ctx.NArg() != 1 {
		fmt.Println(ctx.Command.Usage)
		// TODO: Set error code on exit.
		return
	}
	hail.Defaults()
	hail.Debug(hail.Frain, "If your are seeing this message than DEBUG is ON.")

	// Get Config.
	conffile := ctx.Args().First()
	conf, err := rbot.NewConfigFromFile(conffile)
	if err != nil {
		hail.Fatalf(hail.Frain, "Error reading config file: %s\n", err.Error())
		os.Exit(1)
	}

	// Create bot.
	bot, err := rbot.NewBot(os.Getenv("DG_TOKEN"), conf)
	hail.Info(hail.Frain, "Bot created.")
	if err != nil {
		hail.Errf(hail.Frain, "Unable to create bot: err: %s", err)
		os.Exit(1)
	}

	// Create and attach handler to bot.
	handler.CreateHandler()
	handler.Attach(bot)
	hail.Info(hail.Frain, "Handler created and attached.")

	// Attach internals to bot.
	internal.Attach(bot)
	hail.Info(hail.Frain, "Internals attached.")

	// Open discord session.
	bot.Connect()
	hail.Info(hail.Frain, "Discord connection initiated.")
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
