hashmap = require('hashmap')
fs      = require('fs')
data    = undefined
buffer  = undefined

eplist = new hashmap
file   = '/eplist.txt'

# Read eplist.txt and load data into hashmap
fs.readFile __dirname + file, (err, data) ->
  raw    = data.toString()
  buffer = raw.split('\n')
  for line in buffer
    continue if line[0] == '#'
    ep = line.split(':')
    eplist.set ep[0], ep[1]

# (exports) get_rand_ep ()
#
#  Returns a random hashmap value, or in
#  other terms, returns a random episode.

exports.get_rand_ep = ->
  keys = eplist.keys()
  rand_index = Math.floor(Math.random() * keys.length)
  [ keys[rand_index], eplist.get(keys[rand_index]) ]

# (exports) get_ep (String key)
# Retrieves episode from hashmap with given key

exports.get_ep = (key) ->
  eplist.get key

# (exports) key_to_readable (String key)
#
# Returns a user friendly formatted string with
# season and episode information

exports.key_to_readable = (key) ->
  s_pos = key.indexOf('s')
  e_pos = key.indexOf('e')
  string =
    '(Season: ' +
    key.substr(s_pos + 1, e_pos - 1) +
    ', Episode: ' +
    key.substr(e_pos + 1, key.length) + ')'
  return string
