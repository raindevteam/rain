"use strict";

var resultNest = {};

module.exports = {
  fillNests(text) {
    const nests = text.match(/\{(.*?)\}/g);
    if (nests && nests.length > 0) {
      for (let nest of nests) {
        nest = nest.replace('{', '').replace('}', '');
        if (resultNest[nest]) {
          text = text.replace('{'+ nest+'}', resultNest[nest].response);
        }
      }
    }
    return text;
  },

  nest(name, respondData) {
    var thisNest = resultNest[name] = {};
    if (respondData.output) {
      thisNest.response = respondData.output;
    } else {
      thisNest.response = '';
      for (let response of respondData.responses) {
        thisNest.response += " " + response.data;
        thisNest.response = thisNest.response.trim();
      }
    }
  },

  reset() {
    resultNest = {};
  }
};
