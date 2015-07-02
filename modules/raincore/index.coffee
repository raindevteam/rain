module.exports = (Module) ->
  RainCore = new Module('RainCore')
  RainCore.addCommands __dirname + '/hooks/commands'
  return RainCore
