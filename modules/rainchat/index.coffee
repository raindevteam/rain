module.exports = (Module) ->
  RainChat = new Module('RainChat')
  RainChat.addCommands(__dirname + '/listeners/commands')
  return RainChat
