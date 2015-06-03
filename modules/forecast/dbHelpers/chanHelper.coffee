rainUtil = require __rainUtil
async = require 'async'
_  = require 'lodash'

Channel = require './../models/channel'

# worker (task Object, Callback)
#
# Handles task objects holding event, channel, and argument data
# for mongo database

worker = (task, callback) ->
  Channel.findOne name: task.channel, (err, channel) ->
    if !(_.isEmpty channel)
      args = task.args
      switch task.event
        when "start" # channel, nicks
          start(channel, args[0])
        when "reset" # channel
          reset(channel)
        when "push" # channel, nick
          push(channel, args[0])
        when "pop" # channel, nick
          pop(channel, args[0])
        when "nick" # channel, oldnick, newnick
          nick(channel, args[0], args[1])
        when "message" # channel, nick, message
          message(channel, args[0], args[1])
        when "action" # channel, nick, message
          action(channel, args[0], args[1])
      save channel, -> callback()
    else callback()

# queue for database requests
queue = async.queue(worker, 1)

# start (Mongoose Object, String Array)
#
# Takes an array of nicks and then pushes individual nicks to
# usersSeen of channel document if not already there. Also pushes
# to currentUsers and then checks if current user count is higher
# than highestUserCount of channel document.

start = (channel, nicks) ->
  for nick in nicks
    if !(_.contains channel.usersSeen, nick)
      channel.usersSeen.push(nick)
      channel.totalUsers++
    if !(_.contains channel.currentUsers, nick)
      channel.currentUsers.push(nick)
      channel.usersAsOfNow = channel.currentUsers.length
      if channel.usersAsOfNow > channel.highestUserCount
        channel.highestUserCount = channel.usersAsOfNow

# reset (Mongoose Object)
# Resets current user data for channel document

reset = (channel) ->
  channel.currentUsers = []
  channel.usersAsOfNow = 0

# pop (Mongoose Object, String)
# Removes nick from channel document's array and updates usersAsOfNow

pop = (channel, nick) ->
  _.pull channel.currentUsers, nick
  channel.markModified 'currentUsers'
  channel.usersAsOfNow = channel.currentUsers.length

# push (Mongoose Object, String)
#
# Pushes nick to channel document, and updates usersAsOfNow accordingly.
# Also updates usersSeen if not present

push = (channel, nick) ->
  if !(_.contains channel.usersSeen, nick)
    channel.usersSeen.push(nick)
    channel.totalUsers++
  if !(_.contains channel.currentUsers, nick)
    channel.currentUsers.push(nick)
    channel.usersAsOfNow = channel.currentUsers.length
    if channel.usersAsOfNow > channel.highestUserCount
      channel.highestUserCount = channel.usersAsOfNow

# nick (Mongoose Object, String, String)
#
# Updates channel document by pulling from currentUsers oldnick and
# then pushing newnick into the array. If newnick is not present
# in usersSeen, it is pushed.

nick = (channel, oldnick, newnick) ->
  if !(_.contains channel.currentUsers, oldnick)
    return
  _.pull(channel.currentUsers, oldnick)
  channel.currentUsers.push(newnick)
  channel.markModified 'currentUsers'
  if !(_.contains channel.usersSeen, newnick)
    channel.usersSeen.push(newnick)
    channel.totalUsers++

# message (Mongoose Object, String, String)
# Updates channel document's message data (Currently WIP)

message = (channel, nick, message) ->
  channel.messages++

# message (Mongoose Object, String, String)
# Updates channel document's action data (Currently WIP)

action = (channel, nick, message) ->
  channel.actions++

# save (Mongoose Object, Callback)
# Saves channel document

save = (channel, callback) ->
  channel.save (err) ->
    return callback()

class ChanHelper

  # ChanHelper :: constructor ()
  # Initializes connectedChans (Array)

  constructor: () ->
    @connectedChans = []

  # ChanHelper :: connected (String)
  # Checks if bot is connected to channel

  connected: (name) ->
    return @connectedChans.indexOf(name.lower()) > -1

  # ChanHelper :: ready (String)
  # Checks if channel exists in database

  ready: (name) ->
    Channel.findOne name: name, (err, channel) ->
      return !(_.isEmpty channel)

  # ChanHelper :: init (String, Callback)
  # Creates new channel for specified channel name

  init: (name, callback) ->
    if !@connected(name) then @connectedChans.push(name.lower())
    Channel.findOne name: name, (err, channel) ->
      if _.isEmpty channel
        channel = new Channel()
        channel.name = name
        channel.actions = 0
        channel.messages = 0
        channel.totalUsers = 0
        channel.usersAsOfNow = 0
        channel.highestUserCount = 0
        channel.save (err) ->
          rainUtil.logf 'forecast', name + " was initialized in DB"
          if callback then return callback()
      else
        if callback then return callback()

  # ChanHelper :: do (String, String, Args...)
  # Takes database operation requests

  do: (event, name, args...) ->
    if __core.helpers.isFunction args[args.length - 1]
      callback = _.takeRight(args)[0]
      _.dropRight(args) if args.length >  1
      args = [null]     if args.length == 1
    rainUtil.logf 'forecast', event + " -> " + name + ' : ' + args.join ' '
    if !@ready(name) then return
    queue.push event: event, channel: name, args: args, ->
      return callback() if callback

module.exports = ChanHelper
