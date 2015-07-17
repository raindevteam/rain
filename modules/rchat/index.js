"use strict";

module.exports = function(Module) {
  const rainChat = new Module("RainChat", "rchat");
  rainChat.load(__dirname + '/hooks/commands');
  return rainChat;
};
