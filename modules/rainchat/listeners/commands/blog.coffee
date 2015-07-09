module.exports = (Module) ->

  # blog
  # ----

  'blog':
    help: 'Gives link to "The RainBot Blog" (rainbot.info/userguide#blog)'
    ASAP: false
    action: (data, respond, done) ->
      respond.say 'rainbot-irc.blogspot.com'
      return done()
