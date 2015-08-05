"use strict";

const ircRanks = {
  voice: 1,
  op: 2,
  admin: 3,
  owner: 4
}

const async = require('async');
const _ = require('lodash');

module.exports = function (bot) {
  return {

    connected(channel, who, callback) {
      if (_.isFunction(who)) {
        callback = who;
        who = bot.nick;
      }
      _.forEach(bot.chans, function (chan, key) {
        if (chan.key === channel.toLowerCase()) {
          if (chan.users.hasOwnProperty(who)) return callback(true);
                                         else return callback(false);
        }
      });
    },

    isAdmin(channel, nick) {
      // Todo?
    }

  };
};
