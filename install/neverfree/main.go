package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/raindevteam/rain/cbot"
	"github.com/raindevteam/rain/rain"
	"github.com/raindevteam/rain/rbot"
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
	rconf, err := rbot.ReadConfigFile(os.Args[confarg])
	if err != nil {
		panic(err)
	}

	var (
		preConnectMsg, postConnectMsg string
		bot                           *rbot.Bot
	)

	if *i {
		preConnectMsg, postConnectMsg = "", ""
		bot = cbot.NewCLIBot(rconf)
	} else {
		preConnectMsg, postConnectMsg = "Connecting... ", "Done"
		bot = rbot.NewBot(rain.Version, rconf)
	}

	setup.Default(bot)
	bot.EnableModules()
	bot.DefaultConnectWithMsg(preConnectMsg, postConnectMsg)

	result := <-bot.Quit()
	if result != nil {
		fmt.Println(result.Error())
	}
}
