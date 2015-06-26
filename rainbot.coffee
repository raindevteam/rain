# Globals for use in modules and listeners
global.__config    = __dirname + '/config/config'
global.__rainUtil  = __dirname + '/lib/rainlog'

# Core          = require './lib/core'
# global.__core = new Core()

bot      = require './lib/bot'
config   = require __config
# rainUtil = require __rainUtil
string   = require './lib/string'

# Set the level of debugging
#rainlog.setLoggingModes config.logging
#rainlog.setPrompt "RainBot "

# Create the bot
RainBot = new (bot) config.server, config.nick,
  userName: config.userName
  realName: config.realName
  autoConnect: false
  port: config.port
  debug: config.debug

start = ->
  require(__dirname + '/config/init')(bot)
  if config.password
    bot.send 'ns', 'identify', config.password
  bot.send 'mode', config.nick, '+B'
  __core.gate channel for channel in config.channels

RainBot.load ->
  RainBot.listen ->
    console.log 'done'
