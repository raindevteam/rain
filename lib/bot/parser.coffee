hooks = require './hookhandler/hooks'
alias = require './alias'

normal = __config.nick + ","
minimal = __config.commandSymbol

# getRawCommands (String)
#
# Just removes the command denotion header if any

getRawCommands = (text) ->
  if X(text).has(normal, 0)
    return X(text).after(normal).s
  if X(text).has(minimal, 0)
    return X(text).after(minimal).s
  else return text

module.exports =

  # isCommand (String)
  #
  # Checks whether the string is a command to the bot. Checks
  # first if command has long denotion, then checks minimal
  # and finally uses the experimental method last.

  isCommand: (text) ->
    if X(text).has(normal, 0) then return true
    if X(text).has(minimal, 0) then return true
    if __config.preBang
      start = X(text).before(' ').s
      if !start then start = text
      if !start.match(/^!/) then return false
      matches = text.match(/!(\w+)/g)
      for match in matches
        if hooks.getCommand(match[1..]) or alias.isAlias(match[1..])
          return true
    else
      start = X(text).before(' ').s
      if !start.match(/!$/) then return false
      if !start then start = text
      matches = text.match(/(\w+)!/g)
      for match in matches
        match = X(match).before('!').s
        if hooks.getCommand(match) or alias.isAlias(match)
          return true
    return false

  # getCommands (String)
  #
  # The magical function that splits and parses a raw text command.
  # I forgot how it worked since its been while (since RainBot 2.0),
  # but considering it hasn't broken up until now, we are going to
  # keep it as is ("If it ain't broke, don't fix it").
  #
  # Things of Interest: (Things I know it does)
  #  + Splits commands on & (amp)
  #  + It will take \& (eAmp) and replace it with &
  #  + Doesn't split on && (dAmp)

  getCommands: (text) ->
    # if @aliasHandler.isAliasCmd(text) then return [text.trim()]
    [text, cmds, lastAmp, amp] = [getRawCommands(text), [], 0, 0]
    # Magic, caution hot, do not touch!
    while (amp = text.indexOf('&', lastAmp)) > -1
      eAmp = text.indexOf('\\&', lastAmp)
      dAmp = text.indexOf('&&', lastAmp)
      if eAmp and eAmp + 1 == amp # We encountered an eAmpaped amp
        # We obviously want to eAmpape the amp
        lastAmp = eAmp + 1
        text = text.substring(0, eAmp) + "&" + text.substring(eAmp+2)
      else if eAmp + 1 != amp and dAmp != amp # We encountered a single amp
        # Grab everything before amp
        cmd = X(text).before('&', lastAmp).s
        # Push to cmds array if cmd exists
        if cmd.trim() then cmds.push(cmd.trim())
        # Remove the cmd we just pushed from the original text
        text = text.substring(amp + 1)
        # Keep track of what amp we just processed
        lastAmp = amp - (cmd.length)
      else if dAmp and dAmp == amp # We encountered a double amp
        # Here, I have no idea what the algorithm does
        # I'll have to write everything down on paper to find out again
        # However, if I'm correct, I'm pretty sure we wanted to skip All
        # double amps until we got to a regular amp.
        ampContRe = /&(?!&)/g
        sub = text.substring(lastAmp)
        cutLen = text.indexOf(sub)
        ampContRe.exec(sub)
        lastAmp = ampContRe.lastIndex + cutLen
    if text.trim() then cmds.push(text.trim())
    return cmds
