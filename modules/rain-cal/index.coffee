math   = require('mathjs')
parser = math.parser()

commands =
  cal:
    nest: true
    ASAP: false
    action: (args, respond, done) ->
      eqt = args.join(" ")
      try # Try parsing
        res = parser.eval(eqt)
      catch e # Quietly fail
      # Return result if not a function
      respond.say to, res if !__core.helpers.isFunction(res)
      done()

module.exports = (Module) ->
  RainCal = new Module('RainCal')
  RainCal.addCommands(commands)
  return RainCal
