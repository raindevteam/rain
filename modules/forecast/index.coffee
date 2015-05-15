triggers = require('./nicks')

module.exports = (Module) ->
  Forecast = new Module('Forecast')
  # Forecast.addTriggers(triggers)
  return Forecast
