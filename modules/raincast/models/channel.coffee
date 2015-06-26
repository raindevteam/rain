mongoose = require('mongoose')

channelSchema = mongoose.Schema
  tag: String
  name: String
  usersSeen: [String]
  currentUsers: [String]
  totalUsers: Number
  usersAsOfNow: Number
  highestUserCount: Number
  messages: Number
  actions: Number

module.exports = mongoose.model 'Channel', channelSchema
