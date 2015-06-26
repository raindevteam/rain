eplist  = require('../lib/eplist')

module.exports =
  # Suggest an episode
  episode:
    action: (args, respond, done) ->
      ep = eplist.get_rand_ep()
      respond.say 'You should watch ' +
      ep[1] + ' ' + eplist.key_to_readable(ep[0])
      return done()

  # Get name of key
  nameof:
    action: (args, respond, done) ->
      find = args[0].match(/^s(.*)e(.*)/)
      if !find
        respond.say (
          'Wrong usage of \'Name of\', try' +
          ' something like \'nameof! s2e4\'' )
        return done()
      if !(find[1]?.match(/^\d/))
        respond.say 'Invalid arg format for \'s\''
        return done()
      if !(find[2]?.match(/^\d{1,2}$/))
        respond.say 'Invalid arg format for \'e\''
        return done()
      else
        if find[2].match(/^\d$/)
          find[2] = '0' + find[2]

      result = eplist.get_ep('s'+find[1]+'e'+find[2])
      if result then respond.say result
      else respond.say 'Couldn\'t find anything'
      return done()
