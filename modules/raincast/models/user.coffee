mongoose = require('mongoose')

memberSchema = mongoose.Schema
  acc: String
  nicks: [String]
  channels: [{
    tag: String
    name: String
    stats: {
      messages: [String]
      actions: [String]
      averageMessageLength: Number
      longestMessageLength: Number
      averageWordCount: Number
      longestWordCount: Number
    }
  }]

module.exports = mongoose.model 'User', memberSchema
