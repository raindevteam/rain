"use strict";

const hashmap = require('hashmap');
const fs = require('fs');
const async = require('async');
const _ = require('lodash');

const parser = require('./parser');

var aliases = new hashmap();
const file = '/aliases/aliases.txt';

/**
 * FIXME: Check for err on file write
 * TODO: Find better way to write to file
 */

function writeAlias(done) {
  var string = "";
  aliases.forEach(function(value, key) {
    if (value) string += (key + ':' + value.args.join(' ') + ':' + value.cmd + '\r\n');
  });
  fs.writeFile(__dirname + file, string, 'utf8',
    function(err, written, string) { return done(); }
  );
}

function fillParams(argv, command) {
  _.forEach(argv, function(val, key) {
    command = command.replace('{'+ key +'}', argv[key]);
  });
  return command;
}



module.exports = {

  isAlias(alias) {
    return aliases.get(alias) !== undefined;
  },

  get(name, args, callback) {
    const alias = aliases.get(name);
    const argv = {};
    console.log(alias);
    if (args.length !== alias.args.length) 
      return callback(new Error('Invalid number of params'));
    for (let i = 0; i < alias.args.length; i++) {
      argv[alias.args[i]] = args[i];
    }
    return callback(null, fillParams(argv, alias.cmd));
  },

  add(alias, val, done) {
    aliases.set(alias, val);
    writeAlias(done);
  },

  delete(alias, done) {
    aliases.remove(alias);
    writeAlias(done);
  },

  loadAliases() {
    fs.readFile(__dirname + file, function(err, data) {
      if (data) {
        const raw = data.toString();
        const buffer = raw.split('\n');
        for (let line of buffer) {
          if (line[0] === '#' || !(line.trim())) continue;
          const alias = line.split(':');
          aliases.set(alias[0], {args: alias[1].trim().split(' '), cmd: alias[2]});
        }
      }
    });
  }
};
