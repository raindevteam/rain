mongoose = require('mongoose')

userSchema = mongoose.Schema 
  accName: String
  nicks: [String]

module.exports = mongoose.model 'Users', userSchema