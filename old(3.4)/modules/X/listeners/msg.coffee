core    = require(__core)
trigger = core.trigger
sandbox = require('sandbox')
rainbox = new sandbox()

randomIntFromInterval = (min,max) ->
  return Math.floor(Math.random()*(max-min+1)+min);

module.exports = {
  ###
  command: (nick, to, text, msg) ->
    return {
      name: 'test'
      nest: true
      ASAP: false
      fire: (respond) ->
        respond.say to, 'test command'
    }
  ###
  
  read:
    event: 'message'
    nest: 
      name: 'jrun'
      ASAP: false
    action: (channel, to, text, msg) ->
      return {
        trigger: trigger.cmd(text, 'jrun')
        fire: (respond, done) ->
          rainbox.run text.after('!'), (output) ->
            text = output.result
            text = text.replace(/^'/, '').replace(/'$/, '')
            respond.say to, text
            done()
      }

  roll:
    event: 'message'
    nest: 
      name: 'roll'
      ASAP: false
    action: (channel, to , text, msg) ->
      return {
        trigger: trigger.cmd(text, 'roll')
        fire: (respond, done) ->
          args = text.after('roll!')
          reps = 0
          if args.trim() == ''
            reps = 1
          else if args.trim().split(' ').length == 1
            reps = args.trim() if (!isNaN(args.trim()))
          else done()
          rolls = []
          for i in [0...reps]
            rolls.push(randomIntFromInterval(1, 6))
          result = rolls.join(', ')
          respond.say to, result
          respond.output rolls
          done()
      }
}
  