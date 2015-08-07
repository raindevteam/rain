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

      let wild = channel === 'any' ? true : false;
      for (let key in bot.chans) {
        if (wild || key === channel.toLowerCase()) {
          for (let key in bot.chans[key].users) {
            if (key.toLowerCase() === who.toLowerCase()) return callback(true);
          }
        }
      }
      return callback(false);
    },

    isAdmin(channel, nick) {
      // Todo?
    }

  };
};
