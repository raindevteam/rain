# Weave is my own string library incarnation, inspired by
# stringjs found at stringjs.com

class X
  constructor: (string) ->
    @s = string
    @from = 0

  # X :: hasAt (String, Int, Boolean)
  # Checks if substr is at index pos of @s

  hasAt: (substr, pos, i) ->
    str = @s
    if i? is true
      str    = @s.toLowerCase()
      substr = substr.toLowerCase()
    return str.indexOf(substr) == pos

  # X :: has (String, Boolean)
  # Checks if substring exists anywhere in @s

  has: (substr, i) ->
    str = @s
    if i? is true
      str    = @s.toLowerCase()
      substr = substr.toLowerCase()
    return str.indexOf(substr) > -1

  # X :: after (String, Boolean)
  # Gets string after substring from

  after: (from, i) ->
    str = @s
    if i? is true
      str = @s.toLowerCase()
    pos = str.indexOf(from)
    pos += from.length
    @s = str.substr(pos)
    return @

  # X :: before (String, Int, Boolean)
  # Gets string before substring from starting from pos start

  before: (from, start, i) ->
    str = @s
    if i? is true
      str = @s.toLowerCase()
    pos = str.indexOf(from, start)
    @s = str.substring(0, pos)
    return @

  # X :: lower ()
  # Makes string lower case

  lower: () ->
    @s = @s.toLowerCase()
    return @

  # X :: clean ()
  # Removes ALL extra whitespace

  clean: () ->
    @s = @s.replace(/^\s+|$\s+/, '')
    .replace(/\s{2,}/g, ' ')
    return @

  # X :: nospace ()
  # Removes whitespace COMPLETELY

  nospace: () ->
    @s = @s.replace(/ /g, '')
    return @

module.exports = (string) -> return new X(string)
