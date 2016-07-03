[![Build Status](https://travis-ci.org/wolfchase/rainbot.svg?branch=master)](https://travis-ci.org/wolfchase/rainbot)
[![GitHub version](https://badge.fury.io/gh/wolfchase%2Frainbot.svg)](https://badge.fury.io/gh/wolfchase%2Frainbot)
RainBot
=======

A multi-purpose IRC library for quickly developing useful IRC bots.

#### A Note on Development

College proved rather difficult last quarter and therefore I was not able to work on RainBot as I
had wished. However, now that I am on summer break, I have plenty of time to waste! If you are 
viewing this message on the master branch, then 0.4.0 has been pushed out! And honsetly, is very
sloppy! This was due to ill-mannered development (shame on me!). However, with 0.5.0 development,
a new repo for documentation has been created and features the **Development Tracks**, an effort
to go from chaos to organized chaos when it comes to developing RainBot. I have no idea how it will
work, but it is obvious what I am doing right now is definitely not working! The Development Track
System will be matured along the way as development pushes on.

And yes, this README is pretty lame, but hopefully it will be much more cooler when 0.5.0 comes around!

----------------------------------

Overview
--------

RainBot was essentially just another IRC bot. However, the goal now is to give a user an API, or
rather a framework, that they can leverage to quickly develop IRC bots, simple or complex. This
package already includes two predefined installs of the bot, the default install shows how one
can build a simple but very useful IRC bot that leverages RPC to enable pluggable modules. The
second install is more of an example, and shows almost the bare mininmum of code to get an IRC
bot up and running. With future development, we hope to provide more installs to garner to the
needs that users may have. 

Installation
------------
First make sure you have properly installed go on your system (paths set and everything)

Then simply find the nearest terminal or command prompt and run:

    $ go get github.com/wolfchase/rainbot

Using the Premade Installations
-------------

If you're content enough with using one of the premade installation packages, you can 
leverage this packages commandline utilities to install. After successfully installing 
via go get, you may run:

    $ rainbot mk <install> <bot-name>

This will build a new bot from the specified install package to your GOPATH/bin.

RainBot Module Libraries (RMLS)
-------------------------------

In an effort to abstract the details of RPC when building your modules, we are currently supporting
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

**RainBot received a huge reimplemenation for its cli capabilities and this section is no longer
correct, sorry about that! Since 0.4.0 was also a bit of a slogged mess, we hope to reimpelment
the following functionality in 0.5.0. We hope 0.5.0 will be much smoother of an update!**

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
