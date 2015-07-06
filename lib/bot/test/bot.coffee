# Test integrity: 2

expect = require('chai').expect
globals = require './helpers/globals'

botHelper = require './helpers/botHelper'
events = require './../events'

bot = botHelper.makeTestBot()

describe 'Bot', () ->

  before (done) ->
    bot.loadModules(done)

  it 'Should load modules', ->
    expect(bot.modules.length).to.not.equal(0)
    foundModule = undefined
    for module in bot.modules
      if module.name == 'RainCore'
        foundModule = module
    expect(foundModule).to.exist

  it 'Should attach hooks for events', () ->
    for event in events
      expect(bot.listeners(event)).to.exist

  it 'Should have a config', ->
    expect(bot.config).to.exist
