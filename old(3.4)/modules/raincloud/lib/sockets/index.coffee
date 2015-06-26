module.exports = (io) ->
  Console = io.of('/console');

  Console.on 'connection', (socket) ->
    socket.on 'backlog', (data) ->
      Console.emit 'backlog', [
        message:
          from: 'MistaWolf'
          text: 'This is a test message'
      , action:
          from: 'MistaWolf'
          text: 'This is a test action'
      ]

  return io
