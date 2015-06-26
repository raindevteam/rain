irc = require 'irc'

class bot extends irc.Client
  constructor: (server, nick, options) ->
    super server, nick, options

  load: (callback) ->
    console.log 'loaded things'
    return callback()

  listen: (callback) ->
    console.log 'listening'
    return callback()

module.exports = bot
