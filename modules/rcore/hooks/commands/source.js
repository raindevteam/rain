"use strict";

const source = {
  help: 'Links you to RainBot\'s github (rainbot.info/userguide#source)',
  ASAP: false,
  action: function(data, respond, done) {
    respond.say('https://github.com/Wolfchase/rainbot');
	return done();
  }
};

module.exports = function(module) {
  module.addCommand('source', source);
};