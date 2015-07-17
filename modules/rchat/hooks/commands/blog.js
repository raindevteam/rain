"use strict";

const blog = {
  help: 'Gives link to "The RainBot Blog" (rainbot.info/userguide#blog)',
  ASAP: false,
  action(data, respond, done) {
    respond.say('rainbot-irc.blogspot.com');
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('blog', blog);
};
