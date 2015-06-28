rainUtil = require __rainUtil
bot = __core.bot
async = require('async')

User = require './../models/user'
Channel = require './../models/channel'
_    = require 'lodash'

chan = require './chanHelper'

# resolveAcc (info Object, String, Callback)

resolveAcc = (info, nick, callback) ->
  User.findOne acc: info.acc, (err, user) ->
    if (_.isEmpty user) and info.level > 0
      rainUtil.logf "forecast", "Creating user for " + nick
      newUser = new User()
      newUser.nicks.push(nick)
      newUser.num = 0
      newUser.acc = info.acc
      newUser.save (err) ->
        rainUtil.logf "forecast", "Done creating user for " + nick
        return callback()
    else if !(_.isEmpty user)
      rainUtil.logf "forecast",
        "Pushing nick (" + nick + ") to ACC " + user.acc
      user.nicks.push(nick)
      user.save (err) ->
        rainUtil.logf "forecast",
          "Done pushing nick (" + nick + ") to ACC " + user.acc
        return callback()
    else return callback()

getNsAcc = (bot, nick, callback) ->
  bot.addListener 'raw', getAccInfo = (message) ->
    if message.user == 'NickServ' and
    message.args[1].hasAt(nick + " ->", 0)
      args = message.args[1].after('-> ').split(' ')
      bot.removeListener 'raw', getAccInfo
      return callback acc: args[0], level: args[2]
  bot.say('NickServ', 'acc ' + nick + ' ' + nick)

worker = (task, callback) ->
  User.findOne nicks: task.nick, (err, user) ->
    setTimeout getNsAcc, 1000, __core.bot, task.nick, (info) ->
      if !(_.isEmpty user) # User Exists
        if info.acc != user.acc # Acc does not match
          rainUtil.logf "forecast",
            task.nick + " had conflicting ACC, resolving it"
          _.pull(user.nicks, task.nick)
          user.markModified('nicks')
          user.save (err) ->
            resolveAcc info, task.nick, -> return callback()
        else # Acc matched
          return callback()
      else # User Does not Exist
        resolveAcc info, task.nick, -> return callback()

queue = async.queue(worker, 1)

module.exports =

  resolveNick: (nick, done) ->
    queue.push nick: nick, done

  addChannel: (nick, channel, done) ->
    User.findOne nicks: nick, (err, user) ->
      if user and !(_.find user.channels, (chan) ->
        return chan.tag == channel.lower())
        user.channels.push
          tag: channel.lower()
          name: channel
          stats:
            messages: []
            totalMessages: 0
            longestMessageLength: 0
            averageMessageLength: 0
            averageWordCount: 0
            longestWordCount: 0
        user.save (err) ->
          rainUtil.logf "forecast",
            "Added channel " + channel + " to " + nick + " (" + user.acc + ")"
          if done then return done()
      else
        if done then return done()

  connectedTo: (channel, user) ->
    for key, val of bot.chans
      if channel.lower() == key.lower()
        for key, val of val.users
          if key == user and val != undefined then return true
    return false
