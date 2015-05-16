module.exports =
  server   : process.env.IRC_SERVER or 'irc.canternet.org'
  nick     : process.env.IRC_NICK or 'SnowBot'
  password : process.env.IDENTIFY_PASSWORD or undefined
  channels : [ process.env.CHANNEL or '#Snowybottest' ]
  db       : process.env.DB_URL or 'mongodb://mista:BCo1370S@ds043947.mongolab.com:43947/rainbot'
  helpLink : 'rainbot.info/userguide'
  debug    : true
  bang     : 'end'

  info:
    description : 'a multipurpose IRC botpony'
    version     : '3.3.15'
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
