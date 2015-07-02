expect = require('chai').expect

Hook = require './../hook'

testtrigger =
  event: 'message'
  fireOnCommand: true
  trigger: (data) -> return true
  action: (data, respond, done) ->
    return done()

testcommand =
  ASAP: false
  action: (data, respond, done) ->
    return done()

hookTrigger = undefined
hookCommand = undefined

describe 'Hook', ->

  before () ->
    hookTrigger = new Hook('trigger', 'trigger', {}, testtrigger)
    hookCommand = new Hook('command', 'command', {}, testcommand)

  it 'Should have a name', ->
    expect(hookTrigger.name).to.exist
    expect(hookCommand.name).to.exist

  it 'Should have a parent', ->
    expect(hookTrigger.parent).to.exist
    expect(hookCommand.parent).to.exist

  it 'Should have a type', ->
    expect(hookTrigger.type).to.exist
    expect(hookCommand.type).to.exist

  it 'Should have an action', ->
    expect(hookTrigger.action).to.exist
    expect(hookCommand.action).to.exist

  it 'Should have the property "listening"', ->
    expect(hookTrigger.listening).to.exist
    expect(hookCommand.listening).to.exist

  describe 'Trigger', ->

    it 'Should have an event', ->
      expect(hookTrigger.event).to.exist

    it 'Should have a fireOnCommand property', ->
      expect(hookTrigger.fireOnCommand).to.exist

    it 'Should have a trigger', ->
      expect(hookTrigger.trigger).to.exist

  describe 'Command', ->

    it 'Should have an ASAP property', ->
      expect(hookCommand.ASAP).to.exist
