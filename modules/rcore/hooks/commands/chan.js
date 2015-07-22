"use strict";

const chan = {
  help: 'Allows the bot to join and leave channels (rainbot.info/userguide#chan)',
  ASAP: false,
  action: function(data, respond, done) {
    const cmd = data.args.join(' ');

    let join  = cmd.match(/\bjoin\s*(.*(?= leave)|(?=leave)|.*)/i);
    let leave = cmd.match(/\bleave\s*(.*(?= join)|(?=join)|.*)/i);

    if (join && join[1] && join[1].match(new RegExp(' '))) {
      respond.say('Too many arguments to command join');
      return done();
    } if (leave && leave[1] && leave[1].match(new RegExp(' '))) {
      respond.say('Too many arguments to command leave');
      return done();
    }

    if (join && join[1]) {
      join = join[1].split(':').join(' ');
      if (join[0] !== '#') {
        respond.say('Invalid channel format for join');
        return done();
      }
    } else if (join) { join = null; }

    if (leave && leave[1]) {
      leave = leave[1].split(':').join(' ');
      if (leave[0] !== '#') {
        respond.say('Invalid channel format for leave');
        return done();
      }
    } else if (leave) leave = data.to;

    if (join) {
      data.bot.gate(join);
      respond.say('Joined ' + join).now();
    } if (leave) {
      data.bot.part(leave);
      respond.say("Left " + leave).now();
    }
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('chan', chan);
};
