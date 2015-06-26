bot = undefined

class ResponseHandler
  constructor: (bot) ->
    @responses = []
    @properties = {}

  setProperties: (properties) ->
    @properties = {}
    for name, prop of properties
      @[name] = prop

  say: (response) ->
    @responses.push
      method: 'say'
      to: @to
      res: response

  action: (response) ->
    @responses.push
      method: 'action'
      to: @to
      res: response

  respond: () ->
    for response in @responses
      if (response.length > 250) then return # temp
      self = @
      setImmediate(() -> bot[response.method](response.to, response.res))

  output: (data) ->
    @data = data

  reset: () ->
    @responses = []
    @data      = undefined

module.exports = (_bot) ->
  bot = _bot
  return ResponseHandler
