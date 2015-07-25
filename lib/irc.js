"use strict";

const async = require('async');
const _ = require('lodash');

module.exports = function(bot) {
  return {

    connected(channel, who, callback) {
      if (_.isFunction(who)) {
        callback = who;
        who = bot.nick;
      }
      _.forEach(bot.chans, function(chan, key) {
        if (chan.key === channel.toLowerCase()) {
          if (chan.users.hasOwnProperty(who)) return callback(true);
          else                                return callback(false);
        }
      });
    }

  };
};
