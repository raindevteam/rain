config = require(__config)

module.exports =
  ### == help (command) ==
  Returns link set in config.
  If nested, only returns link
  ###
  help:
    nest: true
    ASAP: false
    action: (args, respond, done) ->
      respond.say 'Have a link: ' + config.helpLink
      respond.output config.helpLink
      return done()
