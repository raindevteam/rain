hashmap = require('hashmap')
fs      = require('fs')

parser = require './parser'

aliases = new hashmap()
file = '/aliases/aliases.txt'

writeAlias = (alias, cmd) ->
  f = fs.createWriteStream(__dirname + file)
  aliases.forEach (value, key) ->
    if value
      f.write(key + ':' + value + '\n')
  f.end()

module.exports =

  isAlias: (alias) ->
    return aliases.get(alias) != undefined

  get: (alias) ->
    return aliases.get(alias)

  add: (alias, cmd) ->
    aliases.set(alias, cmd)
    writeAlias()

  delete: (alias) ->
    aliases.remove(alias)
    writeAlias()

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
