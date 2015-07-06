module.exports  = (bot) ->

  return {

    connected: (channel, who) ->
      if !who then who == bot.nick
      for chan in bot.chans
        if chan.key == channel.toLowerCase()
          return true
      return false

  }
