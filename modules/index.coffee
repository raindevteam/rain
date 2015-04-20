fs    = require('fs')
path  = require('path')
core  = require(__core)

exports.load = (bot, callback) ->
  fs.readdir __dirname, (err, modules) ->
    for module in modules
      if path.extname(module)
        continue # Don't load stray files
      module = require(__dirname + '/' + module)(core.module)
      core.addModule module
    return callback()
