"use strict";

const hashmap = require('hashmap');
const fs = require('fs');
const async = require('async');

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
    if (value) string += (`${key}:${value}\r\n`);
  });
  fs.writeFile(__dirname + file, string, 'utf8',
    function(err, written, string) { return done(); }
  );
}

module.exports = {

  isAlias(alias) {
    return aliases.get(alias) !== undefined;
  },

  get(alias) {
    return aliases.get(alias);
  },

  add(alias, cmd, done) {
    aliases.set(alias, cmd);
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
          if (line[0] === '#') continue;
          aliases.set(X(line).before(':').s, X(line).after(':').s);
        }
      }
    });
  }
};
