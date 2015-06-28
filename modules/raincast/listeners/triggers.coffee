rainUtil = require __rainUtil
config   = require __config
async    = require 'async'
_        = require 'lodash'

rainUtil.newLog 'forecast', __dirname + '/../log.txt'

User     = require './../models/user'
Channel  = require './../models/channel'

user = require './../dbHelpers/userHelper'
chan = require './../dbHelpers/chanHelper'

module.exports =
  updateOnNames:
    event: 'names'
    trigger: (message) -> return true
    action: (respond, done) ->
      chan.init respond.channel, ->
        chan.do "reset", respond.channel, ->
          chan.do "start", respond.channel, respond.nicks
          async.forEachOfSeries respond.nicks, ((rank, nick, next) ->
            user.resolveNick nick, ->
              user.addChannel nick, respond.channel, next
          ), (err) -> return done()

  updateOnJoin:
    event: 'join'
    trigger: (message) -> return true
    action: (respond, done) ->
      if respond.nick == config.nick then return done()
      chan.do "push", respond.channel, respond.nick
      user.resolveNick respond.nick, ->
        user.addChannel respond.nick, respond.channel, done

  updateOnPart:
    event: 'part'
    trigger: (message) -> return true
    action: (respond, done) ->
      chan.do "pop", respond.channel, respond.nick, done

  updateOnQuit:
    event: 'quit'
    trigger: (message) -> return true
    action: (respond, done) ->
      for channel in respond.channels
        rainUtil.logf 'forecast', "Inspecting " + channel + " for nick " + respond.nick
        rainUtil.logf 'forecast', "Channel connected -> " + chan.connected(channel)
        chan.connected channel, (result) ->
          if result == true
            chan.do "pop", channel, respond.nick
      return done()

  updateOnNickChange:
    event: 'nick'
    trigger: (message) -> return true
    action: (respond, done) ->
      async.forEachSeries respond.channels, ((channel, next) ->
        if user.connectedTo(channel, respond.newnick)
          chan.do "nick", channel, respond.oldnick, respond.newnick, next
      ), (err) ->
      user.resolveNick respond.newnick, -> # (created) ->
        async.forEachSeries respond.channels, ((channel, next) ->
          if user.connectedTo(channel, respond.newnick)
            chan.connected channel, (result) ->
              if result == true
                rainUtil.logf 'forecast',
                  respond.newnick + ' is connected to ' + chan.getName(channel)
                user.addChannel respond.newnick, chan.getName(channel), next
          else
            next()
        ), (err) ->
          return done()

  actionUpdate:
    event: 'action'
    trigger: (message) -> return true
    action: (respond, done) ->
      # Update Channel's recorded messages
      chan.do "action", channel, respond.from, respond.text

      User.findOne nicks: respond.from,
      'channels.name': respond.to, (err, user) ->
        if !(_.isEmpty user)
          channel = respond.to
          message = respond.text

          # User stats
          chanStats =
            (_.find user.channels, (chan) -> return chan.name == channel).stats

          chanStats.actions.push(message)
          if (chanStats.actions.length > 1000) then chanStats.actions.shift()

          user.save (err) -> return done()

  messageUpdate:
    event: 'message'
    fireOnCommand: true
    trigger: (message) -> return true
    action: (respond, done) ->
      # Update Channel's recorded messages
      chan.do "message", respond.to, respond.from, respond.text

      User.findOne nicks: respond.from,
      'channels.tag': respond.to.lower(), (err, user) ->
        if !(_.isEmpty user)
          channel = respond.to
          message = respond.text

          # User stats
          chanStats =
            (_.find user.channels, (chan) -> return chan.name == channel).stats

          # Calculate average message length
          sum = chanStats.averageMessageLength * chanStats.messages.length
          chanStats.averageMessageLength =
            Math.round (sum + message.length) / (chanStats.messages.length + 1)

          # Update longest message length
          if (message.length > chanStats.longestMessageLength)
            chanStats.longestMessageLength = message.length

          words = message.split(' ').length

          # Update Longest Word length
          if (words > chanStats.longestWordCount)
            chanStats.longestWordCount = words

          # Update Average Word Length
          sum = chanStats.averageWordCount * chanStats.messages.length
          chanStats.averageWordCount =
            Math.round (sum + words) / (chanStats.messages.length + 1)

          # Push message to logs
          chanStats.messages.push(message)
          if (chanStats.messages.length > 1000) then chanStats.messages.shift()

          # Save user data
          user.save (err) ->
            return done()
