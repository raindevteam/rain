irc = require 'irc'
fs = require 'fs'

HookHandler = require './hookhandler'
RespondQueue = require './responders/respondqueue'
Alias = require './alias'

MODULES_DIRECTORY = __dirname + '/../../modules/'

class bot extends irc.Client

  # bot :: constructor (String, String, options Object)
  #
  # The constructor passes itself over to the RespondQueue which
  # uses it to send out responses from the responders. The irc.Client
  # superclass is also instantiated.

  constructor: (server, nick, options) ->
    rainlog.info 'Bot', 'Initializing bot...'
    @version = 'EXP pre:4.0.0 (Pyrelight)'
    @Module = require('./module')(@)
    @alias = new Alias()
    @modules = []
    @config = __config
    RespondQueue.setBot(@)
    super server, nick, options

  # bot :: loadModules (Callback)
  # Loads up the modules

  loadModules: (callback) ->
    rainlog.info 'Bot', 'Loading modules...'
    self = @
    fs.readdir MODULES_DIRECTORY, (err, modules) ->
      for moduleDir in modules
        if !fs.lstatSync(MODULES_DIRECTORY + moduleDir).isDirectory
          rainlog.warn 'Bot', moduleDir + ' is not a module directory'
          continue
        try
          module = require(MODULES_DIRECTORY + moduleDir)(self.Module)
        catch e
          rainlog.err 'Bot', 'Couldn\'t load module: ' + moduleDir
          rainlog.err 'Bot', 'Check if the module has a (proper) index file!'
          continue
        if !(module instanceof self.Module)
          rainlog.err 'Bot', moduleDir + ' is not a module'
          rainlog.err 'Bot', 'Make sure that module exports a Module instance'
          continue
        self.modules.push(module)
        HookHandler.extractHooks(module)
        rainlog.info 'Bot', 'Loaded module: ' + module.name
      return callback()

  start: () ->
    self = @
    rainlog.info 'Bot', 'Starting bot...'
    self.loadModules -> self.attachHooks ->
      require('./../../config/init')(@, ->
        self.connect ->
          rainlog.info 'Bot', 'Bot connected to IRC'
          if __config.nsPassword
            self.send 'ns', 'identify', __config.nsPassword
          self.send 'mode', __config.nick, '+B'
          rainlog.info 'Bot', 'Connecting to channels'
          self.join '#snowybottest'
      )

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

  ###
  gate: (channel) ->
    self = @
    self.addListener 'raw', gate = (message) ->
      if self.sleep and message.command != 'PRIVMSG'
        self.removeListener 'raw', gate
        self.bot.sleep = false
      if message.args[1]?.indexOf('Replaying up to') > -1
        self.sleep = true
    if channel then @bot.join channel
  ###

module.exports = bot
