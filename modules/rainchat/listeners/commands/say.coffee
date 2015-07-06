module.exports = (Module) ->

  # say
  # ---

  'say':
    help: 'Repeats arguments immediately (rainbot.info/userguide#say)'
    ASAP: true
    action: (data, respond, done) ->
      respond.say data.args.join(" ") if data.args
      return done()
