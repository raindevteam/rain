# Globals for use in modules and listeners
global.__config    = __dirname + '/config'
global.__models    = __dirname + '/config/models'
global.__rainUtil  = __dirname + '/lib/rainUtil'

Core          = require './lib/core'
global.__core = new Core()

irc      = require 'irc'
config   = require __config
rainUtil = require __rainUtil
string   = require './lib/string'

# Set the level of debugging
rainUtil.loggingLevel config.loggingLevel
rainUtil.setPrompt "RainBot "

# Create the bot
bot = new (irc.Client) config.server, config.nick,
  userName: "RainBotDev"
  realName: "RainBot"
  autoConnect: false
  port: 6667
  debug: true

# Assign bot to core
__core.setBot(bot)

init = ->
  if config.password
    bot.send 'ns', 'identify', config.password
  bot.send 'mode', config.nick, '+B'
  __core.gate channel for channel in config.channels

# Load Modules, Core and connect to IRC
__core.load () -> __core.listen () -> bot.connect init
