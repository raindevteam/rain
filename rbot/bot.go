// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package bot

import (
	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/hail"
)

// DiscordSession is an interface to specify discord session capability. It is
// implemented mostly for testing of the bot (so that mocks can be created for
// the discord session struct from bwmarrin's discordgo).
type DiscordSession interface {
	Open() (err error)
	AddHandler(handler interface{}) func()
}

// The Bot struct holds all information needed to do any general work on the
// bot. It also contains a DropletManager struct to manage modules at a high
// level.
type Bot struct {
	config  *Config
	Session DiscordSession
}

// NewBot will create a new instance of a bot, almost everything will be
// initialized, but some things such as connecting to discord will not. It'll
// be wise to read the documentation for the rest of this API.
func NewBot(conf *Config, token string) (*Bot, error) {
	var (
		ds  DiscordSession
		err error
	)

	if conf.testing {
		ds = &DST{0}
		err = nil
	} else {
		ds, err = discordgo.New("Bot " + token)
		hail.Defaults()
	}
	if err != nil {
		return nil, err
	}

	b := &Bot{
		config:  conf,
		Session: ds,
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

// AddHandler mocks discordgo.Session.AddHandler()
func (dst *DST) AddHandler(handler interface{}) func() {
	return func() {}
}
