core = require(__core)
trigger = core.trigger
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

module.exports = {
  # Get a random quotes
  quote:
    event: 'message'
    nest:
      name: 'quote'
      ASAP: false
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'random')
        fire: (respond, done) ->
          text = get_rand_quote()
          text = text.replace('{from}', channel)
          if text.indexOf('*') > -1
            while text.indexOf('*') > -1
              resPos = text.indexOf('*')
              resNext = text.indexOf('*', resPos+1)
              resText = text.before('*')
              respond.say to, resText if resText != ''
              respond.action to, text.substring(resPos+1, resNext)
              text = text.substring(resNext+1)
            if text then respond.say to, text
            done() 
          else
            respond.say to, text
            done()                    
      }
  # Spits out rules
  rules:
    event: 'message'
    nest:
      name: 'rules'
      ASAP: false
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'rules')
        fire: (respond, done) ->
          respond.say to, 'bucket, rules'
          respond.output 'http://tinyurl.com/ks6yml4'
          done()
      }

  say:
    event: 'message'
    nest:
      name: 'say'
      ASAP: true
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'say')
        fire: (respond, done) ->
          respond.say to, text.after('say!').clean()
          done()
      }  
  
  # Hashes input
  hash:
    event: 'message'
    nest:
      name: 'hash'
      ASAP: false
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'hash')
        fire: (respond, done) ->
          respond.say to, strhash(text.after('!').clean())
          done()
      }
  
  # Respond to master race
  masterrace:
    event: 'message'
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.hasAnywhere(text, 'is masterrace')
        fire: (respond, done) ->
          respond.say to, 'botponies are masterrace'
          done()
      }
  
  # Respond to who is best pone
  bestpone:
    event: 'message'
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.hasAnywhere(text, 'who is best pony')
        fire: (respond, done) ->
          respond.say to, 'bucket is best pony'
          respond.action to, 'kicks bucket'
          done()
      }
}