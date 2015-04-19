# Requires
config  = require(__config)
defined = require('./defined')
_       = require('lodash')
async   = require('async')
module  = require('./module')
helpers = require('./helpers')

# Locals
_bot    = undefined
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
triggerHandler  = require('./handlers/trigger')
trigger         = new triggerHandler()
aliasHandler    = require('./handlers/alias')
alias           = new aliasHandler(trigger)
alias.loadAliases()
commandHandler  = require('./handlers/cmds')
cmd             = new commandHandler(alias)
exports.trigger = trigger
exports.cmd     = cmd
exports.alias   = alias

# Modules
exports.module = module
exports.modules = modules
exports.addModule = (module) ->
  modules.push module

# Load Bot
exports.load = (bot) ->

  Core = new module('Core')
  Core.addListeners(require('./listeners/notice'))
  Core.addListeners(require('./listeners/names'))
  Core.addListeners(require('./listeners/join'))
  Core.addListeners(require('./listeners/part'))
  Core.addListeners(require('./listeners/msg'))
  Core.addListeners(require('./listeners/raw'))
  modules.push(Core)

  # Ready the bot
  _bot      = bot
  bot.nicks = []
  bot.sleep = false


# Attach Listeners
exports.listen = (callback) ->
  _bot.addListener 'notice', (nick, to, text, msgs) ->
    if !_bot.sleep
      for module in modules
        module.fire('notice', [nick, to, text, msgs], exports.ACTION_RESPOND)

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

  # Message Events
  _bot.addListener 'message', (nick, to, text, msgs) ->
    if !_bot.sleep
      if cmd.isCommand(text)
        trigger.acceptingCommands(true)
        cmd.handle(nick, to, text.after(defined.MSG_TRIGGER+1).trim(), msgs)
      else
        trigger.acceptingCommands(false)
        for module in modules
          module.fire('message', [nick, to, text, msgs], exports.ACTION_RESPOND)

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
  if (action.responses)
    for response in action.responses
      if (response.length > 250) then return # temp
      _bot[response.method](response.to, response.res)

exports.IDLE = () ->
  _bot.sleep = true
  _bot.addListener 'raw', wake = (message) ->
    if message['command'] == 'PRIVMSG'
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
