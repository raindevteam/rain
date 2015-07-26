"use strict";

module.exports = function(Module, done) {
  const rainCore = new Module("RainCore", "rcore");

  rainCore.setConfig({
    disableOnHookError: false
  });

  rainCore.load(__dirname + '/hooks/commands');
  return done(rainCore);
};
