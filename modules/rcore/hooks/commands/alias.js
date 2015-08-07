"use strict";

const _ = require('lodash');

const alias = {
  help: 'Lets you alias a/some command(s) (rainbot.info/userguide#alias)',
  ASAP: 'false',
  action: function(data, respond, done) {
    const args = data.args;
    const aliaser = data.bot.alias;
    if (args[0] === 'rm') {
      aliaser.delete(args[1], function() {
        respond.say('Removed Alias: ' + args[1]);
        return done();
      });
    } else {
      aliaser.parseAlias(args.join(" "), function(err, alias) {
        if (err) {
          respond.say(err.message);
          return done();
        }
        respond.say('Added alias: ' + alias.name);
        aliaser.add( alias.name
                  , { params: alias.params, alias: alias.alias }
                  , done );
      });
    }
  }
};

module.exports = function(module) {
  module.addCommand('alias', alias);
};
