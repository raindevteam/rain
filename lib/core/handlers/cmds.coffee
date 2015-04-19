core    = require(__core)
async   = require('async')
defined = core.defined
_       = require('lodash')

class CmdHandler
  constructor: (@alias) ->
    
  isCommand: (text) ->
    return text.trim().lower().has(defined.MSG_TRIGGER)

  splitCommands: (text) ->
    if @alias.isAliasCmd(text) then return [text]
    cmds     = []
    lastpipe = 0
    pipe     = 0
    lastnest = 0
    # Magic, caution hot, do not touch!
    while (pipe = text.indexOf('|', lastpipe)) > -1
      esc   = text.indexOf('\\|', lastpipe)
      dpipe = text.indexOf('||', lastpipe)
      if esc + 1 != pipe and dpipe != pipe
        cmd = text.before('|', lastpipe).trim()
        cmds.push cmd
        text = text.substring(pipe + 1)
        lastpipe = pipe+1 - (cmd.length)
      else if dpipe and dpipe == pipe
        pipeContRe = /\|(?!\|)/g
        sub = text.substring(lastpipe)
        cutLen = text.indexOf(sub)
        pipeContRe.exec(sub)
        lastpipe = pipeContRe.lastIndex + cutLen
      else if esc and esc+1 == pipe
        lastpipe = esc + 1
        text = text.substring(0, esc) + "|" + text.substring(esc+2)
    cmds.push(text)
    return cmds

  fireAlias: (cmd, args, callback) ->
    aliasCmds = @splitCommands(@alias.getAlias(cmd))
    @executeCmds aliasCmds, args, (action) ->
      callback(action)

  fireModules: (cmd, args, callback) ->
    async.eachSeries core.modules, ((module, next) ->
      module.fire 'message', [args[0], args[1], cmd, args[3]], 
      (action) ->
        if !(_.isEmpty action)
          return callback(action)
        next()
    ), (err) ->
      callback({})

  executeCmd: (cmd, args, callback) ->
    if !core.WHITELISTED(cmd.before('!'), args[0]) 
      return callback({})
    if (@alias.isAlias(cmd))
      @fireAlias cmd, args, (lastAction) ->
        lastAction.nest = cmd
        lastAction.wasAlias = true
        return callback(lastAction)
    else
      @fireModules cmd, args, (actionFired) ->
        return callback(actionFired)

  executeCmds: (cmds, args, callback) ->
    cmdsProcessed = 0
    results       = {}
    self          = this
    async.eachSeries cmds, ((cmd, next) ->
      cmdsProcessed++
      # Nest class?
      nests = cmd.match(/\{(.*?)\}/g)
      if nests?.length > 0
        for nest in nests   
          nest = nest.replace('{', '').replace('}', '')
          if results[nest]
            cmd = cmd.replace('{'+nest+'}', results[nest].res)

      self.executeCmd(cmd.trim(), args, (action) ->
        if action.responses or action.output
          if action.nest
            results[action.nest] = {}
            if action.output
              results[action.nest].res = action.output
            else
              results[action.nest].res = ''
              for response in action.responses
                results[action.nest].res +=
                if results[action.nest]
                then response.res
                else " " + response.res
          if action.ASAP and !action.wasAlias
            core.ACTION_RESPOND(action)
          if cmdsProcessed == cmds.length
            return callback(action)
        next()
      )
    ), (err) ->
      callback({})

  handle: (channel, to, text, msg) ->
    @executeCmds @splitCommands(text), [channel, to, text, msg],
    (lastAction) ->
      if lastAction?.ASAP != true
        core.ACTION_RESPOND(lastAction)
        
module.exports = CmdHandler