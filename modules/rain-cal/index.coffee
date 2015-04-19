# Requires
core    = require(__core)
math    = require('mathjs')
trigger = core.trigger

# Locals
parser = math.parser()

listeners = {
  parse:
    event: 'message'
    nest:
      name: 'cal'
      ASAP: false
    action: (nick, to, text, msgs) ->
      {
        trigger: trigger.cmd(text, 'cal')
        fire: (respond, done) ->
          eqt = text.after('cal!') # We need a args handler
          try # Try parsing
            res = parser.eval(eqt)
          catch e # Quietly fail
          # Return result if not a function
          respond.say to, res if !core.helpers.isFunction(res)
          done()
      }
}

module.exports = (Module) ->
  RainCal = new Module('RainCal')
  RainCal.addListeners(listeners)
  return RainCal
