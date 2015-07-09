sandbox = require('sandbox')
rainbox = new sandbox()

randomIntFromInterval = (min,max) ->
  return Math.floor(Math.random()*(max-min+1)+min)

strhash = (str) ->
  hash = 0
  for letter in str
    hash = Math.abs(letter.charCodeAt(0) +
    (hash << 6) + (hash << 16) - hash)
  return hash

module.exports =
  # Run a javascript snippet
  jsnip:
    help: 'Runs a javascript snippet (rainbot.info/userguide#jsnip)'
    action: (data, respond, done) ->
      rainbox.run data.args.join(" "), (output) ->
        text = output.result
        text = text.replace(/^'/, '').replace(/'$/, '')
        respond.say text
        return done()

  # Roll some dice
  roll:
    help: 'Roll some dice (rainbot.info/userguide#roll)'
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
    help: 'Repeats arguments (rainbot.info/userguide#echo)'
    action: (data, respond, done) ->
      respond.say data.args.join ' ' if data.args
      return done()

  # Hash some input
  hash:
    help: 'Hashes a string, or your nick if called by itself (rainbot.info/userguide#hash)'
    action: (data, respond, done) ->
      if data.args[0]
        respond.say strhash(data.args[0])
      else
        respond.say strhash(data.from)
      return done()
