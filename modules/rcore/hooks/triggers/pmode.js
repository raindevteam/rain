"use strict";

const pmode = {
  event: "+mode",
  trigger(data) { return true; },
  action(data, respond, done) {
    const chans = data.bot.chans;
    console.log(data);
  }
}

module.exports = function(module) {
  module.addTrigger('rcore-pmode', pmode);
}
