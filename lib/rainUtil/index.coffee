fs = require('fs')

debug = false

logFiles = {}
prompt = undefined

get_time = () ->
  return new Date().toISOString()
  .replace(/T/, ' ').replace(/\..+/, '').split(' ')[1]

exports.newLog = (log, dir) ->
  logFiles[log] = dir

exports.setPrompt = (_prompt) ->
  prompt = _prompt

exports.loggingLevel = (opt) -> debug = opt

exports.log = (msg) ->
  console.log(prompt + '[' + get_time() + '][log] ' + msg) if debug

exports.warn = (msg) ->
  console.log(prompt + '[' + get_time() + '][warn] ' + msg) if debug

exports.err = (msg) ->
  console.error(prompt + '[' + get_time() + '][err] ' + msg)

exports.logf = (log, msg) ->
  if debug
    if logFiles[log] != undefined
      file = logFiles[log]
    else return
    console.log(prompt + '[' + get_time() + '][log] ' + msg)
    fs.appendFile file, prompt + '[' + get_time() + '][log] ' + msg + '\n', (err) ->
      if err then console.error err

exports.warnf = (log, msg) ->
  if debug
    if logFiles[log] != undefined
      file = logFiles[log]
    else return
    console.log(prompt + '[' + get_time() + '][warn] ' + msg) if debug
    fs.appendFile file, prompt + '[' + get_time() + '][warn] ' + msg + '\n', (err) ->
      if err then console.error err

exports.errf = (log, msg) ->
  if debug
    if logFiles[log] != undefined
      file = logFiles[log]
    else return
    console.error(prompt + '[' + get_time() + '][err] ' + msg)
    fs.appendFile file, prompt + '[' + get_time() + '][err] ' + msg + '\n', (err) ->
      if err then console.error err
