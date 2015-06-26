hashmap = require('hashmap')
fs      = require('fs')
data    = undefined
buffer  = undefined

eplist = new hashmap
file   = '/eplist.txt'

# == readFile(callback)
#
#  Reads a file with episode identifiers and
#  titles, then loads them into the hashmap eplist.
#
#  @Param callback
#  Function to call when finished
#  buffering lines in file.
#
#  @Return callback()
#  Fires callback
#
fs.readFile __dirname + file, (err, data) ->
  raw    = data.toString()
  buffer = raw.split('\n')
  for line in buffer
    continue if line[0] == '#'
    ep = line.split(':')
    eplist.set ep[0], ep[1]

# == get_rand_ep()
#
#  Returns a random hashmap value, or in
#  other terms, returns a random episode.
#
exports.get_rand_ep = ->
  keys = eplist.keys()
  rand_index = Math.floor(Math.random() * keys.length)
  [ keys[rand_index], eplist.get(keys[rand_index]) ]

# get_ep(key)
#   Returns the associated value of key
# Params:
#   key - The key of the episode in the hashmap
# Return:
#   The value of the key in hashmap
exports.get_ep = (key) ->
  eplist.get key

# key_to_readable()
#   Returns a properly formatted string from key
exports.key_to_readable = (key) ->
  s_pos = key.indexOf('s')
  e_pos = key.indexOf('e')
  string =
    '(Season: ' +
    key.substr(s_pos + 1, e_pos - 1) +
    ', Episode: ' +
    key.substr(e_pos + 1, key.length) + ')'
  return string
