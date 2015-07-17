"use strict";

const echo = {
  help: 'Repeats arguments (rainbot.info/userguide#echo)',
  action(data, respond, done) {
    if (data.args) respond.say(data.args.join(' '));
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('echo', echo);
};
