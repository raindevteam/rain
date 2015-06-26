fs = require('fs')

quotes = []
file = '/../random.txt'

fs.readFile __dirname + file, (err, data) ->
  raw    = data.toString()
  buffer = raw.split('\n')
  for line in buffer
    continue if line[0] == '#'
    quotes.push(line)

get_rand_quote = ->
  rand_index = Math.floor(Math.random() * quotes.length)
  return quotes[rand_index]

strhash = (str) ->
  hash = 0
  for letter in str
    hash = Math.abs(letter.charCodeAt(0) +
    (hash << 6) + (hash << 16) - hash)
  return hash

module.exports =
  # Get a random quote
  quote:
    action: (args, respond, done) ->
      text = get_rand_quote()
      text = text.replace('{from}', respond.channel)
      if text.indexOf('*') > -1
        while text.indexOf('*') > -1
          resPos = text.indexOf('*')
          resNext = text.indexOf('*', resPos+1)
          resText = text.before('*')
          respond.say resText if resText != ''
          respond.action text.substring(resPos+1, resNext)
          text = text.substring(resNext+1)
        if text then respond.say text
        return done()
      else
        respond.say text
        return done()

  # Spits out rules
  rules:
    action: (args, respond, done) ->
      respond.say 'bucket, rules'
      respond.output 'http://tinyurl.com/ks6yml4'
      return done()

  say:
    ASAP: true
    action: (args, respond, done) ->
      respond.say args.join(" ") if args
      return done()

  # Hashes input
  hash:
    action: (args, respond, done) ->
      respond.say strhash(args[0])
      return done()
