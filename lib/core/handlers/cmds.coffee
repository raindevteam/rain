async   = require('async')
defined = __core.defined
_       = require('lodash')
config  = require(__config)

class CmdHandler
  constructor: (aliasHandler, commands) ->
    @commands = commands
    @aliasHandler = aliasHandler

  isCommand: (text) ->
    return text.trim().lower().has(defined.MSG_TRIGGER)

  # Receives a raw command text line and parses it into an array
  getCommands: (text) ->
    if @aliasHandler.isAliasCmd(text) then return [text.trim()]
    cmds = []
    lastpipe = 0
    pipe = 0
    lastnest = 0
    # Magic, caution hot, do not touch!
    while (pipe = text.indexOf('|', lastpipe)) > -1
      esc   = text.indexOf('\\|', lastpipe)
      dpipe = text.indexOf('||', lastpipe)
      if esc + 1 != pipe and dpipe != pipe
        cmd = text.before('|', lastpipe).trim()
        if cmd.trim() then cmds.push(cmd.trim())
        text = text.substring(pipe + 1)
        lastpipe = pipe + 1 - (cmd.length)
      else if dpipe and dpipe == pipe
        pipeContRe = /\|(?!\|)/g
        sub = text.substring(lastpipe)
        cutLen = text.indexOf(sub)
        pipeContRe.exec(sub)
        lastpipe = pipeContRe.lastIndex + cutLen
      else if esc and esc + 1 == pipe
        lastpipe = esc + 1
        text = text.substring(0, esc) + "|" + text.substring(esc+2)
    if text.trim() then cmds.push(text.trim())
    return cmds

  getCommandName: (commandRaw) ->
    if config.bang == 'front'
      if commandRaw.match(/^!/)
        return commandRaw.match(/^!(\S+)/)[1]
    else
      if commandRaw.has('!')
        return commandRaw.before('!').trim()
      else null

  getCommandArgs: (commandRaw) ->
    if config.bang == 'front'
      if commandRaw.match(/^!/)
        return _.drop(commandRaw.split(' '))
    else
      if commandRaw.has('!')
        return commandRaw.after('!').trim().split(' ')
      else null

  executeCommand: (commandRaw, responseHandler, done) ->
    commandName = @getCommandName(commandRaw)
    commandArgs = @getCommandArgs(commandRaw)
    if !commandName then return done()
    # if !core.WHITELISTED commandName, ar then done({})
    if @aliasHandler.isAlias commandName
      aliasCommands = @getCommands(@aliasHandler.getAlias(commandName))
      @run aliasCommands, responseHandler, (lastCommand) ->
         if lastCommand
           lastCommand.name = commandName
           lastCommand.wasAlias = true
         return done lastCommand
    else if @commands.has(commandName.trim())
      command = @commands.get(commandName)
      if responseHandler.to == config.nick # is PM
        if command.pm && command.pm == false then return done {}
      else
        if command.msg && command.msg == false then return done {}
      action = command.action commandArgs, responseHandler, () ->
        command.name = commandName
        return done command

  run: (commands, responseHandler, done) ->
    commandsProcessed = 0
    results = {}
    self = this
    async.eachSeries commands, ((commandRaw, next) ->
      commandsProcessed++
      nests = commandRaw.match(/\{(.*?)\}/g)
      if nests?.length > 0
        for nest in nests
          nest = nest.replace('{', '').replace('}', '')
          if results[nest]
            commandRaw =
              commandRaw.replace('{'+nest+'}', results[nest].response)
      self.executeCommand(commandRaw, responseHandler, (firedCommand) ->
        if !firedCommand then return done()
        commandName = firedCommand.name
        responses   = responseHandler.responses
        output      = responseHandler.data
        if responses or output
          results[commandName] = {}
          if output then results[commandName].response = output
          else
            results[commandName].response = ''
            for response in responses
              results[commandName].response += " " + response.res
              results[commandName].response =
              results[commandName].response.trim()
        if commandsProcessed == commands.length
          return done()
        if firedCommand.ASAP and !firedCommand.wasAlias
          responseHandler.respond()
        responseHandler.reset()
        next()
      )
    ), (err) ->
      return done()

module.exports = CmdHandler
