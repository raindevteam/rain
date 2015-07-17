"use strict";

const respondQueue = require('./respondqueue');

class Responder {
  constructor(name, data) {
    this.name = name;
    this.responses = [];
    this.flushedResponses = [];
    this.outputData = undefined;
    this.deftarget = undefined;
    this.drained = false;
    return this.setDefaultTarget(data);
  }

  setDefaultTarget(data) {
    if (data.to) {
      this.deftarget = data.to;
    } else if (data.channel) {
      this.deftarget = data.channel;
    } else if (data.nick) {
      this.deftarget = data.nick;
    } else if (data.newnick) {
      this.deftarget = data.newnick;
    }
    return this;
  }

  handleResponseArgs(arg1, arg2) {
    var to;
    var response;
    if (arg2 !== undefined) {
      to = arg1;
      response = arg2;
    } else {
      to = undefined;
      response = arg1;
    }
    if (to === undefined) to = this.deftarget;
    return [to, response];
  }

  say(arg1, arg2) {
    const args = this.handleResponseArgs(arg1, arg2);
    const to = args[0];
    const response = args[1];
    if (!response) return;
    console.log(response);
    console.log(to);
    this.responses.push({
      method: 'say',
      to,
      data: response
    });
    this.drained = false;
    return this;
  }

  action(arg1, arg2) {
    const args = this.handleResponseArgs(arg1, arg2);
    const to = args[0];
    const response = args[1];
    if (!response) return;
    this.responses.push({
      method: 'action',
      to,
      data: response
    });
    this.drained = false;
    return this;
  }

  now() {
    const response = this.responses.pop();
    respondQueue.push(response);
    this.flushedResponses.push(response);
    if (this.responses.length === 0) this.drained = true;
  }

  flush() {
    for (let response of this.responses) {
      this.flushedResponses.push(response);
      respondQueue.push(response);
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

  reset() {
    this.responses = [];
    this.outputData = undefined;
  }
}

module.exports = {

  assignResponder(name, data) {
    return new Responder(name, data);
  }
};
