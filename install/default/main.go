package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/raindevteam/rain/bot"
	"github.com/raindevteam/rain/setup"
)

func main() {
	if len(os.Args) < 2 {
		program := strings.TrimSuffix(filepath.Base(os.Args[0]), ".exe")
		fmt.Println("Usage: " + program + " config.json")
		os.Exit(1)
	}

	rconf, err := rbot.ReadConfig(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bot := rbot.NewBot("0.5.0", rconf)
	setup.Default(bot)

	bot.EnableModules(rconf)
	bot.DefaultConnectWithMsg("Connecting...", "Done")
}
