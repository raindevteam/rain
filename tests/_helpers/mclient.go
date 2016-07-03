package Thelpers

import (
	"strings"

	"github.com/RyanPrintup/nimbus"
	"github.com/wolfchase/rainbot/bot"
)

type MockClient struct {
	listeners map[string][]nimbus.Listener
	config    *rbot.Config
	quit      chan error
}

func NewMockClient(config *rbot.Config, msgs []string) *MockClient {
	mc := &MockClient{
		listeners: make(map[string][]nimbus.Listener),
		config:    config,
		quit:      make(chan error),
	}
	return mc
}

/****                                        Getters                                           ****/

func (mc *MockClient) GetNick() string {
	return mc.config.Nick
}

func (mc *MockClient) GetChannels() []string {
	return mc.config.Channel
}

/**************************************************************************************************/

/****                                          IRC                                             ****/

func (mc *MockClient) Connect() error {
	return nil
}

func (mc *MockClient) Listen() {
}

func (mc *MockClient) Quit() chan error {
	return mc.quit
}

func (mc *MockClient) Send(raw ...string) {
	msg, _ := nimbus.ParseMessage(strings.Join(raw, " "))
	mc.handleMsg(msg)
}

func (mc *MockClient) Say(channel string, text string) {
	mc.Send(nimbus.PRIVMSG, channel, text)
}

func (mc *MockClient) handleMsg(msg *nimbus.Message) {
	switch msg.Command {
	case "PRIVMSG":
		mc.Emit(msg.Command, msg)
	case "JOIN":
		raw := ":" + mc.config.Nick + "!fake.host" + " JOIN :" + msg.Middle
		mcmsg, _ := nimbus.ParseMessage(raw)
		mc.Emit(msg.Command, mcmsg)
	}
}

/**************************************************************************************************/

/****                                        Events                                            ****/

type Handle struct {
	listener nimbus.Listener
	done     chan bool
}

func (h *Handle) Run(msg *nimbus.Message) {
	h.done <- true
	h.listener(msg)
}

func (mc *MockClient) AddListener(event string, l nimbus.Listener) {
	mc.listeners[event] = append(mc.listeners[event], l)
}

func (mc *MockClient) GetListeners(event string) []nimbus.Listener {
	return mc.listeners[event]
}

func (mc *MockClient) Emit(event string, msg *nimbus.Message) {
	for _, listener := range mc.listeners[event] {
		h := Handle{listener, make(chan bool)}
		go h.Run(msg)
		<-h.done
	}
}

/**************************************************************************************************/
