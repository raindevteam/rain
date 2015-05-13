_       = require('lodash')

module.exports = () ->
  {
    user_join: 
      event  : 'join'
      action : (channel, nick, msg) ->
        {
          trigger : true
          fire : (respond) ->
            if _bot.nicks.indexOf(nick) < 0
              _bot.nicks.push(nick)
        }
  }
