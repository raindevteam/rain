bot = __core.bot
_   = require('lodash')

module.exports =
  ### == user_part (event:part) ==
  Removes nick on user part
  ###
  user_part:
    event: 'part'
    trigger: (message) -> return true
    action: (respond, done) ->
      if bot.nicks.indexOf(respond.nick)
        _.pull(bot.nicks, respond.nick)
      return done()
