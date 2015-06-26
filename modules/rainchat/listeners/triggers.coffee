module.exports =
  masterrace:
    event: 'message'
    trigger: (message) -> return false
    action: (respond, done) ->
      respond.say 'botponies are masterrace'
      return done()

  # Respond to who is best pone
  bestpone:
    event: 'message'
    trigger: (message) -> return false
    action: (respond, done) ->
      respond.say 'bucket is best pony'
      respond.action 'kicks bucket'
      return done()
