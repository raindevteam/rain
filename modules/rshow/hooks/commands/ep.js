"use strict";

var eplist;

const ep = {
  help: 'Recommends a random episode, or finds the title from a key (rainbot.info/userguide#ep)',
  action: function(data, respond, done) {
    const args = data.args;
    if (args.length > 2) {
      respond.say('Too many params, try something like \'1 12\'');
    } else if (args[0]) {
      if (!args[1]) {
        respond.say('No episode number entered');
      } else if (isNaN(args[0])) {
        respond.say(args[0] + ' is not a valid param');
      } else if(isNaN(args[1])) {
        respond.say(args[1] + ' is not a valid param');
      } else {
        const season = args[0];
        const episode = args[1].length === 1 ? '0' + args[1] : args[1];
        const result = eplist.getEp('s'+season+'e'+episode);
        if (result) respond.say(result);
        else        respond.say('Couldn\'t find anything');
      }
    } else {
      const ep = eplist.getRandEp();
      console.log(ep);
      respond.say('You should watch ' +
      ep[1] + ' ' + eplist.keyToReadable(ep[0]));
    }
    return done();
  }
};

module.exports = function(module) {
  eplist = require('../../lib/eplist')(module);
  module.addCommand('ep', ep);
};
