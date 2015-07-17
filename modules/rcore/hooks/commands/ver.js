"use strict";

const ver = {
  help: 'Shows the bot\'s version (rainbot.info/userguide#version)',
  ASAP: false,
  action: function(data, respond, done) {
    respond.say(data.bot.version);
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('ver', ver);
};
