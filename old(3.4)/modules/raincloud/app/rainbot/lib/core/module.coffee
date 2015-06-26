core       = require(__core)
events     = core.events
async      = require('async')

class ResponseHandler
  constructor: () ->
    @responded = false
    @responses = []
    
  say: (to, response) ->
    @responded = true
    @responses.push({
      method : 'say'
      to  : to
      res : response
    })
  
  action: (to, response) ->
    @responded = true
    @responses.push({
      method : 'action'
      to  : to
      res : response
    })
    
  output: (data) ->
    @data = data
    
  reset: () ->
    @responses = []
    @data      = undefined


class Module
  constructor: (name) ->
    @name = name
    @fired = undefined
    @resHandler = new ResponseHandler()
    @listeners = {}
    for event in core.events
      @listeners[event] = []
  
  setBot: (bot) ->
    @bot = bot
    
  addListeners: (listeners) ->
    for name,listener of listeners
      @listeners[listener.event].push(listener)
      
  fire: (event, args, handler) ->
    @resHandler.reset()
    self = this
    async.eachSeries @listeners[event], ((list, next) ->
      action = list.action.apply(undefined, args)
      if (action.trigger)
        self.fired = list.name
        action.fire self.resHandler, () ->
          return handler
            responses : self.resHandler.responses
            output    : self.resHandler.data
            nest      : if list.nest then list.nest.name or null
            ASAP      : if list.nest then list.nest.ASAP or false
      else return next()
    ), (err) ->
      return handler({})

module.exports = Module