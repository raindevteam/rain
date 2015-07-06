_    = require('lodash')

module.exports = (module) ->

  # Alias
  # -----

  alias:
    action: (data, respond, done) ->
      args = data.args
      alias = data.bot.alias
      if (args[0] == 'rm')
        alias.deleteAlias(args[1])
        respond.say 'Removed Alias: ' + args[1]
      else
        alias = args[0]
        command = _.drop(args).join(' ')
        if !alias
          respond.say 'No alias name entered!'
          return done()
        if !command
          respond.say 'No command to alias entered!'
          return done()
        respond.say 'Added command alias: ' + alias
        alias.addAlias(alias, command)
      return done()
