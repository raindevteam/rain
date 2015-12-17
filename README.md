[![Build Status](https://travis-ci.org/wolfchase/rainbot.svg?branch=develop)](https://travis-ci.org/wolfchase/rainbot)
# RainBot

This is the current go project for RainBot. It is a total rewrite of the original JavaScript codebase. However, some
features of the bot have dramatically changed, such as the plugin manager (which now uses rpc). We hope that in using
go, RainBot will have improved performance, a better memory footprint and better maintainability. This readme will
be updated as early development releases are pushed out.

## Current features

The bot really doesn't offer much right now, as it is still in very early development. However, the module library uses
RPC server communication via JSON codecs. In other words, if you want to program modules for the bot, you could do it
in any language that supports RPC with a JSON codec. We currently have a Go module library to get you on your feet
for writing modules. Any other language however, will need RPC set up. I already have planned a JavaScript library,
and hopefully I can pull together others in the future.

## Setting Up the Bot

// Todo
