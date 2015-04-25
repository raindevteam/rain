class ResponseHandler
  constructor: () ->
    @responses = []
    @properties = {}

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

  output: (data) ->
    @data = data

  reset: () ->
    @responses = []
    @data      = undefined

module.exports = ResponseHandler
