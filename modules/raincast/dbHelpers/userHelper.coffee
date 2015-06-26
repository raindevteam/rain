rainUtil = require __rainUtil

User = require './../models/user'
_    = require 'lodash'

getNsAcc = (bot, nick, callback) ->
  bot.addListener 'raw', getAccInfo = (message) ->
    if message.user == 'NickServ' and
    message.args[1].hasAt(nick + " ->", 0)
      args = message.args[1].after('-> ').split(' ')
      bot.removeListener 'raw', getAccInfo
      return callback acc: args[0], level: args[2]
  bot.say('NickServ', 'acc ' + nick + ' ' + nick)

resolveAcc = (info, nick, channel, callback) ->
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

class UserHelper
  constructor: () ->

  resolveNick: (nick, channel, next) ->
    User.findOne nicks: nick, (err, user) ->
      setTimeout getNsAcc, 1000, __core.bot, nick, (info) ->
        if !(_.isEmpty user) # User Exists
          if info.acc != user.acc # Acc does not match
            rainUtil.logf "forecast",
              nick + " had conflicting ACC, resolving it"
            _.pull(user.nicks, nick)
            user.markModified('nicks')
            user.save (err) ->
              resolveAcc info, nick, channel, -> return next()
          else # Acc matched
            return next()
        else # User Does not Exist
          resolveAcc info, nick, channel, -> return next()

  addChannel: (nick, channel, done) ->
    User.findOne nicks: nick, (err, user) ->
      if user and !(_.find user.channels, (chan) -> return chan.name == channel)
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

module.exports = UserHelper
