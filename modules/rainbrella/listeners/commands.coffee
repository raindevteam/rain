sandbox = require('sandbox')
rainbox = new sandbox()

randomIntFromInterval = (min,max) ->
  return Math.floor(Math.random()*(max-min+1)+min);

module.exports =
  # Run a javascript snippet
  jrun:
    action: (data, respond, done) ->
      rainbox.run data.args.join(" "), (output) ->
        text = output.result
        text = text.replace(/^'/, '').replace(/'$/, '')
        respond.say text
        return done()

  # Roll some dice
  roll:
    action: (data, respond, done) ->
      reps = 1
      if data.args.length > 0 and !isNaN(data.args[0])
        reps = data.args[0]
      rolls = []
      for i in [0 .. reps]
        rolls.push(randomIntFromInterval(1, 6))
      respond.say rolls
      return done()

  # Repeat args
  echo:
    action: (data, respond, done) ->
      respond.say data.args.join ' ' if data.args
      return done()
