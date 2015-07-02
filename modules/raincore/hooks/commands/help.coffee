module.exports = (module) ->

  # help (Command)
  # Returns help info for Command

  'help':
    help: 'rainbot.info/userguide'
    ASAP: false
    action: (data, respond, done) ->
      if data.args[0]
        hook = data.hooks.getCommand(data.args[0])
        if !hook then respond.say 'Don\'t know anything about that'
        else          respond.say hook.help
      return done()
