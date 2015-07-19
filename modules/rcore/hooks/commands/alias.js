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
      const aliasArgs = X(_.drop(args).join(' ')).s;
      const params = X(aliasArgs).before('->').clean().s.trim().split(' ');
      const command = X(aliasArgs).after('->').clean().s;
      console.log(params);
      if (!aliasName) {
        respond.say('No alias name entered!');
        return done();
      } if (!command) {
        respond.say('No command to alias entered!');
        return done();
      }
      respond.say('Added command alias: ' + aliasName);
      alias.add(aliasName, {args: params, cmd: command}, done);
    }
  }
};

module.exports = function(module) {
  module.addCommand('alias', alias);
};
