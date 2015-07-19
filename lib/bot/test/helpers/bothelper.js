/* globals __config */
"use strict";

const Bot = require('./../../bot');

module.exports = {

  Bot,

  makeBot() {
    return new Bot(__config.server, __config.nick, {
      autoConnect: false
    });
  }
};
