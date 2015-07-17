"use strict";

const say = {
  help: 'Repeats arguments immediately (rainbot.info/userguide#say)',
  ASAP: true,
  action(data, respond, done) {
    if (data.args) respond.say(data.args.join(" "));
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('say', say);
};
