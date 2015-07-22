"use strict";

module.exports = function(Module, done) {
  const rainBrella = new Module('RainBrella', 'rbrella');
  rainBrella.load(__dirname + '/hooks/commands');
  return done(rainBrella);
};
