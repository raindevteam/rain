RespondQueue = require './respondqueue'

class Responder
  constructor: (name) ->
    @name = name
    @responses = []
    @outputData = undefined
    @deftarget = undefined
    @drained = false

  setDefaultTarget: (data) ->
    if data.to then @deftarget = data.to
    if data.channel then @deftarget = data.channel
    if data.nick then @deftarget = data.nick
    if data.newnick then @deftarget = data.newnick
    return @

  handleResponseArgs: (args) ->
    if args.length > 2 then return null
    to = undefined
    response = undefined
    if args[1]
      to = args[0]
      response = args[1]
    else response = args[0]
    if !to then to = @deftarget
    return [to, response]

  say: (args...) ->
    args = @handleResponseArgs(args)
    [to, response] = [args[0], args[1]]
    if !response then return
    @responses.push
      method: 'say'
      to: to
      data: response
    @drained = false
    return @

  action: (args...) ->
    [to, response] = @handleResponseArgs(args)
    if !response then return
    @responses.push
      method: 'action'
      to: to
      data: response
    @drained = false
    return @

  now: () ->
    RespondQueue.push(@responses.pop())
    if @responses.length == 0 then @drained == true

  flush: () ->
    oldResponses = []
    for response in @responses
      oldResponses.push response
      RespondQueue.push(response)
    @drained = true
    @responses = []
    return {output: @outputData, responses: oldResponses}

  flushed: () ->
    return @drained

  output: (data) ->
    @outputData = data

  getResponse: () ->
    return {output: @outputData, responses: @responses}

  reset: () ->
    @responses = []
    @outputData = undefined

module.exports =

  assignResponder: (name) ->
    return new Responder(name)
