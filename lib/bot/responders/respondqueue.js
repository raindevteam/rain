"use strict";

const async = require('async');

var bot;

function worker(response, callback) {
  bot[response.method](response.to, response.data);
  if (callback) callback();
}

const RespondQueue = async.queue(worker, 1);

module.exports = {

  setBot(_bot) {
    bot = _bot;
  },

  push(response) {
    RespondQueue.push(response, null);
  }
};
