/* globals */
"use strict";

const test = require('tape');
const globals = require('./helpers/globals');

const hookHandler = require('./../lib/hookhandler');
const hooks = require('./../lib/hookhandler/hooks');
const helpers = require('./../lib/hookhandler/helpers');

const respond = require('./../lib/responders');

const moduleHelper = require('./helpers/modulehelper');
const testModule = moduleHelper.makeModule();
moduleHelper.addValidCommands(testModule);

const mockParams = {from: 'broot', to: 'you', text: 'iambroot', msg: {}};
var responder;

hookHandler.extractHooks(testModule);

test('respond should assign a hook responder', function(t) {
  const hook = hooks.getCommand('echo');
  const data = helpers.packageData('message', mockParams,
                                   testModule, ['i', 'am', 'broot']);
  responder = respond.assign('hook', 'test', mockParams);
  t.ok(responder, "Responder created");
  t.equal( responder.deftarget, 'you', "Has a default target");
  t.end();
});

test('responder should take say response', function(t) {
  responder.say('iambroot');
  t.notEqual(responder.responses.length, 0, "Has a response");
  const response = responder.responses[0];
  t.ok(response, "There is a response");
  t.equal(response.method, 'say', "Uses say as method");
  t.equal(response.to, 'you', "Has proper 'to' set");
  t.equal(response.data, 'iambroot', "Has 'iambroot' as response");
  t.end();
});

test('responder should take action response', function(t) {
  responder.action('kicking this test');
  t.notEqual(responder.responses.length, 0, "Has a response");
  const response = responder.responses[1];
  t.ok(response, "There is a response");
  t.equal(response.method, 'action', "Uses action as method");
  t.equal(response.to, 'you', "Has proper 'to' set");
  t.equal(response.data, 'kicking this test', "Has 'kicking this test' as response");
  t.end();
});
