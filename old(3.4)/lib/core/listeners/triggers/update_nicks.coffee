bot = __core.bot

module.exports =
  ### == update_nicks (event:names) ==
  Pushes a nick on join
  ###
  user_join:
    event: 'names'
    trigger: (message) -> return true
    action: (respond, done) ->
      if !bot.nicks then bot.nicks = []
      for key of bot.nicks
        if !bot.nicks.indexOf(key)
          bot.nicks.push key
      return done()
