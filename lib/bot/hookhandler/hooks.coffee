hashmap = require 'hashmap'

events = require './../events'

CommandHooks = new hashmap()
TriggerHooks = {}

for event in events
  TriggerHooks[event] = []

module.exports =

  setCommand: (name, command) ->
    CommandHooks.set(name, command)

  setTrigger: (trigger) ->
    TriggerHooks[trigger.event].push trigger

  getCommand: (name) ->
    return CommandHooks.get(name)

  getTrigger: (name) ->
    for event, triggers in TriggerHooks
      for trigger in triggers
        if trigger.name == name then return trigger

  getTriggers: () -> return TriggerHooks
