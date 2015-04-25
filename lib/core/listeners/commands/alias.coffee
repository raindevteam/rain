core = require(__core)

module.exports =
  ### == alias (command) ==
  Sets and removes aliases
  ###
  alias:
    action: (args, respond, done) ->
      if (args[0] == 'rm')
        core.alias.deleteAlias(args[1])
        respond.say 'Removed Alias: ' + args[1]
      else
        alias = args[0]
        command = args[1]
        if !alias
          respond.say 'No alias name entered!'
          return done()
        if !command
          respond.say 'No command to alias entered!'
          return done()
        respond.say 'Added command alias: ' + alias
        core.alias.addAlias(alias, command)
      return done()
