module.exports =
  # Configuration options for RainBot v4 (Pyrelight)

  # server -- address to IRC server
  #  port  -- port number

  server: process.env.IRC_SERVER or 'irc.canternet.org'
  port: process.env.IRC_PORT or 6667

  #   nick   -- Nick to be displayed
  # realName -- Your 'realName'
  # userName -- (Optional) Your username

  nick: process.env.IRC_NICK or 'HailBot'
  realName: process.env.IRC_REAL or 'RainBotExp'
  userName: process.env. IRC_USER or 'RainBot'

  # If your server supports NickServ, enter your password here
  # to be authed when the bot connects or set it to false or undefined

  nsPassword: process.env.IDENTIFY_PASSWORD or undefined

  # channels: (The channels to join on connect)

  channels: [ process.env.CHANNELS or '#snowybottest' ]

  # logging: A string specifying logging modes, all modes include:
  #   x: No logging      a: Log everything
  #   i: Log info        e: Log errors
  #   w: log warnings    f: Allow file logging (regular logs are file logged)
  #   F: Only file log when logf is used (regular logs are not file logged)
  #
  #     To log all output (and include file logging) use '+af'
  #     To exclude output (as an example, no info logging) use '+af-i'

  logging: true

  # debug: Set to true if you want debug logs from the backbone IRC library

  debug: true

  # Option to set the bang for commands in front of the command names
  # True: '!cmd'
  # False: 'cmd!'

  preBang: false

  # Universal Whitelist for commands, only nicks listed here
  # can use any commands from core
  #
  # If you want to whitelist module commands, use the module class method
  # 'addToWhitelist(name, function)'

  whitelist:
    Wolf: [
      'tuner'
      'sleep'
      'wake'
    ]
    Fox: [
      'wake'
    ]
