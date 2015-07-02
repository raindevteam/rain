module.exports = (module) ->

  # version (Command)
  # Says the bot's version

  'version':
    help: 'Shows the bot\'s version (rainbot.info/userguide#version)'
    ASAP: false
    action: (data, respond, done) ->
      respond.say data.bot.version
      return done()
