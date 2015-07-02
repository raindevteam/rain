botHelper = require './botHelper'
Module = require('./../../module')(botHelper.makeTestBot())

bot = botHelper.makeTestBot()

module.exports =

    makeTestModule: () ->
      return new Module('test')

    addValidCommands: (module) ->
      module.addCommand('say',
        ASAP: false
        action: (data, respond, done) ->
          respond.say data.args.join ' '
          return done()
      )

      module.addCommands
        'testcmd2':
          ASAP: false
          action: (data, respond, done) ->
            respond.say('test command 2').now()
            return done()
        'testcmd3':
          ASAP: false
          action: (data, respond, done) ->
            respond.say 'test command 3'
            return done()
