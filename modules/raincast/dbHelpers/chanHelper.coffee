rainUtil = require __rainUtil
async = require 'async'
_  = require 'lodash'

Channel = require './../models/channel'

# Keep track of connected channels locally
# (I think the bot already keeps this in check though, must double check)

connectedChans = []

# connected (String, callback)
# Checks if bot is connected to channel

connected = (name, callback) ->
  console.log 'checking'
  for channel in connectedChans
    console.log name.lower() + ' against ' + channel.lower()
    if channel.lower() == name.lower() then return callback(true)
  return callback(false)

# worker (task Object, Callback)
#
# Handles task objects holding event, channel, and argument data
# for mongo database

worker = (task, callback) ->
  Channel.findOne tag: task.tag, (err, channel) ->
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
    else
      bucket.push task, null
      callback()

# queue for database requests
queue = async.queue(worker, 1)

# bucket for db operations that couldn't be performed due
# to channel initialization
bucket = async.queue(worker, 1)
bucket.drain = ->
  queue.resume()
  bucket.pause()

# Flushes bucket if any tasks present, otherwise resumes regular queue
flushBucket = ->
  if bucket.length() > 1
    rainUtil.logf 'forecast', "Flushing bucket"
    bucket.resume()
  else queue.resume()

# start (Mongoose Object, String Array)
#
# Takes an array of nicks and then pushes individual nicks to
# usersSeen of channel document if not already there. Also pushes
# to currentUsers and then checks if current user count is higher
# than highestUserCount of channel document.

start = (channel, nicks) ->
  queue.pause() # Ensure that after we finish setting up the channel
                # we don't process anything before the bucket is flushed
  for nick, rank of nicks
    if !(_.contains channel.usersSeen, nick)
      channel.usersSeen.push(nick)
      channel.totalUsers++
    if !(_.contains channel.currentUsers, nick)
      channel.currentUsers.push(nick)
      channel.usersAsOfNow = channel.currentUsers.length
      if channel.usersAsOfNow > channel.highestUserCount
        channel.highestUserCount = channel.usersAsOfNow
  flushBucket() # Flush any operations we couldn't do

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
  channel.currentUsers = _.without(channel.currentUsers, oldnick)
  channel.currentUsers.push(newnick)
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

module.exports =

  # export connected function

  connected: connected

  # getName (String)
  # Returns original name of channel

  getName: (name) ->
    for channel in connectedChans
      if channel.lower() == name.lower() then return channel

  # init (String, Callback)
  # Creates new channel for specified channel name

  init: (name, callback) ->
    connected name, (result) ->
       if result == false then connectedChans.push(name)
    Channel.findOne name: name, (err, channel) ->
      if _.isEmpty channel
        channel = new Channel()
        channel.tag = name.lower()
        channel.name = name
        channel.actions = 0
        channel.messages = 0
        channel.totalUsers = 0
        channel.usersAsOfNow = 0
        channel.highestUserCount = 0
        channel.save (err) ->
          rainUtil.logf 'forecast', name + " was initialized in DB"
          return callback() if callback
      else
        return callback() if callback

  # do (String, String, Args...)
  # Takes database operation requests

  do: (event, name, args...) ->
    if __core.helpers.isFunction args[args.length - 1]
      callback = _.takeRight(args)[0]
      _.dropRight(args) if args.length >  1
      args = [null]     if args.length == 1
    rainUtil.logf 'forecast', event + " -> " + name + ' : ' + args.join ' '
    queue.push event: event, tag: name.lower(), args: args, ->
      return callback() if callback
