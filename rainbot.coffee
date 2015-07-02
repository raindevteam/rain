# RainBot 4.0 "Pyrelight"

# To encourage use of the rainlog utility class, we declare it here
# as a global for easy use within modules

global.rainlog  = require __dirname + '/lib/rainlog'
global.__config = require './config/config'

bot      = require './lib/bot/bot'
string   = require './lib/string'

# Set the level of debugging

#rainlog.setLoggingModes config.logging
#rainlog.setPrompt "RainBot "

# Instantiate a new bot with config settings

RainBot = new (bot) __config.server, __config.nick,
  userName: __config.userName
  realName: __config.realName
  autoConnect: false
  port: __config.port
  debug: __config.debug

RainBot.start()
