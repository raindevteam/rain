"use strict";

function strhash(str) {
  var hash = 0;
  for (let letter of str) {
    hash = Math.abs(letter.charCodeAt(0) +
    (hash << 6) + (hash << 16) - hash);
  }
  return hash;
}

const hash = {
  help: 'Hashes a string, or your nick if called by itself (rainbot.info/userguide#hash)',
  action(data, respond, done) {
    if (data.args[0]) {
      respond.say(strhash(data.args[0]));
    } else {
      respond.say(strhash(data.from));
    }
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('hash', hash);
};
