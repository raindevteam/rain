core    = require(__core)

module.exports =
  ### == sleep (command) ==
  Puts the bot to sleep, during which the bot will
  not respond to any messages except for the wake listener
  ###
  sleep:
    nest: true
    ASAP: false
    action: (args, respond, done) ->
      respond.say 'Don\'t forget to wake me'
      respond.action 'Lays down and idles'
      core.IDLE()
      return done()
