Console = undefined

module.exports = (Module, io) ->
  Console = io.of('/console')
  return {
    message:
      event: 'message'
      trigger: (message) -> return true
      action: (respond, done) ->
        Console.emit 'message',
          text: respond.text
          from: respond.from
          chan: respond.to

        if !console[respond.to]
           console[respond.to] = {}
           console[respond.to].messages = []

        console[respond.to].messages.push
          message:
            from: respond.from
            text: respond.text

        if console[respond.to].messages.length > 100
          console[respond.to].messages.shift();

        return done()
  }
