rainUtil = require __rainUtil
config   = require __config
async    = require 'async'
_        = require 'lodash'

rainUtil.newLog 'forecast', __dirname + '/../log.txt'

User     = require './../models/user'
Channel  = require './../models/channel'

UserHelper = require './../dbHelpers/userHelper'
ChanHelper = require './../dbHelpers/chanHelper'

user = new UserHelper()
chan = new ChanHelper()

module.exports =
  update:
    event: 'names'
    trigger: (message) -> return true
    action: (respond, done) ->
      nicks = []
      for nick, rank of respond.nicks
        nicks.push nick
      chan.init respond.channel, ->
        chan.do "reset", respond.channel, ->
          chan.do "start", respond.channel, nicks
          async.eachSeries nicks, ((nick, next) ->
            user.resolveNick nick, respond.channel, ->
              user.addChannel nick, respond.channel, next
          ), (err) -> return done()

  updateOnJoin:
    event: 'join'
    trigger: (message) -> return true
    action: (respond, done) ->
      if respond.nick == config.nick then return done()
      chan.do "push", respond.channel, respond.nick
      user.resolveNick respond.nick, respond.channel, ->
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
        if chan.connected(channel)
          chan.do "pop", channel, respond.nick
      done()

  updateOnNickChange:
    event: 'nick'
    trigger: (message) -> return true
    action: (respond, done) ->
      for channel in respond.channels
        if chan.connected(channel)
          chan.do "nick", channel, respond.oldnick, respond.newnick
      user.resolveNick respond.newnick, respond.channel, -> # (created) ->
        for channel in respond.channels
          if chan.connected(channel)
            user.addChannel respond.newnick, channel, null
        done()

  actionUpdate:
    event: 'action'
    trigger: (message) -> return true
    action: (respond, done) ->
      User.findOne nicks: respond.from,
      'channels.name': respond.to, (err, user) ->
        if !(_.isEmpty user)
          channel = respond.to
          message = respond.text

          # Update Channel's recorded messages
          chan.do "action", channel, respond.from, respond.text

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
      User.findOne nicks: respond.from,
      'channels.tag': respond.to.lower(), (err, user) ->
        if !(_.isEmpty user)
          console.log respond.msg
          channel = respond.to
          message = respond.text

          # Update Channel's recorded messages
          chan.do "message", channel, respond.from, respond.text

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
