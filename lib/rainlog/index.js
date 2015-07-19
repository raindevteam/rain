"use strict";

const clc = require('cli-color');

var debug = false;
var info = false;
var warn = false;
var err = false;

var fileLog = false;
var fileLogOnExplicit = false;

var prompt = "Rainlog ";

// pc (String)
// Returns text in designated prompt color (pc)

function pc(text) {
  return clc.cyanBright(text);
}

// nc (String)
// Returns text in designated normal color (nc)

function nc(text) {
  return clc.white(text);
}

// wc (String)
// Returns text in designated warn color (wc)

function wc(text) {
  return clc.yellowBright(text);
}

// ec (text) ->
// Returns text in designated err color (ec)

function ec(text) {
  return clc.redBright(text);
}

function getArgs(arg1, arg2) {
  var file, text;
  if (arg2 !== undefined) {
    file = arg1;
    text = arg2;
  } else if (arg1 !== undefined) {
    file = '';
    text = arg1;
  }
  return [file, text];
}

module.exports = {

  setPrompt(_prompt) {
    prompt = _prompt;
  },

  setLoggingModes(modes) {
    for (let char of modes) {
      switch (char) {
        case 'd':
          debug = true;
          break;
        case 'i':
          info = true;
          break;
        case 'w':
          warn = true;
          break;
        case 'e':
          err = true;
          break;
        case 'f':
          fileLog = true;
          break;
        case 'F':
          fileLog = false;
          fileLogOnExplicit = true;
          break;
      }
    }
  },

  debug(arg1, arg2) {
    if (!debug) return;
    const args = getArgs(arg1, arg2);
    const file = args[0];
    const text = args[1];
    console.log(pc(prompt) + nc('[debug][' + file + '] ' + text));
  },

  info(arg1, arg2) {
    if (!info) return;
    const args = getArgs(arg1, arg2);
    const file = args[0];
    const text = args[1];
    console.log(pc(prompt) + nc('[info][' + file + '] ' + text));
  },

  warn(arg1, arg2) {
    if (!warn) return;
    const args = getArgs(arg1, arg2);
    const file = args[0];
    const text = args[1];
    console.log(pc(prompt) + wc('[warn][' + file + '] ' + text));
  },

  err(arg1, arg2) {
    if (!err) return;
    const args = getArgs(arg1, arg2);
    const file = args[0];
    const text = args[1];
    console.log(pc(prompt) + ec('[err][' + file + '] ' + text));
  }
};
