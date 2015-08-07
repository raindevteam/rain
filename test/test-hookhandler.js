/* globals __config*/
"use strict";

const test = require('tape');
const globals = require('./helpers/globals');
const moduleHelper = require('./helpers/modulehelper');

const testModule = moduleHelper.makeModule();
const hookHandler = require('./../lib/hookhandler');
const nester = require('./../lib/hookhandler/nest');
const helpers = require('./../lib/hookhandler/helpers');
const hooks = require('./../lib/hookhandler/hooks');

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
  t.equal(nester.fillNests(';say {test}'), ';say iambroot', "Nester fills nest");
  t.end();
});

/* End tests for hook handler -> nest */

/* Tests for hook handler -> helpers */

test('hook handler -> helpers should package data', function(t) {
  const data = helpers
                .packageData('test', mockParams, testModule, ['iambroot']);
  t.equal(data.event, 'test', "data.event");
  t.equal(data.from, 'broot', "data.from");
  t.equal(data.text, 'iambroot', "data.text");
  t.equal(data.parent, testModule, "data.parent");
  t.equal(data.args[0], 'iambroot',"data.args");
  t.ok(data.bot, "data.bot");
  t.end();
});

test('hook handler -> helpers should get command names', function(t) {
  __config.commandPrefix = ';';
  const command = ';SAY i am broot';
  const commandName = helpers.getCommandName(command);
  t.equal(commandName, 'say', 'Command name should be say');
  t.end();
});

test('hook handler -> helpers should get command args', function(t) {
  const commands = [';SAY i am broot', ';say'];
  let commandArgs = helpers.getCommandArgs(commands[0]);
  t.deepEqual(commandArgs, ['i', 'am', 'broot'], 'Command arguments are parsed');
  commandArgs = helpers.getCommandArgs(commands[1]);
  t.deepEqual(commandArgs, [], 'Returns empty param list on no args');
  t.end();
});

/* End tests for hook handler -> helpers */

/* Tests for hook handler -> hooks */

test('hook handler -> hooks sets command', function(t) {
  hooks.setCommand('test', {
    name: 'test'
  });
  t.equal(hooks.getCommand('test').name, 'test', 'Gets the command \'test\'');
  t.end();
});

test('hook handler -> hooks sets trigger', function(t) {
  hooks.setTrigger({
    name: 'test',
    event: 'message'
  });
  hooks.getTrigger('test', function(trigger) {
    t.equal(trigger.name, 'test', 'Gets the trigger \'test\'');
    t.end();
  });
});

test('hook handler -> hooks gets triggers', function(t) {
  const triggers = hooks.getTriggers();
  t.ok(triggers.message[0], 'Returns a filled triggers object');
  t.end();
});

/* End tests for hook handler -> hooks */

/* Tests for hook handler */

test('hook handler should extract commands', function(t) {
  hookHandler.extractHooks(testModule);
  t.ok(hooks.getCommand('say'), 'Properly extracts hooks');
  t.end();
});

test('hook handler should run commands', function(t) {
  hookHandler.execute(';echo', mockParams, function(responder) {
    const response = responder.getResponse();
    t.equal(response.responses[0].method, 'say', "method is say");
    t.equal(response.responses[0].to, 'you', "to is you");
    t.equal(response.responses[0].data, 'from echo', "data is from echo");
    t.end();
  });
});

test('hook handler should run aliases', function(t) {
  t.end();
});
