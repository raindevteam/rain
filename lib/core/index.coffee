config  = require(__config)
_       = require('lodash')
async   = require('async')
Hashmap = require('hashmap')
fs      = require('fs')

defined = require('./lib/defined')
Module  = require('./lib/module')
helpers = require('./helpers')
events  = require('./lib/events')

class Core
  constructor: () ->
    @commands = new Hashmap()
    @triggers = {}
    @modules  = []
    @defined  = defined
    @helpers  = helpers
    @action   = undefined
    @response = undefined
    @alias    = undefined
    @command  = undefined
    @trigger  = undefined
    @events   = events
    @Module   = Module

  setBot: (bot) ->
    # Ready the bot
    @bot       = bot
    @bot.nicks = []
    @bot.sleep = false

  loadModules: (callback) ->
    self = @
    fs.readdir __dirname + '/../../modules', (err, modules) ->
      for module in modules
        module = require(__dirname + '/../../modules/' + module)(self.Module)
        for command in module.commands
          self.commands.set(command.name.lower(), command.command)
        for event, listeners of module.triggers
          for listener in listeners
            self.triggers[event].push listener
      return callback()

  getNewActionHandler: () ->
    return new @ActionHandler(@ResponseHandler, @command)

  preload: () ->
    for event in @events
      @triggers[event] = []

    # Trigger Handler
    TriggerHandler = require('./handlers/trigger')
    @trigger        = new TriggerHandler()

    # Response Handler
    @ResponseHandler = require('./handlers/response')(@bot)

    # Alias Handler
    AliasHandler    = require('./handlers/alias')
    @alias           = new AliasHandler(@trigger)
    @alias.loadAliases()

    # Command Handler
    CommandHandler  = require('./handlers/cmds')
    @command         = new CommandHandler(@alias, @commands)

    # Action Handler
    @ActionHandler = require('./handlers/action')

  load: (callback) ->
    self = @

    @preload()

    # Create Core
    Core = new @Module('Core')

    # Load core commands
    fs.readdir __dirname + '/listeners/commands', (err, modfiles) ->
      for modfile in modfiles
        command = require(__dirname + '/listeners/commands/' + modfile)
        Core.addCommands(command)
      for command in Core.commands
        self.commands.set(command.name.lower(), command.command)

    # Load core triggers
    fs.readdir __dirname + '/listeners/triggers', (err, modfiles) ->
      for modfile in modfiles
        trigger = require(__dirname + '/listeners/triggers/' + modfile)
        Core.addTriggers(trigger)
      for event, listeners of Core.triggers
        for listener in listeners
          self.triggers[event].push listener

    @loadModules ->
      return callback()

  listen: (callback) ->
    self = @

    # Notice
    @bot.addListener 'notice', (nick, to, text, msg) ->
      action = self.getNewActionHandler()
      if !self.bot.sleep
        action.setResponseProperties from: nick, to: to, text: text, msg: msg
        async.each self.triggers['notice'],
        action.handle.bind(action), (err) ->
          if err then console.error 'Problem firing notice triggers'

    # Names Events
    @bot.addListener 'names', (channel, nicks) ->
      action = self.getNewActionHandler()
      if !self.bot.sleep
        action.setResponseProperties channel: channel, nicks: nicks
        async.each self.triggers['names'],
        action.handle.bind(action), (err) ->
          if err then console.error 'Problem firing names triggers'


    # Join Events
    @bot.addListener 'join', (channel, nick, msg) ->
      action = self.getNewActionHandler()
      action.setResponseProperties channel: channel, nick: nick, msg: msg
      async.each self.triggers['join'],
      action.handle.bind(action), (err) ->
        if err then console.error 'Problem firing names triggers'

    @bot.addListener 'part', (channel, nick, reason, msg) ->
      console.log 'parting for ' + nick + ' from ' + channel
      action = self.getNewActionHandler()
      action.setResponseProperties
        channel: channel, nick: nick, reason: reason, msg: msg
      async.each self.triggers['part'],
      action.handle.bind(action), (err) ->
        if err then console.error 'Problem firing names triggers'

    @bot.addListener 'quit', (nick, reason, channels, msg) ->
      action = self.getNewActionHandler()
      action.setResponseProperties
        nick: nick, reason: reason, channels: channels, msg: msg
      async.each self.triggers['quit'],
      action.handle.bind(action), (err) ->
        if err then console.error 'Problem firing names triggers'

    # Nicks Events
    @bot.addListener 'nick', (oldnick, newnick, channels, msg) ->
      action = self.getNewActionHandler()
      action.setResponseProperties
        oldnick: oldnick, newnick: newnick, channels: channels, msg: msg
      async.each self.triggers['nick'],
      action.handle.bind(action), (err) ->
        if err then console.error 'Problem firing names triggers'

    # Message Events
    @bot.addListener 'message', (nick, to, text, msg) ->
      action = self.getNewActionHandler()
      if !self.bot.sleep
        action.setResponseProperties
          from: nick, to: to, text: text, msg: msg
        if self.command.isCommand(text)
          action.fireCommand text.after(defined.MSG_TRIGGER+1).clean()
          async.each self.triggers['message'],
          action.onCommandTrigger.bind(action)
        else
          async.each self.triggers['message'],
          action.handle.bind(action), (err) ->

    @bot.addListener 'action', (from, to, text, msg) ->
      action = self.getNewActionHandler()
      action.setResponseProperties
        from: from, to: to, text: text, msg: msg
      async.each self.triggers['action'],
      action.handle.bind(action), (err) ->
        if err then console.error 'Problem firing names triggers'

    # PM Events
    @bot.addListener 'pm', (nick, text, msg) ->
      action = self.getNewActionHandler()
      if !self.bot.sleep
        action.setResponseProperties
          from: nick, to: config.nick, text: text, msg: msg
        action.fireCommand text.clean()
        async.detect self.triggers['message'],
        action.onCommandTrigger.bind(action), action.fireTrigger.bind(action)

    # Channellist Events
    @bot.addListener 'channellist', (channel_list) ->

    # Error Events
    @bot.addListener 'error', (message) ->
      console.log "from error " + message

    # Raw Events
    @bot.addListener 'raw', (message) ->

    # Everythings ready! Return to caller!
    return callback()

  idle: () ->
    self = @
    @bot.sleep = true
    @bot.addListener 'raw', wake = (message) ->
      if message['command'] == 'PRIVMSG'
        if self.trigger.cmd(message['args'][1], 'wake') and
        _.includes(config.whitelist, message.nick)
          self.bot.sleep = false
          self.bot.removeListener 'raw', wake
          self.bot.action message['args'][0], 'Gets up from idling'
          self.bot.say message['args'][0], 'Systems online'

  channelSwitch: (JOIN, LEAVE, TO) ->
    if JOIN
      @gate JOIN
      @bot.say TO, 'Joined ' + JOIN
    if LEAVE
      @bot.part LEAVE
      @bot.say TO, "Left " + LEAVE

  checkWhitelist: (func, nick) ->
    if _.includes(config.whitelisted_funcs, func)
      return _.includes(config.whitelist, nick)
    else return true

  gate: (channel) ->
    self = @
    @bot.addListener 'raw', gate = (message) ->
      if self.bot.sleep and message['command'] != 'PRIVMSG'
        self.bot.removeListener 'raw', gate
        self.bot.sleep = false
      if message['args'][1]?.indexOf('Replaying up to') > -1
        self.bot.sleep = true
    if channel then @bot.join channel

module.exports = Core
