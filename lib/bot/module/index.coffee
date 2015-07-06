# Requires
# --------

# From node_modules

hashmap = require 'hashmap'
callsite = require 'callsite'
fs = require 'fs'
_ = require 'lodash'

# From Project

mpi = require './mpi'
Hook = require './hook'
events = require './../events'

# isValidCommand (command Object)

isValidCommand = (command) ->
  return false if !command
  return false if !command.action
  return false if !(_.isFunction command.action)
  return true

# isValidTrigger (trigger Object)

isValidTrigger = (trigger) ->
  return false if !trigger
  return false if !trigger.event
  return false if !trigger.trigger
  return false if !(_.isFunction trigger.trigger)
  return false if !trigger.action
  return false if !(_.isFunction trigger.action)
  return true

module.exports = (bot) ->

  # Module class
  # ------------

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
    # Creates a new command hook and adds it to the command hashmap

    addCommand: (name, command) ->
      if !isValidCommand(command)
        err = new Error('Invalid command: ' + name)
        rainlog.err @.name, err.message
        return err
      hook = new Hook(name, 'command', @, command)
      @commands.set(name, hook)

    # Module :: addCommands (commands Object) throws Error
    # Similar to @addCommand, but adds from an object collection

    addCommands: (commands) ->
      if _.isObject(commands)
        @addCommand(name, command) for name, command of commands
      else
        if fs.lstatSync(commands).isDirectory()
          files = fs.readdirSync(commands)
          for file in files
            command = require(commands + '/' + file)(@)
            @addCommand(key, val) for key, val of command

    # Module :: addTrigger (String, trigger Object) throws Error
    # Creates new hook trigger and adds it to the triggers object

    addTrigger: (name, trigger) ->
      if !isValidTrigger(trigger) then throw new Error('Invalid trigger')
      hook = new Hook(name, 'trigger', @, trigger)
      @triggers[trigger.event].push(hook)

    # Module :: addTriggers (triggers Object) throws Error
    # Similar to @addTrigger, but adds from an object collection

    addTriggers: (triggers) ->
      for name, trigger of triggers
        @addTrigger(name, trigger)

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

    # Module :: enable ()
    # Enables module

    enable: () -> @enabled = true

    # Module :: disable ()
    # Disables module

    disable: () -> @enabled = false

    # Module :: status ()
    # Returns @enabled

    status: () -> return @enabled

  return Module
