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
    module = require('./../../modules/rainbin/index')(Module)
    for command in module.commands
      @commands.set(command.name, command.command)
    for event, listeners of module.triggers
      for listener in listeners
        @triggers[event].push listener
    return callback()

  preload: () ->
    for event in @events
      @triggers[event] = []

    # Trigger Handler
    TriggerHandler = require('./handlers/trigger')
    @trigger        = new TriggerHandler()

    # Response Handler
    ResponseHandler = require('./handlers/response')
    @response        = new ResponseHandler()

    # Alias Handler
    AliasHandler    = require('./handlers/alias')
    @alias           = new AliasHandler(@trigger)
    @alias.loadAliases()

    # Command Handler
    CommandHandler  = require('./handlers/cmds')
    @command         = new CommandHandler(@alias, @commands)

    # Action Handler
    ActionHandler = require('./handlers/action')
    @action        = new ActionHandler(@response, @command, @bot)

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
        self.commands.set(command.name, command.command)

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

    @bot.addListener 'notice', (nick, to, text, msg) ->
      if !self.bot.sleep
        self.action.setResponseProperties
           from: nick, to: to, text: text, msg: msg
        async.detect self.triggers['notice'],
        self.action.triggered.bind(self.action),
        self.action.fireTrigger.bind(self.action)

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
    @bot.addListener 'message', (nick, to, text, msg) ->
      if !self.bot.sleep
        self.action.setResponseProperties
          from: nick, to: to, text: text, msg: msg
        if self.command.isCommand(text)
          self.action.fireCommand text.after(defined.MSG_TRIGGER+1).clean()
        else
          async.detect self.triggers['message'],
          self.action.triggered.bind(self.action),
          self.action.fireTrigger.bind(self.action)

    # PM Events
    @bot.addListener 'pm', (nick, text, msgs) ->
      if !self.bot.sleep then return

    # Channellist Events
    @bot.addListener 'channellist', (channel_list) ->
      console.log channel_list

    # Error Events
    @bot.addListener 'error', (message) ->
      console.log message

    # Raw Events
    @bot.addListener 'raw', (message) ->
      if !self.bot.sleep then return

    # Everythings ready! Return to caller!
    return callback()

  idle: () ->
    @bot.sleep = true
    @bot.addListener 'raw', wake = (message) ->
      if message['command'] == 'PRIVMSG'
        if @trigger.cmd(message['args'][1], 'wake') and
        _.includes(config.whitelist, message.nick)
          @bot.sleep = false
          @bot.removeListener 'raw', wake
          @bot.action message['args'][0], 'Gets up from idling'
          @bot.say message['args'][0], 'Systems online'

  channelSwitch: (JOIN, LEAVE, TO) ->
    if JOIN
      @gate JOIN
      _bot.say TO, 'Joined ' + JOIN
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
