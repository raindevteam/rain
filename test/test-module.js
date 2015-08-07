/* globals __config*/
"use strict";

const test = require('tape');
const globals = require('./helpers/globals');
const moduleHelper = require('./helpers/modulehelper');

const Module = require('./../lib/module');
const Hook = require('./../lib/module/hook');
//const mpi = require('./../lib/module/mpi');

const testModule = moduleHelper.makeModule();

const testTrigger = {
  event: 'message',
  onCommand: true,
  trigger(data) { return true; },
  action(data, respond, done) {
    return done();
  }
};

const testCommand = {
  help: 'iambroot',
  ASAP: true,
  whitelist: ['broot'],
  action(data, respond, done) {
    return done();
  }
};

const hookTrigger = new Hook('testTrigger', 'trigger', {}, testTrigger);
const hookCommand = new Hook('testCommand', 'command', {}, testCommand);

/* Tests module -> hook */

test('module -> hook should set basic properties', function(t) {
  t.plan(10);
  t.ok(hookTrigger.name, "Has trigger name");
  t.ok(hookCommand.name, "Has command name");
  t.ok(hookTrigger.parent, "Has parent (trigger)");
  t.ok(hookCommand.parent, "Has parent (command)");
  t.ok(hookTrigger.type, "Has a type (trigger)");
  t.ok(hookCommand.type, "Has a type (command)");
  t.ok(hookTrigger.action, "Has an action (trigger)");
  t.ok(hookCommand.action, "Has an action (command)");
  t.ok(hookTrigger.listening, "Has listening (trigger)");
  t.ok(hookCommand.listening, "Has listening (command)");
  t.end();
});

test("module -> hook should set trigger properties", function(t) {
  t.plan(3);
  t.ok(hookTrigger.event, "Has event");
  t.ok(hookTrigger.onCommand, "Has onCommand");
  t.ok(hookTrigger.trigger, "Has trigger");
  t.end();
});

test("module -> hook should set command properties", function(t) {
  t.plan(3);
  t.ok(hookCommand.help, "Has help text");
  t.ok(hookCommand.ASAP, "Has ASAP");
  t.ok(hookCommand.whitelist, "Has whitelist");
  t.end();
});

/* End tests for module -> hook */

/* Tests for module */

test('module should have a name', function(t) {
  t.equal(testModule.name, 'TestModule', "Has name");
  t.end();
});

test('module should have an abbreviation', function(t) {
  t.equal(testModule.abbrev, 'tm', "Has abbreviation");
  t.end();
});
