"use strict";

module.exports = function(Responder) {

class HookResponder extends Responder {
  constructor(name, data) {
    super(name, data);
    this.outputData = undefined;
    return this;
  }

  say(arg1, arg2) {
    this.pushSay(arg1, arg2);
    return this;
  }

  action(arg1, arg2) {
    this.pushAction(arg1, arg2);
    return this;
  }

  paste(arg1, arg2) {
    this.pushPaste(arg1, arg2);
    return this;
  }

  now() {
    this.pop();
  }

  flush() {
    if (this.responses.length > 0) {
      for (let response of this.responses) {
        this.pop();
      }
    }
    this.drained = true;
    this.responses = [];
    return {output: this.outputData, responses: this.flushedResponses};
  }

  flushed() {
    return this.drained;
  }

  output(data) {
    this.outputData = data;
  }

  getResponse() {
    const responses = this.flushedResponses.concat(this.responses);
    return {
      output: this.outputData,
      responses: responses
    };
  }
}

return HookResponder;

};
