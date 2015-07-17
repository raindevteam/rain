"use strict";

const help = {
  help: 'Displays a help link (rainbot.info/userguide#help)',
  ASAP: false,
  action: function(data, respond, done) {
    if (data.args[0]) {
      const hook = data.hooks.getCommand(data.args[0]);
      if (!hook) respond.say('Don\'t know anything about that');
      else       respond.say(hook.help);
    } else {
      respond.say('rainbot.info/userguide');
    }
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('help', help);
};
