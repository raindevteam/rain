# Test Integrity: 1

expect = require('chai').expect

botHelper = require './helpers/botHelper'
events = require './../events'

bot = botHelper.makeTestBot()

Module = require('./../module')(bot)

testModule = undefined

describe 'Module', () ->

  before () ->
    testModule = new Module('test', 't')

  it 'Should have a name', ->
    expect(testModule.name).to.equal('test')

  it 'Should have an abbreviation', ->
    expect(testModule.abbrev).to.equal('t')

  it 'Should add good Commands', ->
    testModule.addCommands __dirname + '/helpers/commands'

    expect(testModule.commands.get('testcmd4')).to.exist

    testModule.addCommand 'testCmd1',
      ASAP: false,
      action: (data, respond, done) ->
        respond.say 'test command'
        return done()

    expect(testModule.commands.get('testCmd1')).to.exist

    testModule.addCommands
      testCmd2:
        ASAP: false
        action: (data, respond, done) ->
          respond.say 'test cmd2' .now()
          return done()
      testCmd3:
        ASAP: true
        action: (data, respond, done) ->
          respond.say 'test cmd3'
          return done()

    expect(testModule.commands.get('testCmd2')).to.exist
    expect(testModule.commands.get('testCmd3')).to.exist

  it 'Shouldn\'t add bad Commands', ->
    try
      testModule.addCommand 'badCmd1',
        ASAP: true
    catch e
      expect(e.message).to.equal('Invalid command')

    try
      testModule.addCommands
        testCmd2: undefined
        testCmd3:
          ASAP: true
          actio: (data, respond, done) ->
            respond.say 'test cmd3'
            return done()
    catch e
      expect(e.message).to.equal('Invalid command')


  it 'Should add valid triggers', ->
    testModule.addTrigger 'testTrigger',
      event: 'message'
      trigger: (data) -> return true
      action: (data, respond, done) ->
        respond.say('hello from testTrigger')
        return done()

    expect(testModule.triggers['message'][0]).to.exist

    testModule.addTriggers
      tesTrigger2:
        event: 'notice'
        trigger: (data) -> return true
        action: (data, respond, done) ->
          respond.say 'test trigger 2'
          return done()
      testTrigger3:
        event: 'notice'
        trigger: (data) -> return true
        action: (data, respond, done) ->
          respond.say 'test trigger 3'
          return done()

    expect(testModule.triggers['notice'][0]).to.exist
    expect(testModule.triggers['notice'][1]).to.exist

  it 'Shouldn\'t add invalid triggers', ->
    try
      testModule.addTrigger 'badtrigger1',
        event: 'message'
        trigger: true
        action: (data, respond, done) ->
          respond.say 'nope'
          return done()
    catch e
      expect(e.message).to.equal('Invalid trigger')

    try
      testModule.addTriggers
        badtrigger2:
          event: 'message'
          trigger: (data) -> true
        badtrigger3:
          trigger: (data) -> return true
          action: (data, respond, done) ->
            respond.say 'bad trigger 3'
            return done()
    catch e
      expect(e.message).to.equal('Invalid trigger')
