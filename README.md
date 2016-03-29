[![Build Status](https://travis-ci.org/wolfchase/rainbot.svg?branch=develop)](https://travis-ci.org/wolfchase/rainbot)

RainBot
=======

A multi-purpose and extensible IRC bot aimed at providing you as much freedom as possible.

Overview
--------

RainBot is built off the [Nimbus IRC library](https://github.com/RyanPrintup/Nimbus) leveraging its
event driven API to handle IRC. Nimbus is very simple in that it only provides event listening functionality,
RainBot is meant to add on to that by internally handling information such as joined channels and
users. On top of that, it also provides an RPC API that allows you to write plugable modules. Since
RPC is used, you can technically write your own module in any language provided you properly handle
the Bot's RPC protocol (Specs coming in v0.5.0).

Installation
------------

First make sure you have properly installed go on your system (paths set and everything)

Then simply find the nearest terminal or command prompt and run:

    go get github.com/wolfchase/rainbot

Configuration
-------------

// Todo

RainBot Module Libraries (RMLS)
-------------------------------

In an effort to abstract the details of RPC when building your modules, we are current supporting
three RainBot Module Libraries called RMLS.

- [**gorml**](https://github.com/wolfchase/gorml) for Golang
- [**pyrml**](https://github.com/wolfchase/pyrml) for Python and,
- [**jsrml**](https://github.com/wolfchase/jsrml) for Javascript

Most of our efforts are being put into gorml currently, so you might find the other libraries lacking for now.

Current features
----------------

#### Common IRC Events handling

RainBot can handle the management of channels and users. As early development pushes on, we hope to
handle a wider range of common IRC events as need be. Do keep in mind however that you may add listeners
via modules to any event that you may want to listen to as long as that event is handled by Nimbus ([here's
a list](https://github.com/RyanPrintup/Nimbus/blob/master/commons.go)).

#### Modules

As mentioned, you can use the bot's RPC API to write your own modules using a RML. See the RMLS section
for a list of supported RML libraries. RainBot also has some scaffolding generation to let you quickly
create biolerplate code for your module. If you have RainBot installed, you can create a scaffold by using
the following command in a terminal or command prompt

    rainbot -m <rml-prefix> <module-name>

Replace the rml-prefix with your rml of choice, so if you wish to create a Golang module scaffolding,
running

    rainbot -m go <module-name>

Will do the trick.

RainBot also has a builtin internal command that that lets you reload modules. If you have the bot
connected to IRC, and your command prefix is set to ".", you can run:

    .m reload <module-name>

which will reload a module given its module name (it will recompile if need be).

#### CLI

We also have a CLI client for the bot that let's you run IRC commands in a terminal or command prompt.
Simply start the bot with the -i flag:

    rainbot -i <config>

 Do note that this feature is not tested in version 0.4.0 and is to be considered experimental.
