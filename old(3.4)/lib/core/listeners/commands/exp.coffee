core = __core
_    = require('lodash')

module.exports =
  ### == dev (command) ==
  Configures development variables and settings
  ###
  dev:
    action: (args, respond, done) ->
      if (args[0])
        switch args[0]
          when "set"
            if (args[1])
              switch args[1]
                when "parser"
                  if (args[2])
                    if (args[2] == "default")
                      respond.say "Using default parser"
                      return done()
                    if (args[2] == "minimal")
                      respond.say "Using minimal parser"
                      return done()
                    if (args[2] == "experimental")
                      respond.say "Using experimental parser"
                      return done()
                    else
                      respond.say "Unknown option: " + args[2]
                else
                  respond.say "Unknown option: " + args[1]
          else
            respond.say "Unknown option: " + args[0]
