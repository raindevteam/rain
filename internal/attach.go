package internal 

import "github.com/raindevteam/rain/rbot"

func Attach(bot *rbot.Bot) {
	ls := Listeners{bot}
	ls.AddMessageCreate()
}
