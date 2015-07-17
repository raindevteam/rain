"use strict";

module.exports = function(Module) {
  const rainBrella = new Module('RainBrella', 'rbrella');
  rainBrella.load(__dirname + '/hooks/commands');
  return rainBrella;
};
