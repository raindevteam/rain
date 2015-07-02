module.exports = (Module) ->
  RainCore = new Module('RainCore')
  RainCore.addCommand 'module',
    help: 'Lists, enables, and disables modules (rainbot.info/userguide#module)'
    action: (data, respond, done) ->
      if data.args[0]?.lower() is 'list'
        names = (module.name for module in data.bot.modules)
        respond.say names.join ', '
      else if data.args[0]?.lower() is 'disable'
        if data.args[1]
          for module in data.bot.modules
            if module.name.lower() == data.args[1].lower()
              module.disable()
              respond.say 'Disabled module ' + module.name
              return done()
      else if data.args[0]?.lower() is 'enable'
        if data.args[1]
          for module in data.bot.modules
            if module.name.lower() == data.args[1].lower()
              module.enable()
              respond.say 'Enabled module ' + module.name
              return done()
      return done()
  RainCore.addCommand 'ver',
    help: 'Shows the bot\'s version (rainbot.info/userguide#version)'
    action: (data, respond, done) ->
      respond.say data.bot.version
      return done()
  RainCore.addCommand 'help',
    help: 'If you put something else other than help in front it, you know... it might help'
    action: (data, respond, done) ->
      if data.args[0]
        hook = data.hooks.getCommand(data.args[0])
        if !hook then respond.say 'Don\t know anything about that'
        else          respond.say hook.help
      return done()
  return RainCore
