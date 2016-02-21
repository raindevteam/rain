package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/wolfchase/rainbot/bot"
	"github.com/wolfchase/rainbot/cli"
	"github.com/wolfchase/rainbot/template"

	"github.com/RyanPrintup/nimbus"
)

var t = flag.String("t", "", "template type to create (mod or test)")
var i = flag.Bool("i", false, "run the bot in interactive mode (does not connect to IRC)")

func ValidLibraryOpt(opt string) bool {
	switch opt {
	case
		"bot",
		"db",
		"template":
		return true
	}
	return false
}

func HandleFlags() {
	if *t != "" {
		switch *t {
		case "test":
			if len(flag.Args()) != 2 {
				fmt.Println("RainBot: Invalid number of arguments")
				fmt.Println("RainBot: Valid Args for (-t test): lib name")
				os.Exit(1)
			}

			if !ValidLibraryOpt(flag.Args()[0]) {
				fmt.Println("RainBot: " + flag.Args()[0] + " not a testable library")
				os.Exit(1)
			}

			tmpl.CreateTestTemplate(flag.Args()[0], flag.Args()[1])
		}
		os.Exit(0)
	}

}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: rainbot [-t | -m] | config")
		os.Exit(1)
	}

	flag.Parse()
	HandleFlags()

	var bot *rainbot.Bot
	var rainConfig *rainbot.Config
	var nimConfig *nimbus.Config
	var err error

	if *i {
		rainConfig, err = rainbot.ReadConfig(os.Args[2])
		fmt.Println(os.Args[2])

		fmt.Println(rainConfig.Host)

		if err != nil {
			fmt.Println(err)
			fmt.Println("here?")
			os.Exit(1)
		}

		bot = clibot.NewCLIBot(rainConfig)
	} else {
		nimConfig, rainConfig, err = rainbot.GetConfigs(os.Args[1])

		fmt.Println(os.Args[1])

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		bot = &rainbot.Bot{
			/* Client      */ nimbus.NewClient(rainConfig.Host, rainConfig.Nick, *nimConfig),
			/* Version     */ "Alpha 0.4.0 (Monterey Jack)",
			/* Modules     */ make(map[string]*rainbot.Module),
			/* Channels    */ make(map[string]*rainbot.Channel),
			/* Parser      */ rainbot.NewParser(rainConfig.CmdPrefix),
			/* Handler     */ rainbot.NewHandler(),
			/* Mutex       */ sync.Mutex{},
		}
	}

	bot.SetupModules(rainConfig)

	fmt.Print("Connecting... ")

	bot.Connect(func(e error) {
		if e != nil {
			fmt.Println(e)
			return
		}

		fmt.Println("Done")

		bot.LoadModules()

		bot.Handler.AddInternalCommand("m", &rainbot.Command{
			Help: "The RainBot Module Manager helps manage modules.",
			Fun: func(msg *nimbus.Message, args []string) {
				if len(args) == 0 {
					bot.Say(msg.Args[0], "Error: No command specified")
					return
				}
				switch args[0] {
				case "reload":
					if len(args) == 1 {
						bot.Say(msg.Args[0], "Error: No module specified")
						return
					}

					err := bot.ModuleReload(args[1])
					if err != nil {
						bot.Say(msg.Args[0], "Error: Could not reload module, view output in console")
						return
					}
					bot.Say(msg.Args[0], "Reloaded module: "+args[1])
				case "kill":
					if len(args) == 1 {
						bot.Say(msg.Args[0], "Error: No module specified")
						return
					}

				}

			},
			PM: true,
			CM: true,
		})

		// Commands listener
		bot.AddListener(nimbus.PRIVMSG, func(msg *nimbus.Message) {
			if bot.Parser.IsCommand(msg.Trailing) {
				command, args := bot.Parser.ParseCommand(msg.Trailing)
				bot.Handler.Invoke(msg, rainbot.CommandName(command), args)
			}
		})

		bot.AddListener(nimbus.PRIVMSG, func(msg *nimbus.Message) {
			text := msg.Trailing
			if text == "Hello, "+bot.GetNick() {
				bot.Client.Say(msg.Args[0], "Hello there!")
			}
		})

		// Add/Update channel when topic is received
		bot.AddListener(nimbus.RPL_TOPIC, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			name := msg.Args[1]
			topic := msg.Trailing

			bot.Channels[strings.ToLower(name)].Topic = topic

			bot.Mu.Unlock()
		})

		// Update users for channel
		bot.AddListener(nimbus.RPL_NAMEREPLY, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			channel := bot.Channels[strings.ToLower(msg.Args[2])]
			users := strings.Split(strings.Trim(msg.Trailing, " "), " ")

			for _, user := range users {
				var name, rank string

				if strings.ContainsAny(string(user[0]), "+ & ~ & @ & &") {
					name, rank = user[1:], string(user[0])
				} else {
					name, rank = user, ""
				}

				channel.Users[name] = rank
			}

			bot.Mu.Unlock()
		})

		// Update on user Join
		bot.AddListener(nimbus.JOIN, func(msg *nimbus.Message) {
			bot.Mu.Lock()
			defer bot.Mu.Unlock()

			who, _ := bot.Parser.ParsePrefix(msg.Prefix)
			where := msg.Args[0]

			if who == bot.GetNick() {
				channel := rainbot.NewChannel(where)
				bot.Channels[strings.ToLower(where)] = channel
				return
			}

			channel := bot.Channels[strings.ToLower(where)]
			channel.Users[who] = ""
		})

		// Update on user Kick
		bot.AddListener(nimbus.KICK, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			who := msg.Args[2]
			where := msg.Args[0]

			bot.RemoveUser(who, where)

			bot.Mu.Unlock()
		})

		// Update on user Kill
		bot.AddListener(nimbus.KILL, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			// Implement getInfo(msg) function?
			who, _ := bot.Parser.ParsePrefix(msg.Prefix)
			where := msg.Args[0]

			bot.RemoveUser(who, where)

			bot.Mu.Unlock()
		})

		// Update on user part
		bot.AddListener(nimbus.PART, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			who, _ := bot.Parser.ParsePrefix(msg.Prefix)
			where := msg.Args[0]

			bot.RemoveUser(who, where)

			bot.Mu.Unlock()
		})

		// Update on user quit
		bot.AddListener(nimbus.QUIT, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			who, _ := bot.Parser.ParsePrefix(msg.Prefix)

			if who == bot.GetNick() {
				// We quit
			}

			for _, channel := range bot.Channels {
				delete(channel.Users, who)
			}

			bot.Mu.Unlock()
		})

		// Update on nick change
		bot.AddListener(nimbus.NICK, func(msg *nimbus.Message) {
			bot.Mu.Lock()

			fmt.Println(msg)

			bot.Mu.Unlock()
		})

		bot.Listen()
		result := <-bot.Quit()

		fmt.Println(result)
	})
}
