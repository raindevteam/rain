# Requires
fs      = require('fs')
path    = require('path')
core    = require(__core)

# Locals
modules = []

exports.load = (bot, callback) ->
  fs.readdir __dirname, (err, mods) ->
    for mod in mods
      if path.extname(mod)
        continue # Don't load stray files
      module = require(__dirname + '/' + mod)(core.module)
      modules.push module
    return callback modules