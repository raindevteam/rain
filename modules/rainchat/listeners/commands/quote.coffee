fs = require 'fs'

quotes = []
file = '/../../random.txt'

fs.readFile __dirname + file, (err, data) ->
  raw    = data.toString()
  buffer = raw.split('\n')
  for line in buffer
    continue if line[0] == '#'
    quotes.push(line)

get_rand_quote = ->
  rand_index = Math.floor(Math.random() * quotes.length)
  return quotes[rand_index]

module.exports = (Module) ->

  # quote
  # -----

  'quote':
    help: 'Says a random quote (rainbot.info/userguide#quote)'
    ASAP: 'false'
    action: (data, respond, done) ->
      text = get_rand_quote()
      text = text.replace('{from}', respond.channel)
      if text.indexOf('*') > -1
        while text.indexOf('*') > -1
          resPos = text.indexOf('*')
          resNext = text.indexOf('*', resPos+1)
          resText = X(text).before('*').s
          respond.say resText if resText != ''
          respond.action text.substring(resPos+1, resNext)
          text = text.substring(resNext+1)
        if text then respond.say text
        return done()
      else
        respond.say text
        return done()
