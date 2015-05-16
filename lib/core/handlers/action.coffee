class Action
  constructor: (responseHandler, commandHandler) ->
    @commandHandler = commandHandler
    @responseHandler = responseHandler
    @properties = {}

  setResponseProperties: (properties) ->
    for prop, value of properties
      @properties[prop] = value

  onCommandTrigger: (action, callback) ->
    if action.fireOnCommand
      if action.trigger @properties then return callback(true)

  triggered: (action, callback) ->
    if action.trigger @properties then return callback(true)

  handle: (action, callback) ->
    if action.trigger @properties
      @fireTrigger(action)
    callback(null)

  fireTrigger: (trigger) ->
    if !trigger then return
    responseHandler = new @responseHandler()
    responseHandler.setProperties(@properties)
    self = this
    action = trigger.action responseHandler, () ->
      responseHandler.respond()

  fireCommand: (commandText) ->
    if !commandText then return
    responseHandler = new @responseHandler()
    responseHandler.setProperties(@properties)
    commands = @commandHandler.getCommands(commandText)
    @commandHandler.run commands, responseHandler, (lastAction) ->
      console.log 'returned from firing commands'
      responseHandler.respond()

module.exports = Action
