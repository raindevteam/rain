"use strict";

module.exports = function(Module, done) {
  const rainChat = new Module("RainChat", "rchat");
  rainChat.load(__dirname + '/hooks/commands');
  return done(rainChat);
};
