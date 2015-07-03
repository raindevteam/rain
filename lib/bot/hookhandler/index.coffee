hashmap = require 'hashmap'
async = require 'async'

HookRespond = require './../responders/hookrespond'
helpers = require './helpers'
nester = require './nest'
hooks = require './hooks'
parser = require './../parser'

extractTriggers = (module) ->
  for event, triggers of module.triggers
    for trigger in triggers
      hooks.setTrigger(trigger)

extractCommands = (module) ->
  module.commands.forEach (command, name) ->
    hooks.setCommand(name, command)

execute = (command, params, callback) ->
  commandName = helpers.getCommandName(command)
  commandArgs = helpers.getCommandArgs(command)
  return callback(null) if !commandName
  # if alias.isAlias(commandName)
  #  commands = alias.get(commandName)
  #  run commands, params, (responder) -> return callback(responder)
  if true
    hook = hooks.getCommand(commandName)
    if !hook
      return callback(null)
    if !hook.listening or !hook.parent.enabled
      return callback(null)
    if hook.whitelist and !_.contains(hook.whitelist, params.from)
      return callback(null)
    data = helpers.packageData('command', params, hook.parent, commandArgs)
    responder = HookRespond
      .assignResponder(commandName)
      .setDefaultTarget(data)
    hook.action data, responder, () ->
      if hook.ASAP then response = responder.flush()
      else              response = responder.getResponse()
      nester.nest(hook.name, response)
      return callback(responder)


# run (String Array, params Object, Boolean, Callback)
# --
# Runs a set of commands, but only flushes the last responder.
# Also flushes the responder if hook has ASAP set to true. The alias
# argument is used when this function is called recursively from
# the execute function when an alias command is ran.

run = (commands, params, alias, done) ->
  async.forEachOfSeries commands, ((command, i, next) ->
    execute command, params, (responder) ->
      return next() if !responder
      if i + 1 == commands.length and !responder.flushed() and !alias
        responder.flush()
        return done(responder) if done
      else
        next()
  ), (err) ->
    if err
      rainlog.err 'Received error while running commands, err: ' + err.message
      rainlog.err 'Was aliased set of commands: ' + alias

module.exports =

  extractHooks: (module) ->
    extractTriggers(module)
    extractCommands(module)

  execute: execute

  fire: (event, params, callback) ->
    onCommandOnly = false
    if event == 'message' and parser.isCommand(params.text)
      run parser.getCommands(params.text), params, false, null
      onCommandOnly = true
    async.each (hooks.getTriggers)[event], ((hook) ->
      if onCommandOnly then return if !hook.fireOnCommand
      if hook.trigger(params)
        data = helpers.packageData(event, params, hook.parent)
        responder = HookRespond
          .assignResponder(hook.name)
          .setDefaultTarget(data)
        hook.action data, responder, () -> responder.flush()
    ), (err) ->
