_ = require 'lodash'

nester = require './nest'
hooks = require './hooks'

module.exports =

  # packageData (String, params Object, Module Object [, String Args])
  # --
  # Packages arguments into a data object. If called for commands, also
  # packages that command's arguments

  packageData: (event, params, parent, args) ->
    data = {}
    data.event = event
    data.hooks = hooks
    data[key] = val for key, val of params
    data.parent = parent
    data.irc = parent.bot.irc
    data.bot = parent.bot
    if args then data.args = args
    return data



  getCommandName: (command) ->
    if __config.preBang
      if command.match(/^!/)
        return X(command.match(/^!(\S+)/)[1]).lower().s
    else
      if X(command).has('!')
        return X(command).before('!').lower().s.trim()
      else null


  getCommandArgs: (command) ->
    args = _.drop(command.split(' '))
    if args
      args = args.join(' ')
      args = nester.fillNests(args)
      return args.split(' ')
    else
      return []
