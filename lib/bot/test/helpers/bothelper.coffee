bot = require './../../bot'
events = require './../../events'
global.__config = require './../../../../config/config'

module.exports =

    makeTestBot: () ->
      testbot = new (bot) __config.server, __config.nick,
        autoConnect: false
      return testbot
