module.exports = (Module) ->
  RainCopy = new Module('RainCopy')
  # RainCopy.getApiFrom 'RainEcho', (api) ->
  #   respond = RainCopy.responder()
  #   api.on 'echo', (data) -> respond.say()
  return RainCopy
