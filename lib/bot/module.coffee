hashmap = require 'hashmap'
_ = require 'lodash'

mpi = require './mpi'
Hook = require './hook'
events = require './events'

HookHandler = require './hookhandler'

isValidCommand = (command) ->
  return false if !command
  return false if !command.action
  return false if !(_.isFunction command.action)
  return true

isValidTrigger = (trigger) ->
  return false if !trigger
  return false if !trigger.event
  return false if !trigger.trigger
  return false if !(_.isFunction trigger.trigger)
  return false if !trigger.action
  return false if !(_.isFunction trigger.action)
  return true

module.exports = (bot) ->

  class Module

    # Module :: constructor (String, String)
    #
    # Initializes name and abbreviation of module. Also sets the bot
    # and declares the enabled bool, commands array and triggers object

    constructor: (name, abbrev) ->
      @name = name
      @abbrev = abbrev
      @bot = bot
      @enabled = true
      @commands = new hashmap()
      @triggers = {}
      for event in events
        @triggers[event] = []

    # Module :: addCommand (String, command Object) throws Error
    #
    # Takes a command and its name and sets it to the module's hashmap.
    # If the command is invalid, the method throws.

    addCommand: (name, command) ->
      if !isValidCommand(command) then throw new Error('Invalid command')
      hook = new Hook(name, 'command', @, command)
      @commands.set(name, hook)

    # Module :: addCommands (commands Object) throws Error
    #
    # Very similar to @addCommand, except iterates over an
    # object collection

    addCommands: (commands) ->
      for name, command of commands
        if !isValidCommand(command) then throw new Error('Invalid command')
        hook = new Hook(name, 'command', @, command)
        @commands.set(name, hook)

    # Module :: addTrigger (String, trigger Object) throws Error
    #
    # Takes a trigger and its name and loads it into the module's
    # trigger object.

    addTrigger: (name, trigger) ->
      if !isValidTrigger(trigger) then throw new Error('Invalid trigger')
      hook = new Hook(name, 'trigger', @, trigger)
      @triggers[trigger.event].push(hook)

    # Module :: addTriggers (triggers Object) throws Error
    # Similar to addTrigger, except iterates over an object collection

    addTriggers: (triggers) ->
      for name, trigger of triggers
        if !isValidTrigger(trigger) then throw new Error('Invalid trigger')
        hook = new Hook(name, 'trigger', @, trigger)
        @triggers[trigger.event].push(hook)

    # Module :: disableCommand (String)
    # Disables named command

    disableCommand: (name) ->
      hook = @commands.get(name)
      hook.listening = false

    # Module :: enableCommand (String)
    # Enables named command

    enableCommand: (name) ->
      hook = @commands.get(name)
      hook.listening = true

    disableTrigger: (name) ->

    enableTrigger: (name) ->

    enable: () -> @enabled = true

    disable: () -> @enabled = false

    status: () -> return @enabled

  return Module

###
Module.getBot()
Module.disableCommand('paste')
Module.enableCommand('paste')
Module.disable()
Module.enable()
Module.status()
Module.getWhitelist()
###
