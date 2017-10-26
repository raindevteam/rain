// Copyright 2015-2017, Rodolfo Castillo-Valladares. All rights reserved.
//
// This file is governed by the Modified BSD License. You should have
// received a copy of the license (LICENSE.md) with this file's program.
// You may find the program here: https://github.com/raindevteam/rain
//
// Contact Rodolfo at rcvallada@gmail.com for any inquiries of this file.

package rbot

import (
	"errors"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/raindevteam/rain/internal/hail"
)

// DiscordSession is an interface to specify discord session capability. It is
// implemented mostly for testing of the bot (so that mocks can be created for
// the discord session struct from bwmarrin's discordgo).
type DiscordSession interface {
	AddHandler(handler interface{}) func()
	Open() (err error)
}

// The Bot struct holds all information needed to do any general work on the
// bot. It also contains a DropletManager struct to manage modules at a high
// level.
type Bot struct {
	config  *Config
	Session DiscordSession
}

// NewBot will create a new instance of a bot, almost everything will be
// initialized, but some things such as connecting to discord will not.
func NewBot(token string, conf *Config) (*Bot, error) {
	var (
		ds  DiscordSession
		err error
	)

	if conf.testing {
		ds = &DST{0}
		err = nil
	} else {
		// TODO: Test this section of code by verifying that
		//       discordgo.New() sets ds and err accordingly.
		hail.Defaults()
		if token == "" {
			hail.Fatalf(hail.Fbot, "Token is empty, exiting...")
			os.Exit(1)
		}
		ds, err = discordgo.New("Bot " + token)
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

// Connect will call Open() on the discord session.
func (b *Bot) Connect() error {
	err := b.Session.Open()
	if err != nil {
		return err
	}
	return nil
}

///////////////////               Testing                    ///////////////////

// DST stands for DiscordSessionTest. It has mocked methods that are used within
// the bot's functions. The 'id' int for each struct is used to identify the
// structs. Positive IDs, starting with 0, represent properly working mocks.
// Negative IDs, starting with -1, represent intentionally broken mocks.
type DST struct {
	ID int
}

// AddHandler mocks discordgo.Session.AddHandler()
func (dst *DST) AddHandler(handler interface{}) func() {
	return func() {}
}

// Open mocks discordgo.Session.Open().
func (dst *DST) Open() (err error) {
	if dst.ID < 0 {
		return errors.New("can't open discord session (false alarm from DST)")
	}
	return nil
}
