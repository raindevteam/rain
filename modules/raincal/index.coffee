math   = require('mathjs')
_      = require('lodash')
parser = math.parser()

commands =
  cal:
    help: 'Parses math (rainbot.info/userguide#cal)'
    ASAP: false
    action: (data, respond, done) ->
      eqt = data.args.join(" ")
      try # Try parsing
        res = parser.eval(eqt)
      catch e # Quietly fail
      # Return result if not a function
      respond.say res if !_.isFunction(res)
      return done()

module.exports = (Module) ->
  RainCal = new Module('RainCal')
  RainCal.addCommands(commands)
  return RainCal
