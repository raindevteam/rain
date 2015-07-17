"use strict";

const fs = require('fs');

const eplist = new Map();
const file   = '/eplist.txt';

// Read eplist.txt and load data into hashmap
fs.readFile(__dirname + file, function(err, data) {
  const raw    = data.toString();
  const buffer = raw.split('\n');
  for (let line of buffer) {
    if (!line[0] || line[0] === '#') continue;
    const ep = line.split(':');
    eplist.set(ep[0], ep[1]);
  }
});

module.exports = function(module) {

  return {

  /**
   * Returns a random eplist value, or in
   * other terms, returns a random episode.
   */

  getRandEp() {
    const randIndex = Math.floor(Math.random() * eplist.size);
    let iter = eplist.keys();
    for(let i = 0; i < randIndex - 1; i++) { iter.next(); }
    const ep = iter.next().value;
    return [ep, eplist.get(ep)];
  },

  /**
   * Retrieves episode from eplist with given key.
   *
   * @param {String} key - Key to get value for.
   */

  getEp(key) {
    return eplist.get(key);
  },

  /**
   * Returns a user friendly formatted string with
   * season and episode information.
   *
   * @param {String} key - The key to format.
   */

  keyToReadable(key) {
    const spos = key.indexOf('s');
    const epos = key.indexOf('e');
    const string =
      '(Season: ' +
      key.substr(spos + 1, epos - 1) +
      ', Episode: ' +
      key.substr(epos + 1, key.length) + ')';
    return string;
  }

  }; // return
};
