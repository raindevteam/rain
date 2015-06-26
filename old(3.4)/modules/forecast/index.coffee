triggers = require('./listeners/triggers')
commands = require('./listeners/commands')

module.exports = (Module) ->
  Forecast = new Module('Forecast')
  Forecast.addTriggers(triggers)
  Forecast.addCommands(commands)
  return Forecast
