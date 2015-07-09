hashmap = require('hashmap')
fs      = require('fs')
async   = require('async')

parser = require './parser'

aliases = new hashmap()
file = '/aliases/aliases.txt'

writeAlias = (done) ->
  string = ""
  aliases.forEach (value, key) ->
    if value then string += (key + ':' + value + '\r\n')
  fs.writeFile __dirname + file, string, 'utf8', (err, written, string) ->
    return done()

module.exports =

  isAlias: (alias) ->
    return aliases.get(alias) != undefined

  get: (alias) ->
    return aliases.get(alias)

  add: (alias, cmd, done) ->
    aliases.set(alias, cmd)
    writeAlias(done)

  delete: (alias, done) ->
    aliases.remove(alias)
    writeAlias(done)

  isAliasCmd: (text) ->
    if __config.preBang
      if text.match(/^!/)
        return text.match(/^!(\S+)/)[1].toLowerCase() == 'alias'
    else
      if X(text).has('!')
        return X(text).before('!').lower().s.trim() == 'alias'
      else false

  loadAliases: () ->
    fs.readFile __dirname + file, (err, data) ->
      if data
        raw    = data.toString()
        buffer = raw.split('\n')
        for line in buffer
          continue if line[0] == '#'
          alias = []
          alias[0] = X(line).before(':').s
          alias[1] = X(line).after(':').s
          aliases.set alias[0], alias[1]
