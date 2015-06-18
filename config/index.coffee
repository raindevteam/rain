module.exports =
  server   : process.env.IRC_SERVER or 'irc.canternet.org'
  nick     : process.env.IRC_NICK or 'HailBot'
  password : process.env.IDENTIFY_PASSWORD or undefined
  channels : [ process.env.CHANNEL or '#Snowybottest' ]
  helpLink : 'rainbot.info/userguide'
  loggingLevel : true
  bang     : 'end'

  info:
    description : 'a multipurpose IRC botpony'
    version     : 'EXP 3.5.0'
    versionName : 'Xenith'
    developer   : 'MistaWolf'
    writtenIn   : 'node.js'

  whitelist: [
    'Dustrunner'
    'Eventide'
    'Powderprancer'
    'King_Sombra'
    'MistaWolf'
  ]

  whitelisted_funcs: [
    'tuner'
    'sleep'
    'wake'
  ]
