User = require('./../models/user')
Channel = require('./../models/channel')
_    = require('lodash')

channelQuery = (args, respond, done) ->
  Channel.findOne name: chan, (err, channel) ->
    switch args[0]
      when "users"
        if !args[1]
          respond.say "There are " + channel.usersAsOfNow + " users in " +
          channel.name + " right now"
          return done()
        switch args[1]
          when "now"
            respond.say "There are " + channel.usersAsOfNow + " users in " +
            channel.name + " right now"
            return done()
          when "highest"
            respond.say "Highest user count has been " + channel.highestUserCount +
            " for " + channel.name
            return done()
          when "total"
            respond.say "Total users seen on " + channel.name + " has been " + channel.totalUsers
            return done()
          else
            respond.say "Found nothing for query: " + args[1]
            return done()
      when "messages"
        respond.say "Total messages for " + channel.name + " has been " + channel.messages
        return done()
      else
        respond.say "Found nothing for query: " + args[0]
        return done()

userQuery = (args, respond, done) ->
  chan =
    (_.find user.channels, (chan) -> return chan.name == respond.to)
  if !chan
    respond.say "User doesn't have stats for this channel"
    return done()
  chanStats = chan.stats
  if args.length == 1
    if chanStats.messages.length == 0
      respond.say "User has no messages recorded on this channel"
      return done()
    rand_index = Math.floor(Math.random() * chanStats.messages.length)
    respond.say "(" + args[0] + ") " + '"' + chanStats.messages[rand_index] + '"'
    return done()
  switch args[1]
    when "average"
      respond.say "Average message length has been " +
      chanStats.averageMessageLength
      return done()
    when "longest"
      respond.say "Longest message has been " +
      chanStats.longestMessageLength + " in length"
      return done()
    when "total"
      respond.say "User has a total of " +
      chanStats.totalMessages + " messages recorded"
      return done()
    when "word"
      respond.say "Average word count is " +
      chanStats.averageWordCount
      return done()
    else
      respond.say "Found nothing for query: " + args[1]
      return done()

module.exports =
  q:
    action: (args, respond, done) ->
      if args.length == 0 then return done()
      User.findOne nicks: args[0], (err, user) ->
        if (_.isEmpty user)
          if args[0][0] == '#'
            chan = args[0]
            args = _.drop(args)
            channelQuery(args, chan, respond, done)
          else if (args[0][0] == "$")
            qc = args[0]
            args = _.drop(args)
            quickCommand(args, qc, respond, done)
        else
          userQuery(args, respond, done)
