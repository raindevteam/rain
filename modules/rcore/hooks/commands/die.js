"use strict";

const die = {
  help: 'Best you don\'t use it, unless the bot really needs to go bye bye',
  ASAP: true,
  action: function(data, respond, done) {
    process.exit(0);
  }
};

module.exports = function(module) {
  module.addCommand('die', die);
};
