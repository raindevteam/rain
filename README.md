## A New Blog For Development Updates!
http://rainbot-irc.blogspot.com/

## This README
Is incomplete. I hope to add more to it soon such as how to get RainBot working for you and how to make your own module!

# RainBot
Multipurpose, extensible IRC bot written with node.js

## Overview
RainBot is just another IRC bot among the many already made; it started as something for fun rather than need. That being said, plenty of development has still been going into RainBot, well at least as much as a college student can do without drowning themself in work. The master branch is for versions of the bot that are somewhat stable, but not entirely, stable enough that the bot can be used to a reasonable extent. The development branch is where I do all my work and reflects the most up-to-date changes, bugs, and features.

### Features

#### Modules
The Bot has been reworked quite a few times to provide easy extensibility via modules. Writing modules for the bot is a short and sweet process, but still aims to provide as much flexibility as it can.

Here's a small example showing how you can make a module sporting an echo command!

    module.exports = (Module) ->
      # Make a new module!
      myModule = new Module('myModule')
      # What good is a module without hooks!
      myModule.addCommand 'echo',
        help: "Repeats its arguments!"
        action: (data, respond, done) ->
          respond.say data.args.join ' ' if data.args
          return done()
      # Shut up and take my module!
      return myModule


#### Command Nesting
Sometimes you just want to string commands together or have commands pass values to each other. RainBot's command nesting aims to make this a realization. You can pipe together separate commands and have them executed in one go, or pass values from one command to another! Try these examples for a little taste:

    RainBot, cal! f(x) = x * x | cal! f(3)
    RainBot, roll! | jrun! if ({roll} > 3) "What a roll!"; else "Aw what a bad roll"

#### Command Aliasing
Command nesting allows you to string together commands, however they can get pretty lengthy. Aliasing solves that by allowing you to group a set of commands under a single alias name.

    Add alias:
    RainBot, alias! longCommand roll! | jrun! if ({roll} > 3) "What a roll!"; else "Aw what a bad roll"
    Remove alias:
    Rainbot, alias! rm longCommand

Currently you can only set and remove aliases, but listing aliases will be implemented soon.


## Config
// Todo

## Modules
// Todo
