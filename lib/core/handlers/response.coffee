class ResponseHandler
  constructor: (bot) ->
    @responses = []
    @properties = {}
    @bot = bot

  setProperties: (properties) ->
    @properties = {}
    for name, prop of properties
      @properties[name] = prop

  say: (response) ->
    @responses.push
      method: 'say'
      to: @properties.to
      res: response

  action: (response) ->
    @responses.push
      method: 'action'
      to: @properties.to
      res: response

  respond: () ->
    for response in @responses
      if (response.length > 250) then return # temp
      @bot[response.method](response.to, response.res)

  output: (data) ->
    @data = data

  reset: () ->
    @responses = []
    @data      = undefined

module.exports = ResponseHandler
