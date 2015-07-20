/* globals */
"use strict";

const globals = require('./helpers/globals');
const test = require('tape');
const botHelper = require('./helpers/bothelper');
const events = require('./../lib/bot/events');

const Bot = botHelper.makeBot();

test("bot should load modules", function(t) {
  Bot.loadModules(function() {
    t.equal(5, Bot.modules.length, "all modules loaded");
    t.end();
  });
});

test("bot should attach hooks", function(t) {
  t.plan(events.length);
  for (let event of events)
    t.ok(Bot.listeners(event), "Event listeners present for " + event);
  t.end();
});
