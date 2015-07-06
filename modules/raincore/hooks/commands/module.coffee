module.exports = (module) ->

  # module (Command)
  # Handles the bot's modules

  'module':
    help: 'Lists, enables, and disables modules (rainbot.info/userguide#module)'
    ASAP: false
    action: (data, respond, done) ->
      args = data.args
      return done() if !args[0]
      switch args[0].toLowerCase()
        when 'list'
          names = (module.name for module in data.bot.modules)
          respond.say names.join ', '
        when 'enable'
          if data.args[1]
            for module in data.bot.modules
              if module.name.toLowerCase() == data.args[1].toLowerCase()
                module.disable()
                respond.say 'Disabled module ' + module.name
                return done()
        when 'disable'
          if data.args[1]
            for module in data.bot.modules
              if module.name.toLowerCase() == data.args[1].toLowerCase()
                module.enable()
                respond.say 'Enabled module ' + module.name
                return done()
      return done()
