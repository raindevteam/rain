config  = require(__config)
core    = require(__core)
trigger = core.trigger

event = 'message'

module.exports = {
  # info
  #   Display info about the bot
  info :
    event : event
    nest  :
      name : 'info'
      ASAP : false
    action: (nick, to, text, msgs) ->
      return {
        trigger : trigger.cmd(text, 'info')
        fire    : (respond, done) ->
          info = config.info
          respond.say to, 
          'I\'m ' + info['description'] + 
          ' ' + 'written with ' + 
          info['writtenIn'] + ' ' + 
          'by ' + info['developer'] + 
          ' ' + '[Version: ' + 
          info['version'] + ' ' + 
          '(' + info['versionName'] + ')]'
          done()
      }

  # help
  #   Returns a help link pointing
  #   to a user's guide
  help :
    event : event
    nest  :
      name : 'help'
      ASAP : false
    action: (nick, to, text, msgs) ->
      {
        trigger : trigger.cmd(text, 'help') 
        fire    : (respond, done) ->
          respond.say to, 'Have a link: ' + config.helpLink,
          respond.output config.helpLink
          done()
      }

  # help
  #   Returns a help link pointing
  #   to a user's guide
  aliasAdd:
    event : event
    action: (nick, to, text, msgs) ->
      {
        trigger : trigger.cmd(text, 'alias') 
        fire    : (respond, done) ->
          args = text.after('alias!').trim()
          if (args.has('-rm'))
            argsArr = args.split(' ')
            core.alias.deleteAlias(argsArr[1])
            respond.say to, 'Removed Alias: ' + argsArr[1]
          else
            argsArr = []
            argsArr[0] = args.before(':').trim()
            if argsArr[0] == ''
              respond.say to, 'No alias name entered!'
              return done()
            argsArr[1] = args.after(':').trim()
            if argsArr[1] == ''
              respond.say to, 'No command to alias entered!'
              return done()
            respond.say to, 'Added command alias: ' + argsArr[0]
            core.alias.addAlias(argsArr[0], argsArr[1])
          return done()
      }

  # help
  #   Returns a help link pointing
  #   to a user's guide
  aliasDel:
    event : event
    action: (nick, to, text, msgs) ->
      {
        trigger : trigger.cmd(text, 'alias') 
        fire    : (respond, done) ->
          args = text.after('alias!').split(':')
          args[0] = args[0].trim()
          args[1] = args[1].trim()
          respond.say to, 'Added command alias: ' + args[0]
          core.aliasHandler.addAlias(args[0], args[1])
          done()
      }

  # tuner
  #   Allows the bot to switch
  #   between different channels
  tuner :
    event : event
    nest  :
      name : 'chan'
      ASAP : true
    action: (nick, to, text, msgs) ->
      {
        trigger : trigger.cmd(text, 'chan')
        fire    : (respond, done) ->
          cmd   = text.after('!').clean()
          JOIN  = undefined
          LEAVE = undefined

          JOIN  = cmd.match(/\bjoin\s*(.*(?= leave)|(?=leave)|.*)/)
          LEAVE = cmd.match(/\bleave\s*(.*(?= join)|(?=join)|.*)/)

          if JOIN and JOIN[1] and JOIN[1].match(RegExp(' '))
            respond.say to, 'Too many arguments to command join'
            return done()
          if LEAVE and LEAVE[1] and LEAVE[1].match(RegExp(' '))
            respond.say to, 'Too many arguments to command leave'
            return done()

          if JOIN and JOIN[1]
            JOIN = JOIN[1].split(':').join(' ')
            if JOIN[0] != '#'
              respond.say to, 'Invalid channel format for join'
              return done()
          else if JOIN then JOIN = null

          if LEAVE and LEAVE[1]
            LEAVE = LEAVE[1].split(':').join(' ')
            if LEAVE[0] != '#'
              respond.say to, 'Invalid channel format for leave'
              return done()
          else if LEAVE then LEAVE = to

          core.CHANNEL_SWITCH JOIN, LEAVE, to
          done()
      }

  sleep :
    event : event
    nest :
      name : 'sleep'
      ASAP : false
    action : (nick, to, text, msgs) ->
      {
        trigger : trigger.cmd(text, 'sleep')
        fire    : (respond, done) ->
          respond.say to, 'Don\'t forget to wake me'
          respond.action to, 'Lays down and idles'
          core.IDLE()
          done()
      }
}
