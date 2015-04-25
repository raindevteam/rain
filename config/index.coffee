module.exports =
  nick     : process.env.IRC_NICK or 'SnowBot'
  password : process.env.IDENTIFY_PASSWORD or undefined
  channels : [ process.env.CHANNEL or '#Snowybottest' ]
  db       : process.env.DB_URL or undefined
  helpLink : 'rainbot.info/userguide'
  debug    : true
  bang     : 'front'

  info:
    description : 'a multipurpose IRC botpony'
    version     : 'DEV 2.1.40'
    versionName : 'Red Eye'
    developer   : 'MistaWolf'
    writtenIn   : 'io.js'

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
