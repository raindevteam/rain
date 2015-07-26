/* global rainlog, __config */
"use strict";

const Hashmap = require('hashmap');
const fs = require('fs');
const _ = require('lodash');

const mpi = require('./mpi');
const Hook = require('./hook');
const events = require('./../events');

const validConfigOptions = [
  'disableOnHookError'
];

function isValidCommand(command) {
  if (!command) return false;
  if (!command.action) return false;
  if (!(_.isFunction(command.action))) return false;
  return true;
}

function isValidTrigger(trigger) {
  if (!trigger) return false;
  if (!trigger.event) return false;
  if (!trigger.trigger) return false;
  if (!(_.isFunction(trigger.trigger))) return false;
  if (!trigger.action) return false;
  if (!(_.isFunction(trigger.action))) return false;
  return true;
}

module.exports = function(bot) {

  class Module {

    /**
     *
     *
     */

    constructor(name, abbrev) {
      this.name = name;
      this.abbrev = abbrev;
      this.bot = bot;
      this.enabled = true;
      this.commands = new Hashmap();
      this.triggers = {};
      for (let event of events) this.triggers[event] = [];
    }

    setConfig(options) {
      const self = this;
      _.forEach(options, function(option, key) {
        if (validConfigOptions.indexOf(key) < -1) {
          rainlog.err(self.name, key + ' is not a valid module config option');
        } else self[key] = option;
      });
    }

    load(directory) {
      if (fs.lstatSync(directory).isDirectory()) {
        const files = fs.readdirSync(directory);
        for (let file of files) {
          require(directory + '/' + file)(this);
        }
      }
    }

    addCommand(name, command) {
      if (!isValidCommand(command)) {
        const err = new Error('Invalid command: ' + name);
        rainlog.err(this.name, err.message);
        return err;
      }
      const hook = new Hook(name, 'command', this, command);
      this.commands.set(name, hook);
   }

    addCommands(commands) {
      const self = this;
      if (_.isObject(commands)) {
        _.forEach(commands, function(command, name) {
          self.addCommand(name, command);
        });
      } else {
        if (fs.lstatSync(commands).isDirectory()) {
          const files = fs.readdirSync(commands);
          for (let file of files) {
            const command = require(commands + '/' + file)(self);
            _.forEach(command, function(val, key) {
              self.addCommand(key, val);
            });
          }
        }
      }
    }

    addTrigger(name, trigger) {
      if (!isValidTrigger(trigger)) throw new Error('Invalid trigger');
      const hook = new Hook(name, 'trigger', this, trigger);
      this.triggers[trigger.event].push(hook);
    }

    addTriggers(triggers) {
      const self = this;
      _.forEach(triggers, function(trigger, name) {
        self.addTrigger(name, trigger);
      });
    }


    enable() {
      this.enabled = true;
    }

    disable() {
      this.enabled = false;
    }

    status() {
      return this.enabled;
    }

    broke(err) {
      this.enabled = false;
      //bot.reportBroken(this);
    }
  }

  return Module;
};
