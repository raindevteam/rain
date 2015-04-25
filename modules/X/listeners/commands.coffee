sandbox = require('sandbox')
rainbox = new sandbox()

randomIntFromInterval = (min,max) ->
  return Math.floor(Math.random()*(max-min+1)+min);

module.exports =
  jrun:
    action: (args, respond, done) ->
      rainbox.run args.join(" "), (output) ->
        text = output.result
        text = text.replace(/^'/, '').replace(/'$/, '')
        respond.say text
        return done()

  roll:
    action: (args, respond, done) ->
      reps = 0
      if !args.length
        reps = 1
      reps = args[0] if (!isNaN(args[0]))
      rolls = []
      for i in [0...reps]
        rolls.push(randomIntFromInterval(1, 6))
      result += roll for roll in rolls
      respond.say result
      respond.output rolls
      return done()
