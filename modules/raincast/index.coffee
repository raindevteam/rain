triggers = require('./listeners/triggers')
commands = require('./listeners/commands')

module.exports = (Module) ->
  RainCast = new Module('RainCast')
  RainCast.addTriggers(triggers)
  RainCast.addCommands(commands)
  return RainCast
