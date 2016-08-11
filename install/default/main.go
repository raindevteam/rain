package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/raindevteam/rain/rain"
	"github.com/raindevteam/rain/rbot"
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

	bot := rbot.NewBot(rain.Version(), rconf)
	setup.Default(bot)

	bot.EnableModules()
	bot.DefaultConnectWithMsg("Connecting...", "Done")

	reason, err := bot.WaitForQuit()
	if err != nil {
		fmt.Printf("Quit Error: %s\n", reason)
		os.Exit(1)
	}

	fmt.Printf("Quit: %s\n", reason)
}
