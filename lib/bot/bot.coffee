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
    self = @
    fs.readdir MODULES_DIRECTORY, (err, modules) ->
      # Check err
      for module in modules
        # check if is actually a module
        module = require(MODULES_DIRECTORY + module)(self.Module)
        self.modules.push(module)
        HookHandler.extractCommands(module)
        HookHandler.extractTriggers(module)
      return callback()

  start: () ->
    self = @
    self.loadModules -> self.attachHooks ->
      require('./../../config/init')(@, ->
        self.connect ->
          if __config.nsPassword
            self.send 'ns', 'identify', __config.nsPassword
          self.send 'mode', __config.nick, '+B'
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

    @.addListener 'error', (message) ->
      console.log "from error " + message

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
