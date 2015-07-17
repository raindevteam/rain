"use strict";

const rainbox = new (require('sandbox'))();

const jip = {
  help: 'Runs a javascript snippet (rainbot.info/userguide#jip)',
  action(data, respond, done) {
    rainbox.run(data.args.join(' '), function(output) {
      let text = output.result;
      text = text.replace(/^'/, '').replace(/'$/, '');
      respond.say(text);
      return done();
    });
  }
};

module.exports = function(module) {
  module.addCommand('jip', jip);
};
