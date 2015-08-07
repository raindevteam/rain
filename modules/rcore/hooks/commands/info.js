"use strict";

const info = {
  help: 'Information on running build (rainbot.info/userguide#info)',
  ASAP: false,
  action: function(data, respond, done) {
    respond.say('https://github.com/Wolfchase/rainbot/releases/tag/v0.8.0-pre');
	return done();
  }
};

module.exports = function(module) {
  module.addCommand('info', info);
};