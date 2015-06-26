config = require(__config)

module.exports =
  ### == info (command) ==
  Returns info about the bot, including version
  ###
  info:
    nest: true
    ASAP: false
    action: (args, respond, done) ->
      info = config.info
      respond.say(
        'I\'m ' + info['description'] +
        ' ' + 'written with ' +
        info['writtenIn'] + ' ' +
        'by ' + info['developer'] +
        ' ' + '[Version: ' +
        info['version'] + ' ' +
        '(' + info['versionName'] + ')]'
        )
      return done()
