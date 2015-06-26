events     = require('./events.coffee')

class Module
  constructor: (name) ->
    @name = name
    @commands = []
    @triggers = {}
    for event in events
      @triggers[event] = []

  setBot: (bot) ->
    @bot = bot

  addCommands: (commands) ->
    for name, command of commands
      @commands.push
        'name': name
        'command': command

  addTriggers: (triggers) ->
    for name, trigger of triggers
      @triggers[trigger.event].push(trigger)

module.exports = Module
