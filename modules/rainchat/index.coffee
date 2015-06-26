config = require(__config)

commands = require('./listeners/commands')
triggers = require('./listeners/triggers')

module.exports = (Module) ->
  RainChat = new Module('RainChat')
  RainChat.addCommands(commands)
  RainChat.addTriggers(triggers)
  return RainChat
