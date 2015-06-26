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
      console.log args + " " + args.length
      reps = 1
      if args.length > 0 and !isNaN(args[0])
        reps = args[0]
      rolls = []
      for i in [0...reps]
        rolls.push(randomIntFromInterval(1, 6))
      respond.say rolls
      return done()
