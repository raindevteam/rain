module.exports = (module) ->

  # chan
  # ----

  chan:
    help: 'Allows the bot to join and leave channels (rainbot.info/userguide#chan)'
    ASAP: false
    action: (data, respond, done) ->
      cmd   = data.args.join(' ')

      JOIN  = cmd.match(/\bjoin\s*(.*(?= leave)|(?=leave)|.*)/i)
      LEAVE = cmd.match(/\bleave\s*(.*(?= join)|(?=join)|.*)/i)

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
      else if LEAVE then LEAVE = data.to

      if JOIN
        data.bot.gate JOIN
        respond.say('Joined ' + JOIN).now()
      if LEAVE
        data.bot.part LEAVE
        respond.say("Left " + LEAVE).now()
      return done()
