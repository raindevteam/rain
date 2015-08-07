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
  aliases.forEach(function(value, key) {
    fs.writeFile(__dirname + '/aliases/' + key + '.alias'
               , X(key + " " + value.params.join(" ") + " -> " + value.alias).clean().s
               , 'utf8'
               , function(err, written, string) { 
      return done(); 
    });
  });
}

function unlinkAlias(name, done) {
  fs.unlink(__dirname + '/aliases/' + name + '.alias', function(err) {
    return done();
  });
}

function fillParams(argv, command) {
  _.forEach(argv, function(val, key) {
    command = command.replace('{'+ key +'}', argv[key]);
  });
  return command;
}

function parseAlias(text, callback) {
  if (!text) return callback(new Error('Nothing to parse'), null);
  let aliasDeclaration = X(text).before('->').clean().s.trim();
  let aliasDefinition = X(text).after('->').clean().s.trim();
  if (!aliasDeclaration) {
    return callback(new Error('No alias declaration'), null);
  } if (!aliasDefinition) {
    return callback(new Error('No alias definition'), null);
  }
  aliasDeclaration = aliasDeclaration.split(' ');
  const aliasName = aliasDeclaration[0];
  const params = _.drop(aliasDeclaration);
  return callback(null, { name: aliasName, params, alias: aliasDefinition });
}

module.exports = {

  isAlias(alias) {
    return aliases.get(alias) !== undefined;
  },

  get(name, args, callback) {
    const alias = aliases.get(name);
    const argv = {};
    if (args.length !== alias.params.length) 
      return callback(new Error('Invalid number of params'));
    for (let i = 0; i < alias.params.length; i++) {
      argv[alias.params[i]] = args[i];
    }
    return callback(null, fillParams(argv, alias.alias));
  },

  add(alias, val, done) {
    aliases.set(alias, val);
    writeAlias(done);
  },

  delete(alias, done) {
    aliases.remove(alias);
    unlinkAlias(alias, done);
  },

  loadAliases() {
    fs.readdir(__dirname + '/aliases', function(err, files) {
      files.forEach(function(file) {
        let aliasData = fs.readFileSync(__dirname + '/aliases/' + file, 'utf8');
        parseAlias(aliasData, function(err, alias) {
          if (err) {
            rainlog.err('err', err.message);
            return;
          }
          aliases.set( alias.name
                     , { params: alias.params, alias: alias.alias }
                     );
        });
      });
    });
  },
  
  parseAlias
};
