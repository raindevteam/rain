eplist  = require('../lib/eplist')
core    = require(__core)
trigger = core.trigger

module.exports = {
  # Suggest an episode
  episode:
    event: 'message'
    nest:
      name: 'episode'
      ASAP: false
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'episode')
        fire: (respond, done) ->
          ep = eplist.get_rand_ep()
          respond.say to, 'You should watch ' + 
          ep[1] + ' ' + eplist.key_to_readable(ep[0])
          done()
      }  
  # Get name of key
  name:
    event: 'message'
    nest:
      name: 'nameof'
      ASAP: false
    action: (channel, to, text, msgs) ->
      return {
        trigger: trigger.cmd(text, 'nameof')
        fire: (respond, done) ->
          find = text.after('!').trim().match(/^s(.*)e(.*)/)
          if !find 
            respond.say to,
            'Wrong usage of \'Name of\', try' +
            ' something like \'Name of s2e04\''
            done()
          if !(find[1]?.match(/^\d/))
            respond.say to, 'Invalid arg format for \'s\''
            done()
          if !(find[2]?.match(/^\d{1,2}$/))
            respond.say to, 'Invalid arg format for \'e\''
            done()
          else 
            if find[2].match(/^\d$/)
              find[2] = '0' + find[2]

          result = eplist.get_ep('s'+find[1]+'e'+find[2])
          if result then respond.say to, result
          else respond.say to, 'Couldn\'t find anything'
          done()
      }  
}
      