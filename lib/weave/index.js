"use strict";

// Weave is my own string library incarnation, inspired by
// stringjs found at stringjs.com

class X {
  constructor(string) {
    this.s = string;
    this.from = 0;
  }

  // X :: hasAt (String, Int, Boolean)
  // Checks if substr is at index pos of this.s

  hasAt(substr, pos, i) {
    let str = this.s;
    if (i && i === true) {
      str    = this.s.toLowerCase();
      substr = substr.toLowerCase();
    }
    return str.indexOf(substr) === pos;
  }

  // X :: has (String, Boolean)
  // Checks if substring exists anywhere in this.s

  has(substr, i) {
    let str = this.s;
    if (i && i === true) {
      str    = this.s.toLowerCase();
      substr = substr.toLowerCase();
    }
    return str.indexOf(substr) > -1;
  }

  // X :: after (String, Boolean)
  // Gets string after substring from

  after(from, i) {
    let str = this.s;
    if (i && i === true) {
      str = this.s.toLowerCase();
    }
    let pos = str.indexOf(from);
    pos += from.length;
    this.s = str.substr(pos);
    return this;
  }

  // X :: before (String, Int, Boolean)
  // Gets string before substring from starting from pos start

  before(from, start, i) {
    let str = this.s;
    if (i && i === true) {
      str = this.s.toLowerCase();
    }
    let pos = str.indexOf(from, start);
    this.s = str.substring(0, pos);
    return this;
  }

  // X :: lower ()
  // Makes string lower case

  lower() {
    this.s = this.s.toLowerCase();
    return this;
  }

  // X :: clean ()
  // Removes ALL extra whitespace

  clean() {
    this.s = this.s.replace(/^\s+|$\s+/, '').replace(/\s{2,}/g, ' ');
    return this;
  }

  // X :: nospace ()
  // Removes whitespace COMPLETELY

  nospace() {
    this.s = this.s.replace(/ /g, '');
    return this;
  }
}

module.exports = function (string) { return new X(string); };
