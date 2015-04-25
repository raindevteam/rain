commands = require('./listeners/commands')

module.exports = (Module) ->
  X = new Module('X')
  X.addCommands(commands)
  return X
