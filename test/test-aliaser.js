/* globals */
"use strict";

const globals = require('./helpers/globals');
const test = require('tape');
const fs = require('fs');
const aliaser = require('./../lib/alias');

test('aliaser should error on no definition', function(t) {
  aliaser.parseAlias('badAlias ->', function(err, alias) {
    t.equal(err.message, 'No alias definition', 'Errors on no definition');
    t.end();
  });
});

test('aliaser should error on no declaration', function(t) {
  aliaser.parseAlias('-> ;bad alias', function(err, alias) {
    t.equal(err.message, 'No alias declaration', 'Errors on no declaration');
    t.end();
  });
});

test('aliaser should parse correctly', function(t) {
  aliaser.parseAlias('test-alias x -> ;say test', function(err, alias) {
    t.equal(alias.name, 'test-alias', 'correct alias name');
    t.deepEqual(alias.params, ['x'], 'correct params');
    t.equal(alias.alias, ';say test', 'correct alias defintion');
    t.end();
  });
});


test('aliaser should add alias', function(t) {
  aliaser.add('test-alias', { params: ['x'], alias: ';say {x}' });
  aliaser.get('test-alias', ['iambroot'], function(err, alias) {
    t.equal(alias, ';say iambroot');
    t.end();
  });
});

test('aliaser should err on invalid number of params', function(t) {
  aliaser.get('test-alias', [], function(err, alias) {
    t.equal(err.message, 'Invalid number of params', 'Errors on invalid params');
    t.end();
  });
});

test('aliaser should delete alias', function(t) {
  aliaser.delete('test-alias');
  let fileCheck = fs.existsSync(__dirname + '/../lib/aliases/test-alias.alias');
  t.ok(!fileCheck, 'Alias removed');
  t.end();
});
