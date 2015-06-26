config = require(__config)

commands = require('./listeners/commands')
triggers = require('./listeners/triggers')

module.exports = (Module) ->
  Social = new Module('Social')
  Social.addCommands(commands)
  Social.addTriggers(triggers)
  return Social
