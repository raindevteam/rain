commands = require('./listeners/commands')

module.exports = (Module) ->
  MlpRcmdr = new Module('MlpRcmdr')
  MlpRcmdr.addCommands(commands)
  return MlpRcmdr
