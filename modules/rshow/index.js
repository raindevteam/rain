"use strict";

module.exports = function(Module) {
  const rainShow = new Module('RainShow', 'rshow');
  rainShow.load(__dirname + '/hooks/commands');
  return rainShow;
};
