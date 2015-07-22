"use strict";

class Hook {

  constructor(name, type, parent, object) {
    this.name = name;
    this.parent = parent;
    this.type = type;
    this.action = object.action;
    this.listening = true;
    if (type === 'trigger') {
      this.event = object.event;
      this.onCommand = object.onCommand;
      this.trigger = object.trigger;
    } else if (type === 'command') {
      this.whitelist = object.whitelist;
      this.ASAP = object.ASAP;
      this.help = object.help;
    }
  }
}

module.exports = Hook;
