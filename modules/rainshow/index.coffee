commands = require('./listeners/commands')

module.exports = (Module) ->
  RainShow = new Module('RainShow')
  RainShow.addCommands(commands)
  return RainShow
