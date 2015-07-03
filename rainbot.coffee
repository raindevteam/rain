# RainBot 4.0 "Pyrelight"
# An extensible, multipurpose IRC bot written in CoffeeScript with node.js

global.rainlog  = require __dirname + '/lib/rainlog'
global.__config = require './config/config'
global.X        = require './lib/weave'

bot      = require './lib/bot/bot'
string   = require './lib/string'

# Set the level of debugging

#rainlog.setLoggingModes config.logging
#rainlog.setPrompt "RainBot "

RainBot = new (bot) __config.server, __config.nick,
  userName: __config.userName
  realName: __config.realName
  autoConnect: false
  port: __config.port
  debug: __config.debug

RainBot.start()
