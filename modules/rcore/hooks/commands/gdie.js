"use strict";

const gdie = {
  help: 'Disconnects first and then terminates the bot (A graceful death)',
  ASAP: true,
  whitelist: ['batpony', 'Medea', 'PinkDawn', 'Eventide', 'Dustrunner'],
  action: function(data, respond, done) {
    let bye = data.args.length ? data.args.join(' ') : 'Bye from RainBot';
    data.bot.disconnect(bye, function() {
      process.exit(0);
    });
  }
};

module.exports = function(module) {
  module.addCommand('gdie', gdie);
};
