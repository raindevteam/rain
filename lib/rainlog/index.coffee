clc = require 'cli-color'

info = false
warn = false
err = false

fileLog = false
fileLogOnExplicit = false

prompt = "Rainlog "

# pc (String)
# Returns text in designated prompt color (pc)

pc = (text) ->
  return clc.cyanBright(text)

# nc (String)
# Returns text in designated normal color (nc)

nc = (text) ->
  return clc.white(text)

# wc (String)
# Returns text in designated warn color (wc)

wc = (text) ->
  return clc.yellowBright(text)

# ec (text) ->
# Returns text in designated err color (ec)

ec = (text) ->
  return clc.redBright(text)


getArgs = (args) ->
  if args > 2
    console.error(
      pc(prompt) +
      ec('[err][rainlog] Too many arguments were passed to a log function'))
    return null
  if args.length == 1
    return ['', args[0]]
  else if args.length == 2
    return [args[0], args[1]]
  else
    return null

module.exports =

  setPrompt: (_prompt) ->
    prompt = _prompt

  setLoggingModes: (modes) ->
    for char in modes
      switch char
        when 'i'
          info = true
        when 'w'
          warn = true
        when 'e'
          err = true
        when 'f'
          fileLog = true
        when 'F'
          fileLog = false
          fileLogOnExplicit = true

  # Logging

  info: (args...) ->
    if !info then return
    [file, text] = getArgs(args)
    console.log(pc(prompt) + nc('[info][' + file + '] ' + text))

  warn: (args...) ->
    if !warn then return
    [file, text] = getArgs(args)
    console.log(pc(prompt) + wc('[warn][' + file + '] ' + text))

  err: (args...) ->
    if !err then return
    [file, text] = getArgs(args)
    console.log(pc(prompt) + ec('[err][' + file + '] ' + text))
