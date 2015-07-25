/* globals __config */
"use strict";

const Bot = require('./../../lib/bot');

module.exports = {

  Bot,

  makeBot() {
    return new Bot(__config.server, __config.nick, {
      autoConnect: false
    });
  }
};
