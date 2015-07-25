"use strict";

const Hashmap = require('hashmap');
const _ = require('lodash');

const events = require('./../events');

var commandHooks = new Hashmap();
var triggerHooks = {};

for (let event of events) {
  triggerHooks[event] = [];
}

module.exports = {

  setCommand(name, command) {
    commandHooks.set(name, command);
  },

  setTrigger(trigger) {
    triggerHooks[trigger.event].push(trigger);
  },

  getCommand(name) {
    return commandHooks.get(name);
  },

  getTrigger(name, callback) {
    _.forEach(triggerHooks, function(triggers, event) {
      for (let trigger of triggers) {
        if (trigger.name === name) return callback(trigger);
      }
    });
  },

  getTriggers() {
    return triggerHooks;
  }
};
