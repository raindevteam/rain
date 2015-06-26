# Global
User  = require(__models + '/user.coffee')
debug = require(__debug)

# Local
async = require('async')
_     = require('lodash')
log   = __dirname + '/../log.txt'

errHandler = (err) ->
  if err
    debug.err err
    debug.errf log, err
    return true
  else false

resolve = (info, nick, callback) ->
  User.findOne accName: info.acc, (err, user) ->
    console.log 'searched for acc'
    console.log info.level
    if (_.isEmpty user) and info.level > 0
      console.log 'creating new user'
      newUser = new User()
      newUser.nicks.push(nick)
      newUser.accName = info.acc
      newUser.save (err) ->
        return callback()
    else if !(_.isEmpty user)
      user.nicks.push(nick)
      user.save (err) ->
        return callback()
    else return callback()

accRegex = undefined

attachNsListener = () ->
  bot.addListener 'raw', getAccInfo = (message) ->
    if message.user == 'NickServ' and
    accRegex.exec(message.args[1])
      args = message.args[1].after('-> ').split(' ')
      bot.removeListener 'raw', getAccInfo
      return find
        acc: args[0]
        level: args[2]

getNsAcc = (bot, nick, callback) ->
  accRegex = new RegExp("^" + nick + " ->")
  process.nextTick(() ->
    bot.say('MistaWolf', 'spamming on purpose')
    bot.say('MistaWolf', 'sent ns')
    )

find = (info) ->
  if !(_.isEmpty user) and info.acc != user.accName
    _.pull(user.nicks, nick)
    user.save (err) ->
      resolve info, nick, () ->
  else if _.isEmpty user
    resolve info, nick, () ->

module.exports =
  update:
    event: 'names'
    trigger: (message) -> return false
    action: (respond, done) ->
      nicks = []
      for nick, rank of respond.nicks
        nicks.push nick
      async.each nicks, ((nick, callback) ->
        console.log nick
        User.findOne nicks: nick, (err, user) ->
          getNsAcc __core.bot, nick
          callback()
      ), (err) ->
        done()
