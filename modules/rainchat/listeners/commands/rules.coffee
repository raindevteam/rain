module.exports = (Module) ->

  # rules
  # -----

  rules:
    help: 'Spits out rules (rainbot.info/userguide#rules)'
    ASAP: false
    action: (data, respond, done) ->
      irc = data.irc
      if irc.connected(data.to, 'bucket')
        respond.say 'bucket, rules'
      else
        respond.say 'http://tinyurl.com/ks6yml4'
      respond.output 'http://tinyurl.com/ks6yml4'
      return done()
