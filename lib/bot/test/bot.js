// Test integrity: 2
/* globals describe, before, it */
"use strict";

const expect = require('chai').expect;
const globals = require('./helpers/globals');

const botHelper = require('./helpers/botHelper');
const events = require('./../events');

const bot = botHelper.makeBot();

describe('Bot', function() {

  before(function(done) {
    bot.loadModules(done);
  });

  it('should load modules', function() {
    expect(bot.modules.length).to.not.equal(0);
    let foundModule;
    for (let module of bot.modules) {
      if (module.name === 'RainCore')
        foundModule = module;
    }
    expect(foundModule).to.exist();
  });

  it('should attach hooks for events', function() {
    for (let event of events)
      expect(bot.listeners(event)).to.exist();
  });
});
