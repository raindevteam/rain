package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/raindevteam/rain/bot"
	"github.com/raindevteam/rain/cli"
	"github.com/raindevteam/rain/setup"
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

	confarg := 1
	if *i {
		confarg = 2
	}
	rconf, err := rbot.ReadConfig(os.Args[confarg])
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
		bot = rbot.NewBot("Alpha 0.5.0 (Stable 91)", rconf)
	}

	setup.Default(bot)

	bot.EnableModules(rconf)
	bot.DefaultConnectWithMsg(preConnectMsg, postConnectMsg)
}
