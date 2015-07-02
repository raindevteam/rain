module.exports = (Module) ->
  testcmd4:
    action: (data, respond, done) ->
      respond.say 'done'
      return done()
