# Globals for use in modules and listeners
global.__config = __dirname + '/config'
global.__models = __dirname + '/config/models'
global.__debug  = __dirname + '/lib/debug'

Core          = require('./lib/core')
global.__core = new Core()

irc      = require('irc')
config   = require(__config)
debug    = require(__debug)
string   = require('./lib/string')
mongoose = require('mongoose')

# Set the level of debugging
debug.set(config.debug)

# Connect to the database
if config.db
  mongoose.connect(config.db)

# Create the bot
bot = new (irc.Client) config.server, config.nick,
  userName: "RainbotDev"
  realName: "Rainbot"
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
