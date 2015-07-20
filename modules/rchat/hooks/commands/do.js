"use strict";

const _do = {
  help: 'Does an action (rainbot.info/userguide#say)',
  ASAP: true,
  action(data, respond, done) {
    if (data.args) respond.action(data.args.join(" "));
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('do', _do);
};
