class Hook

  # Hook :: constructor (String, String, Module Class, body Object)
  #
  # Creates a new hook from given object, class is really
  # more of a container

  constructor: (name, type, parent, object) ->
    @name = name
    @parent = parent
    @type = type
    @action = object.action
    @listening = true
    if type == 'trigger'
      @event = object.event
      @fireOnCommand = object.fireOnCommand
      @trigger = object.trigger
    else if type == 'command'
      @ASAP = object.ASAP
      @help = object.help

module.exports = Hook
