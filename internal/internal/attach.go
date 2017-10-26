package internal

import "github.com/raindevteam/rain/internal/rbot"

// Attach will insert all internal listeners and commands into the handler. It
// also requires an instance of a bot too pass onto the listener and command
// structs.
func Attach(bot *rbot.Bot) {
	ls := Listeners{bot}
	ls.AddMessageCreate()
}
