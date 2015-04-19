messages = require('../defined')

module.exports = () ->
  { 
    update :
      event : 'names'
      action: (channel, nicks) ->
        {
          trigger : true
          fire : (respond) ->
            if _bot.nicks == undefined
              _bot.nicks = []
            for key of nicks
              if _bot.nicks.indexOf(key) < 0
                _bot.nicks.push key
        }
  }
