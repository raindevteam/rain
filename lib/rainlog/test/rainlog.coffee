expect = require('chai').expect
rainlog = require './../index'

describe 'Rainlog', ->

  it 'Should set modes', ->
    rainlog.info 'This shouldn\'t log'
    rainlog.warn 'Nor this'
    rainlog.err 'Neither should this'

    rainlog.setLoggingModes 'i'
    rainlog.info 'tinfo', 'This is an info test from rainlog'

    rainlog.setLoggingModes 'w'
    rainlog.warn 'twarn', 'This is a warn test from rainlog'

    rainlog.setLoggingModes 'e'
    rainlog.err 'terr', 'This is a err test from rainlog'
