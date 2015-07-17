/* global __config, X */
"use strict";

const hooks = require('./hookhandler/hooks');
const alias = require('./alias');

function normal() {
  return __config.nick + ",";
}

function minimal() {
  return __config.commandSymbol;
}

/**
 *
 */

function getRawCommands(text) {
  const xtext = X(text);
  if (xtext.hasAt(normal(), 0)) return xtext.after(normal()).s;
  if (xtext.hasAt(minimal(), 0)) return xtext.after(minimal()).s;
  return text;
}

module.exports = {

  isCommand(text) {
    const xtext = X(text);

    // Check for normal and shorthand denotion
    if (xtext.hasAt(normal(), 0) || xtext.hasAt(minimal(), 0)) return true;

    if (xtext.hasAt(__config.commandPrefix, 0)) {
      const matches = text.match(
        new RegExp(__config.commandPrefix + '(\\w+)', 'g'));
      if (!matches) return false;
      for (let match of matches) {
        if (match && hooks.getCommand(match.slice(1)) ||
                     alias.isAlias(match.slice(1))) { return true; }
      }
    }
    return false;
  },

  getCommands(text) {
    let raw = getRawCommands(text);
    let cmds = [];
    let lastAmp = 0;
    let amp = 0;
    while ((amp = raw.indexOf('&', lastAmp)) > -1) {
      let eAmp = raw.indexOf('\\&', lastAmp);
      let dAmp = raw.indexOf('&&', lastAmp);
      // Escaped Amp
      if (eAmp && eAmp + 1 === amp) {
        lastAmp = eAmp + 1;
        raw = raw.substring(0, eAmp) + "&" + raw.substring(eAmp+2);
      }
      // Single Amp
      else if (eAmp + 1 !== amp && dAmp !== amp) {
        const cmd = X(raw).before('&', lastAmp).s;
        if (cmd.trim()) cmds.push(cmd.trim());
        raw = raw.substring(amp + 1);
        lastAmp = amp - (cmd.length);
      }
      // Double Amp
      else if (dAmp && dAmp === amp) {
        const ampContRe = /&(?!&)/g;
        const sub = raw.substring(lastAmp);
        const cutLen = raw.indexOf(sub);
        ampContRe.exec(sub);
        lastAmp = ampContRe.lastIndex + cutLen;
      }
    }
    // If raw still exists, push as single command
    if (raw.trim()) cmds.push(raw.trim());
    return cmds;
  }
};
