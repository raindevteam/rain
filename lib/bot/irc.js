"use strict";

const _ = require('lodash');

module.exports = function(bot) {
  return {
    connected(channel, who) {
      if (!who) who = bot.to;
      Object.keys(bot.chans).forEach(function(val, chan) {
        if (chan.key === channel.toLowerCase())
          if (_.contains(val.users, who)) return true;
      });
      return false;
    }
  };
};
