"use strict";

const math   = require('mathjs');
const _      = require('lodash');
const parser = math.parser();

const cal = {
  help: 'Parses math (rainbot.info/userguide#cal)',
  ASAP: false,
  action: function(data, respond, done) {
    const eqt = data.args.join(" ");
    let res;
    try { // Try parsing
      res = parser.eval(eqt);
    } catch (e) { /* Quietly fail */ }
    // Return result if not a function
    if (!_.isFunction(res)) respond.say(res);
    return done();
  }
};

module.exports = function(Module) {
  const rainCal = new Module('RainCal', 'rcal');
  rainCal.addCommand('cal', cal);
  return rainCal;
};
