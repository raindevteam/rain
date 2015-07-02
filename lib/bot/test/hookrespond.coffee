expect = require('chai').expect
globals = require './helpers/globals'

hookHandler = require './../hookhandler'
hooks = require './../hookhandler/hooks'
helpers = require './../hookhandler/helpers'

HookRespond = require './../responders/hookrespond'

moduleHelper = require './helpers/moduleHelper'
module = moduleHelper.makeTestModule()
moduleHelper.addValidCommands(module)

hookHandler.extractCommands(module)

responder = undefined

describe 'HookRespond', () ->

  it 'Should assign a responder and set default target', ->
    params = from: 'tester', to: '#snowybottest', text: '!say is command'
    hook = hooks.getCommand('say')
    data = helpers.packageData('message', params, hook.parent, ['is', 'command'])
    responder = HookRespond.assignResponder('say').setDefaultTarget(data)
    expect(responder).to.exist
    expect(responder.deftarget).to.equal('#snowybottest')

  it 'Should take a say response', ->
    responder.say('Hello world')
    expect(responder.responses.length).to.be.above(0)
    expect(responder.responses[0].method).to.equal('say')
    expect(responder.responses[0].to).to.equal('#snowybottest')
    expect(responder.responses[0].data).to.equal('Hello world')

  it 'Should take an action response', ->
    responder.action('kicking this test')
    expect(responder.responses.length).to.be.above(0)
    expect(responder.responses[1].method).to.equal('action')
    expect(responder.responses[1].to).to.equal('#snowybottest')
    expect(responder.responses[1].data).to.equal('kicking this test')
