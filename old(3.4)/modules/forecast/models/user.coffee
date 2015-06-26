mongoose = require('mongoose')

memberSchema = mongoose.Schema
  acc: String
  nicks: [String]
  num: Number
  channels: [{
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
