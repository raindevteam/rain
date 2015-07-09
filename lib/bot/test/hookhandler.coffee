# Test Integrity: 3

expect = require('chai').expect
globals = require './helpers/globals'

hookHandler = require './../hookhandler'
nester = require './../hookhandler/nest'
helpers = require './../hookhandler/helpers'
parser = require './../parser'
hooks = require './../hookhandler/hooks'

moduleHelper = require './helpers/moduleHelper'
module = moduleHelper.makeTestModule()
moduleHelper.addValidCommands(module)

params = from: 'me', to: 'test', msg: {}

describe 'HookHandler', () ->

  before ->
    hookHandler.extractHooks(module)

  describe 'HookHandler -> nest', () ->

    it 'Should add and fill nests', () ->
      nester.nest 'test',
        output: undefined,
        responses: [
            { data: 'test ' }
            { data: 'another test' }
        ]
      expect(nester.fillNests('say! {test}')).to.equal('say! test another test')

  # End tests for  HookHandler -> nest

  describe 'HookHandler -> helpers', () ->

    it 'Should package data', ->
      data = helpers.packageData('test',
      from: 'mocha', text: 'testing', module, ['test'])
      expect(data.event).to.equal('test')
      expect(data.from).to.equal('mocha')
      expect(data.text).to.equal('testing')
      expect(data.parent).to.equal(module)
      expect(data.args[0]).to.equal('test')

    it 'Should get command names (front bang)', ->
      __config.preBang = false
      command = 'say! testing getCommandNames'
      commandName = helpers.getCommandName(command)
      expect(commandName).to.equal('say')

    it 'Should get command names (end bang)', ->
      __config.preBang = true
      command = '!SAY testing getCommandNames'
      commandName = helpers.getCommandName(command)
      expect(commandName).to.equal('say')

    it 'Should get command args', ->
      command = '!SAY testing getCommandArgs'
      commandArgs = helpers.getCommandArgs(command)
      expect(commandArgs[0]).to.equal('testing')
      expect(commandArgs[1]).to.equal('getCommandArgs')

  # End tests for HookHandler -> helpers

  it 'Should extract commands', ->
    hookHandler.extractHooks(module)
    expect(hooks.getCommand('say')).to.exist

  it 'Should run commands (Test #1)', (done) ->
    params.text = '!say hello'
    commands = parser.getCommands('!say hello')
    hookHandler.execute commands[0], params, (responder) ->
      response = responder.getResponse()
      expect(response.responses[0].method).to.equal('say')
      expect(response.responses[0].to).to.equal('test')
      expect(response.responses[0].data).to.equal('hello')
      return done()

  it 'Should run commands (Test #2)', (done)->
    params.text = '!testcmd2'
    commands = parser.getCommands('!testcmd2')
    hookHandler.execute commands[0], params, (responder)->
      response = responder.getResponse()
      expect(response.responses[1]).to.not.exist
      return done()
