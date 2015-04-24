core       = require(__core)
events     = core.events
async      = require('async')

class Module
  constructor: (name) ->
    @name = name
    @commands = []
    @triggers = {}
    for event in core.events
      @triggers[event] = []

  setBot: (bot) ->
    @bot = bot

  addCommands: (commands) ->
    for name, command of commands
      @commands.push {
        'name': name
        'command': command
      }

  addTriggers: (triggers) ->
    for name, trigger of listeners
      @triggers[trigger.event].push(trigger)

module.exports = Module
