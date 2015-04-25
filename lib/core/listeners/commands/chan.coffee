core = require(__core)

module.exports =
  ### == chan (command) ==
  Allows the bot to switch channels
  ###
  chan:
    nest: true
    ASAP: false
    action: (args, respond, done) ->
      cmd   = args.join(' ')

      JOIN  = cmd.match(/\bjoin\s*(.*(?= leave)|(?=leave)|.*)/)
      LEAVE = cmd.match(/\bleave\s*(.*(?= join)|(?=join)|.*)/)

      if JOIN and JOIN[1] and JOIN[1].match(RegExp(' '))
        respond.say 'Too many arguments to command join'
        return done()
      if LEAVE and LEAVE[1] and LEAVE[1].match(RegExp(' '))
        respond.say 'Too many arguments to command leave'
        return done()

      if JOIN and JOIN[1]
        JOIN = JOIN[1].split(':').join(' ')
        if JOIN[0] != '#'
          respond.say 'Invalid channel format for join'
          return done()
      else if JOIN then JOIN = null

      if LEAVE and LEAVE[1]
        LEAVE = LEAVE[1].split(':').join(' ')
        if LEAVE[0] != '#'
          respond.say 'Invalid channel format for leave'
          return done()
      else if LEAVE then LEAVE = respond.to

      core.CHANNEL_SWITCH JOIN, LEAVE, respond.to
      return done()
