async = require('async')

BOT = undefined

worker = (response, callback) ->
  BOT[response.method](response.to, response.data)
  return callback() if callback

RespondQueue = async.queue(worker, 1)

module.exports =

  setBot: (bot) ->
    BOT = bot

  push: (response) ->
    RespondQueue.push response, () ->
