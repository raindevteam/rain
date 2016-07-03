package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wolfchase/rainbot/bot"
	"github.com/wolfchase/rainbot/cli"
	"github.com/wolfchase/rainbot/setup"
)

var i = flag.Bool("i", false, "run the bot in interactive mode (does not connect to IRC)")

func main() {
	flag.Usage = func() {
		program := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
		fmt.Println("Usage: " + program + " [-i | -h] config.json")
		flag.PrintDefaults()
	}

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	flag.Parse()

	rconf, err := rbot.ReadConfig(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var preConnectMsg, postConnectMsg string
	var bot *rbot.Bot
	if *i {
		preConnectMsg = ""
		postConnectMsg = ""
		bot = clibot.NewCLIBot(rconf)
	} else {
		preConnectMsg = "Connecting... "
		postConnectMsg = "Done"
		bot = rbot.NewBot("0.4.0 (Monterey Jack)", rconf)
	}

	setup.Default(bot)

	bot.EnableModules(rconf)
	bot.DefaultConnectWithMsg(preConnectMsg, postConnectMsg)
}
