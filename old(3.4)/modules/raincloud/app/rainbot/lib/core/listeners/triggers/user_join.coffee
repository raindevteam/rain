bot = __core.bot

module.exports =
  ### == user_join (event:join) ==
  Pushes a nick on join
  ###
  user_join:
    event: 'join'
    trigger: (message) -> return true
    action: (respond, done) ->
      if !bot.nicks.indexOf(respond.nick)
        bot.nicks.push(respond.nick)
      return done()
