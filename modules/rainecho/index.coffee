module.exports = (Module) ->
  # RainEcho, a beginner friend module showing how to Create
  # the simplest of modules.

  RainEcho = new Module('RainEcho')

  # echo (command)
  # Repeats its arguments

  RainEcho.addCommand 'echo',
    ASAP: false
    action: (data, respond, done) ->
      respond.say data.args.join ' '
      return done()

  RainEcho.addCommand 'say',
    ASAP: true
    action: (data, respond, done) ->
      respond.say data.args.join ' '
      return done()

  return RainEcho
