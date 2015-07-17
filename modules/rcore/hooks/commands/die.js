"use strict";

const die = {
  help: 'Best you don\'t use it, unless the bot really needs to go bye bye',
  ASAP: true,
  whitelist: ['MistaWolf', 'PinkDawn', 'Eventide',
              'Dustrunner', 'King_Sombra'],
  action: function(data, respond, done) {
    respond.say('No! Not like...').now();
    process.exit(0);
  }
};

module.exports = function(module) {
  module.addCommand('die', die);
};
