eplist  = require('../lib/eplist')

module.exports =
  # Suggest an episode
  'ep':
    action: (data, respond, done) ->
      if data.args.length > 2
        respond.say 'Too many params, try something like \'1 12\''
        return done()
      if data.args[0]
        season = data.args[0]
        if isNaN season
          respond.say data.args[0] + ' is not a valid param'
          return done()
        if data.args[1]
          episode = data.args[1]
          if isNaN episode
            respond.say data.args[1] + ' is not a valid param'
            return done()
          if episode.toString().length == 1
            episode = '0' + episode
          result = eplist.get_ep('s'+season+'e'+episode)
          if result then respond.say result
          else respond.say 'Couldn\'t find anything'
          return done()
        else
          respond.say 'No episode number entered'
          return done()
      else
        ep = eplist.get_rand_ep()
        respond.say 'You should watch ' +
        ep[1] + ' ' + eplist.key_to_readable(ep[0])
        return done()
