# Add a few methods to the string prototype

String.prototype.hasAt = (substr, pos, i) ->
  str = this
  if i? is true
    str    = this.toLowerCase()
    substr = substr.toLowerCase()
  return str.indexOf(substr) == pos

String.prototype.has = (substr, i) ->
  str = this
  if i? is true
    str    = this.toLowerCase()
    substr = substr.toLowerCase()
  return str.indexOf(substr) > -1

String.prototype.after = (from, i) ->
  str = this
  if i? is true
    str = this.toLowerCase()
  pos = str.indexOf(from)
  pos += from.length
  return str.substr(pos)

String.prototype.before = (from, start, i) ->
  str = this
  if i? is true
    str = this.toLowerCase()
  pos = str.indexOf(from, start)
  return str.substring(0, pos)

String.prototype.lower = () ->
  return this.toLowerCase()

String.prototype.clean = () ->
  return this.replace(/^\s+|$\s+/, '')
  .replace(/\s{2,}/g, ' ')

String.prototype.nospace = () ->
  return this.replace(/ /g, '')