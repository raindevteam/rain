class ResponseHandler
  constructor: () ->
    @responses = []
    @properties = {}

  setProperties: (properties) ->
    @properties = {}
    for name, prop of properties
      console.log name + " : " + prop
      @properties[name] = prop

  say: (response) ->
    console.log @properties
    @responses.push
      method: 'say'
      to: @properties.to
      res: response

  action: (to, response) ->
    @responses.push
      method: 'action'
      to: to
      res: response

  output: (data) ->
    @data = data

  reset: () ->
    @responses = []
    @data      = undefined

module.exports = ResponseHandler
