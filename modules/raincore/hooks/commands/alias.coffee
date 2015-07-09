_    = require('lodash')

module.exports = (module) ->

  # Alias
  # -----

  alias:
    help: 'Lets you alias a/some command(s) (rainbot.info/userguide#alias)'
    action: (data, respond, done) ->
      args = data.args
      alias = data.bot.alias
      if (args[0] == 'rm')
        alias.delete args[1], () ->
          respond.say 'Removed Alias: ' + args[1]
          return done()
      else
        aliasName = args[0]
        command = _.drop(args).join(' ')
        if !aliasName
          respond.say 'No alias name entered!'
          return done()
        if !command
          respond.say 'No command to alias entered!'
          return done()
        respond.say 'Added command alias: ' + aliasName
        alias.add(aliasName, command, done)
