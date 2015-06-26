commands = require('./listeners/commands')

module.exports = (Module) ->
  RainBrella = new Module('RainBrella')
  RainBrella.addCommands(commands)
  return RainBrella
