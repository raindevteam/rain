irc = require 'irc'
fs = require 'fs'

HookHandler = require './hookhandler'
RespondQueue = require './responders/respondqueue'
alias = require './alias'
ircHelpers = require './irc'

MODULES_DIRECTORY = __dirname + '/../../modules/'

class bot extends irc.Client

  # bot :: constructor (String, String, options Object)
  #
  # The constructor passes itself over to the RespondQueue which
  # uses it to send out responses from the responders. The irc.Client
  # superclass is also instantiated.

  constructor: (server, nick, options) ->
    rainlog.info 'Bot', 'Initializing bot...'
    @version = '0.4.0 (Pyrelight)'
    @Module = require('./module')(@)
    @alias = alias
    @irc = ircHelpers(@)
    @modules = []
    @config = __config
    @sleep = false
    alias.loadAliases()
    RespondQueue.setBot(@)
    super server, nick, options

  # bot :: loadModules (Callback)
  # Loads up the modules

  loadModules: (callback) ->
    rainlog.info 'Bot', 'Loading modules...'
    self = @
    fs.readdir MODULES_DIRECTORY, (err, modules) ->
      for moduleDir in modules
        if !fs.lstatSync(MODULES_DIRECTORY + moduleDir).isDirectory()
          rainlog.warn 'Bot', moduleDir + ' is not a module directory'
          continue
        module = require(MODULES_DIRECTORY + moduleDir)(self.Module)
        if !(module instanceof self.Module)
          rainlog.err 'Bot', moduleDir + ' is not a module'
          rainlog.err 'Bot', 'Make sure that module exports a Module instance'
          continue
        self.modules.push(module)
        HookHandler.extractHooks(module)
        rainlog.info 'Bot', 'Loaded module: ' + module.name
      return callback()

  prestart: (callback) ->
    self = @
    self.loadModules -> self.attachHooks callback

  start: () ->
    self = @
    rainlog.info 'Bot', 'Starting bot...'
    rainlog.info 'Bot', 'Temporarily caching nsPassword and Pastebin API'
    if @config.nsPassword then nsPassword = @config.nsPassword
    if @config.pastebinApi then pastebinApi = @config.pastebinApi
    rainlog.info 'Bot', 'Unsetting nsPassword and Pastebin API in config'
    @config.nsPassword = ''
    @config.pastebinApi = ''
    @prestart ->
      require('./../../config/init') @, ->
        self.connect ->
          rainlog.info 'Bot', 'Bot connected to IRC'
          if __config.nsPassword
            self.send 'ns', 'identify', nsPassword
          self.send 'mode', __config.nick, '+B'
          rainlog.info 'Bot', 'Connecting to channels'
          self.gate channel for channel in __config.channels

  # bot :: dispatch (String, params Object)
  # Dispatches events to modules

  dispatch: (event, params) ->
    self = @
    if self.sleep then return
    HookHandler.fire(event, params)

  # bot :: attachHooks (Callback)
  # Attaches listeners to the bot which dispatch events to modules

  attachHooks: (callback) ->
    self = @

    # Notice Events

    @.addListener 'notice', (nick, to, text, msg) ->
      self.dispatch 'notice', from: nick, to: to, text: text, msg: msg

    # Message Events

    @.addListener 'message', (nick, to, text, msg) ->
      self.dispatch 'message', from: nick, to: to, text: text, msg: msg

    @.addListener 'action', (from, to, text, msg) ->
      self.dispatch 'action', from: from, to: to, text: text, msg: msg

    # Names events

    @.addListener 'names', (channel, nicks) ->
      self.dispatch 'names', channel: channel, nicks: nicks

    # Join Events
    @.addListener 'join', (channel, nick, msg) ->
      self.dispatch 'join', channel: channel, nick: nick, msg: msg

    # Part events

    @.addListener 'part', (channel, nick, reason, msg) ->
      self.dispatch 'part', channel: channel, nick: nick, reason: reason, msg: msg

    # Quit events

    @.addListener 'quit', (nick, reason, channels, msg) ->
      self.dispatch 'quit', nick: nick, reason: reason, channels: channels, msg: msg

    # Nicks Events
    @.addListener 'nick', (oldnick, newnick, channels, msg) ->
      self.dispatch 'nick', oldnick: oldnick, newnick: newnick, channels: channels, msg: msg

    # PM events

    @.addListener 'pm', (nick, text, msg) ->
      self.dispatch 'pm', from: nick, text: text, msg: msg

    # Ping events

    @.addListener 'ping', (server) ->
      rainlog.info 'Bot', ':Pong ' + server
      self.dispatch 'ping', server: server

    # Error events

    @.addListener 'error', (message) ->
      rainlog.err 'Server', message.args[1]

    return callback() # Finished attaching hooks

  gate: (channel) ->
    self = @
    self.addListener 'raw', gate = (message) ->
      if self.sleep and message.command != 'PRIVMSG'
        self.removeListener 'raw', gate
        self.bot.sleep = false
      if message.args[1]?.indexOf('Replaying up to') > -1
        self.sleep = true
    if channel then @.join channel
    rainlog.info 'Bot', 'Joined ' + channel

module.exports = bot
