messages = require('../defined')
_ = require('lodash')

module.exports = (_bot) ->
  {
    user_leave: 
      event: 'part'
      action : (channel, nick, reason, msg) ->
        {
          trigger : true
          fire    : (respond) ->
            if _bot.nicks.indexOf(nick) > -1
              _.pull(_bot.nicks, nick)
        }
  }
