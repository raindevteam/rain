expect = require('chai').expect
#HookHandler = require './../hookhandler'

moduleHelper = require './helpers/moduleHelper'
Module = moduleHelper.makeTestModule()
moduleHelper.addValidCommands(Module)
parser = require './../parser'

describe 'Parser', () ->

  before () ->
    #HookHandler.extractCommands(Module)

  it 'Should check if text is a command (normal)', ->
    expect(parser.isCommand('HailBot, !say is a command?')).to.equal(true)

  it 'Should check if text is a command (minimal)', ->
    expect(parser.isCommand('$ !say is a command?')).to.equal(true)

  it 'Should check if text is a command (experimental)', ->
    __config.preBang = true
    expect(parser.isCommand('!say')).to.equal(true)
    __config.preBang = false
    expect(parser.isCommand('say! is a command? &roll!')).to.equal(true)

  it 'Should split commands into an array (Test #1)', ->
    commands =
      parser
        .getCommands(
          'say! testing && say! testing')
    expect(commands[0]).to.equal('say! testing && say! testing')

  it 'Should split commands into an array (Test #2)', ->
    commands =
      parser
        .getCommands(
          'HailBot, say! testing \\&\\&\\&\\&&&\\&& say! commands \\& & say! test')
    expect(commands[0]).to.equal('say! testing &&&&&&&')
    expect(commands[1]).to.equal('say! commands &')
    expect(commands[2]).to.equal('say! test')

  it 'Should split commands into an array (Test #3)', ->
    commands =
      parser
        .getCommands(
          'HailBot, say! testing &&&&&&&&&& commands&say! &&')
    expect(commands[0]).to.equal('say! testing &&&&&&&&&& commands')
    expect(commands[1]).to.equal('say! &&')

  it 'Should split commands into an array (Test #4)', ->
    commands =
      parser
        .getCommands(
          'HailBot, say! testing \\&\\& commands & ' +
          'jrun! true && false & cal! 4 + 5')
    expect(commands[0]).to.equal('say! testing && commands')
    expect(commands[1]).to.equal('jrun! true && false')
    expect(commands[2]).to.equal('cal! 4 + 5')

  it 'Should split commands into an array (Test #5)', ->
    commands =
      parser
        .getCommands(
          'HailBot, say! testing \\&\\& commands & & say! test')
    expect(commands[0]).to.equal('say! testing && commands')
    expect(commands[1]).to.equal('say! test')
