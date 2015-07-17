"use strict";

const _ = require('lodash');

const alias = {
  help: 'Lets you alias a/some command(s) (rainbot.info/userguide#alias)',
  ASAP: 'false',
  action: function(data, respond, done) {
    const args = data.args;
    const alias = data.bot.alias;
    if (args[0] === 'rm') {
      alias.delete(args[1], function() {
        respond.say('Removed Alias: ' + args[1]);
        return done();
      });
    } else {
      const aliasName = args[0];
      const command = _.drop(args).join(' ');
      if (!aliasName) {
        respond.say('No alias name entered!');
        return done();
      } if (!command) {
        respond.say('No command to alias entered!');
        return done();
      }
      respond.say('Added command alias: ' + aliasName);
      alias.add(aliasName, command, done);
    }
  }
};

module.exports = function(module) {
  module.addCommand('alias', alias);
};
