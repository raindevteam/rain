"use strict";

const Hashmap = require('hashmap');

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

  getTrigger(name) {
    Object.keys(triggerHooks).forEach(function(triggers, event) {
      for (let trigger of triggers) {
        if (trigger.name === name) return trigger;
      }
    });
  },

  getTriggers() {
    return triggerHooks;
  }
};
