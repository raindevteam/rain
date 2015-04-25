class Action
  constructor: (responseHandler, commandHandler, bot) ->
    @commandHandler = commandHandler
    @responseHandler = responseHandler
    @bot = bot

  setResponseProperties: (properties) ->
    console.log 'setting response props'
    @responseHandler.setProperties(properties)

  triggered: (action, callback) ->
    console.log 'now'
    if action.trigger @responseHandler.properties
      console.log 'hi'
      return callback(true)

  actionRespond: (responses) ->
      for response in responses
        if (response.length > 250) then return # temp
        @bot[response.method](response.to, response.res)

  fireTrigger: (trigger) ->
    @responseHandler.reset()
    self = this
    action = trigger.action @responseHandler, () ->
      self.actionRespond self.responseHandler.responses

  fireCommand: (commandText) ->
    if !commandText then return
    @responseHandler.reset()
    commands = @commandHandler.getCommands(commandText)
    self = this
    @commandHandler.run commands, @responseHandler, (lastAction) ->
      self.actionRespond self.responseHandler.responses

module.exports = Action
