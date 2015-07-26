"use strict";

module.exports = function(Module, done) {
  const rainCore = new Module("RainCore", "rcore");
  rainCore.load(__dirname + '/hooks/commands');
  return done(rainCore);
};
