"use strict";

module.exports = function(Module) {
  const rainCore = new Module("RainCore", "rcore");
  rainCore.load(__dirname + '/hooks/commands');
  return rainCore;
};
