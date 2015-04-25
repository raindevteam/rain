# Globals for use in modules and listeners
global.__config = __dirname + '/config'
global.__models = __dirname + '/config/models'
global.__core   = __dirname + '/lib/core/core'
global.__debug  = __dirname + '/lib/debug'

irc      = require('irc')
core     = require(__core)
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
bot = new (irc.Client) 'irc.canternet.org', config.nick,
  userName: "RainbotDev"
  realName: "Rainbot"
  autoConnect: false
  port: 6667
  debug: true

init = ()->
  console.log 'here now --------=-=-----------------------=-=====================-=-=---------------------'
  if config.password
    bot.send 'ns', 'identify', config.password
  bot.send 'mode', config.nick, '+B'
  for channel in config.channels
    core.GATE channel

# Load Modules, Core and connect to IRC
core.preload()
core.loadModules ()->
  core.load(bot)
  core.listen () ->
    bot.connect init
