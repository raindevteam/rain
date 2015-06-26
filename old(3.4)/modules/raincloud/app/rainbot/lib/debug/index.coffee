fs = require('fs')

debug = false

get_time = () ->
  return new Date().toISOString()
  .replace(/T/, ' ').replace(/\..+/, '') 

exports.set = (opt) -> debug = opt
  
exports.log = (msg) ->
  console.log('[' + get_time() + '] log: ' + msg) if debug
  
exports.warn = (msg) ->
  console.log('[' + get_time() + '] warn: ' + msg) if debug

exports.err = (msg) ->
  console.log('[' + get_time() + '] err: ' + msg) if debug

exports.logf = (file, msg) ->
  if debug
    fs.appendFile file, '[' + get_time() + '] log: ' + msg + '\n', (err) ->
      if err then console.log err
  
exports.warnf = (file, msg) ->
  if debug
    fs.appendFile file, '[' + get_time() + '] warn: ' + msg + '\n', (err) ->
      if err then console.log err
  
exports.errf = (file, msg) ->
  if debug
    fs.appendFile file, '[' + get_time() + '] err: ' + msg + '\n', (err) ->
      if err then console.log err