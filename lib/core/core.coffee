# Requires
config  = require(__config)
defined = require('./defined')
_       = require('lodash')
async   = require('async')
module  = require('./module')
helpers = require('./helpers')
hashmap = require('hashmap')
fs      = require('fs')
response = require('./handlers/response')
responseHandler = new response()

# Locals
_bot    = undefined
commands = new hashmap()
triggers = {}
modules = []
alias   = undefined
cmd     = undefined
trigger = undefined

# All Events
exports.events = [
  'registered'
  'motd'
  'names'
  'names#channel'
  'topic'
  'join'
  'join#channel'
  'part'
  'part#channel'
  'quit'
  'kick'
  'kick#channel'
  'kill'
  'message'
  'message#'
  'message#channel'
  'notice'
  'ping'
  'pm'
  'ctcp'
  'ctcp-notice'
  'ctcp-privmsg'
  'ctcp-version'
  'nick'
  'invite'
  '+mode'
  '-mode'
  'whois'
  'channellist_start'
  'channelist_item'
  'channellist'
  'raw'
  'error'
]



# Export Handlers, helpers and defined messages
exports.helpers = helpers
exports.defined = defined
action = require('./handlers/action')
triggerHandler  = require('./handlers/trigger')
trigger         = new triggerHandler()
aliasHandler    = require('./handlers/alias')
alias           = new aliasHandler(trigger)
alias.loadAliases()
commandHandler  = require('./handlers/cmds')
cmd             = new commandHandler(alias, commands)
exports.trigger = trigger
exports.cmd     = cmd
exports.alias   = alias
actionHandler = undefined

# Modules
exports.module = module
exports.modules = modules
exports.addModule = (module) ->
  modules.push module

exports.preload = ->
  for event in exports.events
    triggers[event] = []

  ###
  fs.readdir __dirname + "../../modules", (err, modules) ->
    for module in modules
      if path.extname(module)
        continue # Don't load stray files
      module = require(__dirname + '/' + module)(core.module)
      for command in module.commands
         commands.set(command.name, command.value)
      for event, listeners in module.triggers
         triggers[event].push(listeners)
   ###
exports.loadModules = (callback) ->
  lmodule = require('./../../modules/rainbin/index')(exports.module)
  console.log(lmodule.addCommands())
  for command in lmodule.commands
    commands.set(command.name, command.command)
  for event, listeners of lmodule.triggers
    for listener in listeners
      triggers[event].push listener
  return callback()

# Load Bot
exports.load = (bot) ->
  Core = new module('Core')
  # Load commands
  fs.readdir __dirname + '/listeners/commands', (err, modfiles) ->
    for modfile in modfiles
      corecommand = require(__dirname + '/listeners/commands/' + modfile)
      Core.addCommands(corecommand)
    for command in Core.commands
      commands.set(command.name, command.command)
  # Load triggers
  fs.readdir __dirname + '/listeners/triggers', (err, modfiles) ->
    for modfile in modfiles
      coretrigger = require(__dirname + '/listeners/triggers/' + modfile)
      Core.addTriggers(coretrigger)
    for event, listeners of Core.triggers
      for listener in listeners
        triggers[event].push listener

  console.log(Core.commands)
  console.log(triggers)

  # Ready the bot
  _bot      = bot
  _bot.nicks = []
  _bot.sleep = false
  actionHandler = new action(responseHandler, cmd, _bot)


# Attach Listeners
exports.listen = (callback) ->
  _bot.addListener 'notice', (nick, to, text, msg) ->
    if !_bot.sleep
      console.log 'notice was sent ' + text
      actionHandler.setResponseProperties
         from: nick
         to: to
         text: text
         msg: msg
      # async.detect triggers['notice'],
      # actionHandler.triggered.bind(actionHandler),
      # actionHandler.fireTrigger.bind(actionHandler)

  ###
  # Names Events
  _bot.addListener 'names', (channel, nicks) ->
    for module in modules
      module.fire('names', [channel, nicks], exports.ACTION_RESPOND)

  # Join Events
  _bot.addListener 'join', (channel, nick, msg) ->
    for module in modules
      module.fire('join', [channel, nick, msg], exports.ACTION_RESPOND)

  # Nicks Events
  _bot.addListener 'nick', (oldnick, newnick, channels, msg) ->
    for module in modules
      module.fire('nick', [oldnick, newnick, channels, msg], exports.ACTION_RESPOND)
  ###

  # Message Events
  _bot.addListener 'message', (nick, to, text, msg) ->
    if !_bot.sleep
      console.log 'nick is ' + nick, 'to is ' + to
      actionHandler.setResponseProperties
        from: nick, to: to, text: text, msg: msg
      if cmd.isCommand(text)
        actionHandler.fireCommand text.after(defined.MSG_TRIGGER+1).clean()
      else
        console.log 'testing!'
        console.log triggers['message']
        async.detect triggers['message'],
        actionHandler.triggered.bind(actionHandler),
        actionHandler.fireTrigger.bind(actionHandler)

  # PM Events
  _bot.addListener 'pm', (nick, text, msgs) ->
    if !_bot.sleep
      for module in modules
        module.fire('pm', [nick, text, msgs], exports.ACTION_RESPOND)

  # Channellist Events
  _bot.addListener 'channellist', (channel_list) ->
    console.log channel_list

  # Error Events
  _bot.addListener 'error', (message) ->
    console.log message

  # Raw Events
  _bot.addListener 'raw', (message) ->
    if !_bot.sleep
      for module in modules
        module.fire('raw', [message], exports.ACTION_RESPOND)

  # Everythings ready! Return to caller!
  return callback()

exports.ACTION_RESPOND = (action) ->


exports.IDLE = () ->
  _bot.sleep = true
  _bot.addListener 'raw', wake = (message) ->
    if message['command'] == 'PRIVMSG'
      console.log 'here'
      if trigger.cmd(message['args'][1], 'wake') and
      _.includes(config.whitelist, message.nick)
        _bot.sleep = false
        _bot.removeListener 'raw', wake
        _bot.action message['args'][0], 'Gets up from idling'
        _bot.say message['args'][0], 'Systems online'

exports.CHANNEL_SWITCH = (JOIN, LEAVE, TO) ->
  if JOIN
    exports.GATE JOIN
    _bot.say TO, 'Joined ' + JOIN
  if LEAVE
    _bot.part LEAVE
    _bot.say TO, "Left " + LEAVE

exports.WHITELISTED = (func, nick) ->
  if _.includes(config.whitelisted_funcs, func)
    return _.includes(config.whitelist, nick)
  else return true

exports.GATE = (channel) ->
  _bot.addListener 'raw', gate = (message) ->
    if _bot.sleep and message['command'] != 'PRIVMSG'
      _bot.removeListener 'raw', gate
      _bot.sleep = false
    if message['args'][1]?.indexOf('Replaying up to') > -1
      _bot.sleep = true
  if channel then _bot.join channel
