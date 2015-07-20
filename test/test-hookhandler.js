/* globals __config*/
"use strict";

const test = require('tape');
const globals = require('./helpers/globals');
const moduleHelper = require('./helpers/modulehelper');

const testModule = moduleHelper.makeModule();
const hookHandler = require('./../lib/bot/hookhandler');
const nester = require('./../lib/bot/hookhandler/nest');
const helpers = require('./../lib/bot/hookhandler/helpers');
const hooks = require('./../lib/bot/hookhandler/hooks');

moduleHelper.addValidCommands(testModule);
const mockParams = {from: 'broot', to: 'you', text: 'iambroot', msg: {}};

/* Tests for hook handler -> nest */

test("hook handler -> nest should add and fill", function(t) {

  // Add a nest
  nester.nest('test', {
    output: undefined,
    responses: [{
      method: 'say',
      to: 'test',
      data: 'iambroot'
    }]
  });

  // If the nest was added correctly, it should fill a nest properly.
  t.equal(';say iambroot', nester.fillNests(';say {test}'));
  t.end();
});

/* End tests for hook handler -> nest */

/* Tests for hook handler -> helpers */

test('hook handler -> helpers should package data', function(t) {
  const data = helpers
                .packageData('test', mockParams, testModule, ['iambroot']);
  t.equal('test', data.event, "data.event");
  t.equal('broot', data.from, "data.from");
  t.equal('iambroot', data.text, "data.text");
  t.equal(testModule, data.parent, "data.parent");
  t.equal('iambroot', data.args[0], "data.args");
  t.ok(data.bot, "data.bot");
  t.end();
});

test('hook handler -> helpers should get command names', function(t) {
  __config.commandPrefix = ';';
  const command = ';SAY i am broot';
  const commandName = helpers.getCommandName(command);
  t.equal('say', commandName);
  t.end();
});

test('hook handler -> helpers should get command args', function(t) {
  const command = ';SAY i am broot';
  const commandArgs = helpers.getCommandArgs(command);
  t.deepEqual(['i', 'am', 'broot'], commandArgs);
  t.end();
});

/* End tests for hook handler -> helpers */

/* Tests for hook handler -> hooks */

test('hook handler -> hooks sets command', function(t) {
  hooks.setCommand('test', {
    name: 'test'
  });
  t.equal('test', hooks.getCommand('test').name);
  t.end();
});

test('hook handler -> hooks sets trigger', function(t) {
  hooks.setTrigger({
    name: 'test',
    event: 'message'
  });
  hooks.getTrigger('test', function(trigger) {
    t.equal('test', trigger.name);
    t.end();
  });
});

test('hook handler -> hooks gets triggers', function(t) {
  const triggers = hooks.getTriggers();
  t.equal('test', triggers.message[0].name);
  t.end();
});

/* End tests for hook handler -> hooks */

/* Tests for hook handler */

test('hook handler should extract commands', function(t) {
  hookHandler.extractHooks(testModule);
  t.ok(hooks.getCommand('say'));
  t.end();
});

test('hook handler should run commands', function(t) {
  hookHandler.execute(';echo', mockParams, function(responder) {
    const response = responder.getResponse();
    console.log(response);
    t.equal('say', response.responses[0].method, "method is say");
    t.equal('you', response.responses[0].to, "to is you");
    t.equal('from echo', response.responses[0].data, "data is from echo");
    t.end();
  });
});

test('hook handler should run aliases', function(t) {
  t.end();
});
