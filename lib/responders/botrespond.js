"use strict";

module.exports = function(Responder) {

class BotResponder extends Responder {
  constructor(data) {
    super(data);
    return this;
  }

  say(arg1, arg2) {
    this.pushSay(arg1, arg2);
    this.pop();
    return this;
  }

  action(arg1, arg2) {
    this.pushAction(arg1, arg2);
    this.pop();
    return this;
  }
}

return BotResponder;

};
