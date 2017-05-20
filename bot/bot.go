package bot

import (
	"github.com/bwmarrin/discordgo"
)

// DiscordSession is an interface to specify discord session capability. It is
// implemented mostly for testing of the bot (so that mocks can be created for
// the discord session struct from bwmarrin's discordgo).
type DiscordSession interface {
	Open() (err error)
}

// The Bot struct holds all information needed to do any general work on the
// bot. It also contains a DropletManager struct to manage modules at a high
// level.
type Bot struct {
	name string
	ds   DiscordSession
}

// NewBot will create a new instance of a bot, almost everything will be
// initialized, but some things such as connecting to discord will not. It'll
// be wise to read the documentation for the rest of this API.
func NewBot(conf *Config, name, token string) (*Bot, error) {
	var (
		ds  DiscordSession
		err error
	)

	if conf.testing {
		ds = &DST{0}
		err = nil
	} else {
		ds, err = discordgo.New("Bot " + token)
	}
	if err != nil {
		return nil, err
	}

	b := &Bot{
		name: name,
		ds:   ds,
	}
	return b, nil
}

///////////////////               Testing                    ///////////////////

// DST stands for DiscordSessionTest. It has mocked methods that are used within
// the bot's functions. The 'id' int for each struct is used to identify the
// structs. The identifer 0 is special and means it is the primary (first
// ever instanced) test session when testing.
type DST struct {
	id int
}

// Open mocks discordgo.Session.Open().
func (dst *DST) Open() (err error) {
	return nil
}
