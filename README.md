[![Build Status](https://travis-ci.org/raindevteam/rain.svg?branch=master)](https://travis-ci.org/raindevteam/rain)
[![Build status](https://ci.appveyor.com/api/projects/status/7rpa57pck1m3i2a1?svg=true)](https://ci.appveyor.com/project/wolfchase/rain)
[![Go Report Card](https://goreportcard.com/badge/github.com/raindevteam/rain)](https://goreportcard.com/report/github.com/raindevteam/rain)
[![codebeat badge](https://codebeat.co/badges/1e03860e-db6d-4751-81d1-1ed3e414537e)](https://codebeat.co/projects/github-com-raindevteam-rain)
[![GoDoc](https://godoc.org/github.com/raindevteam/rain?status.svg)](https://godoc.org/github.com/raindevteam/rain)
[![GitHub license](https://img.shields.io/badge/license-AGPL-blue.svg)](https://raw.githubusercontent.com/raindevteam/rain/master/LICENSE)


----------------------------------------------------------------------------------------------------
<img align="right" src="https://avatars1.githubusercontent.com/u/20262521?v=3&s=200">
Rain, an IRC Bot Library for Go
===============================

### Overview

1. [Installation](#installation)
2. [Using the Premade Installations](#premade)
3. [Rain Module Libraries (RMLs)](#rmls)
4. [Current Features](#features)
5. [Setups](#setups)

Rain(Bot) was essentially just another IRC bot. However, the goal now is to give a user an API, or
rather a framework, that they can leverage to quickly develop IRC bots, simple or complex. This
package already includes three predefined installs of the bot, the default install shows how one
can build a simple but very useful IRC bot that leverages RPC to enable pluggable modules. The
second install is more of an example, and shows almost the bare mininmum of code to get an IRC
bot up and running. The third install is the build off the unofficial #Neverfree bot on the 
irc.canternet.org server! With future development, we hope to provide more installs to garner to the
needs that users may have. 

Documentation is still pretty lacking, as I am currently developing the bot in a pre-alpha manner (despite the 
versioning saying otherwise). Not all testing is done, and many features are still lacking. It is not recommended
to use the library just yet, unless you understand the consequences (the bot is deemed unstable and many bugs are yet
to be ironed out). If you wish to still use the library, you may view the godoc reference (at least what there is of it)
to find your way around. We hope to change of this on our way to a stable 1.0.0 release!

Installation<a name="installation"></a>
---------------------------------------
First make sure you have properly installed go on your system (paths set and everything)

##### Installing the Rain cli
Find the nearest terminal or command prompt and run:
    
    $ go get github.com/raindevteam/rain

##### Installing the subpackages
If you plan on using the subpackages, you may install their dependencies with the cli:

    $ rain depends

If you don't give a care for the cli (I think it's useful! Don't knock it til' you try it):

    $ go get -d github.com/raindevteam/rain/...

Will suffice

Using the Premade Installations<a name="premade"></a>
-----------------------------------------------------

If you're content enough with using one of the premade installation packages, you can 
leverage Rain's commandline utilities to install. After successfully installing 
via go get, you may run:

    $ rain install <install> <bot-name>

This will build a new bot from the specified install package to your GOPATH/bin. It will also
install dependencies of any of the sub packages for Rain. If you wish to install the premade
installations then make sure you install the libraries you will use from Rain before hand, as to
resolve dependencies.

Rain Module Libraries (RMLs)<a name="rmls"></a>
--------------------------------------------------

In an effort to abstract the details of RPC when building your modules, we are currently supporting
three RainBot Module Libraries called RMLs.

- [**gorml**](https://github.com/wolfchase/gorml) for Golang
- [**pyrml**](https://github.com/wolfchase/pyrml) for Python and,
- [**jsrml**](https://github.com/wolfchase/jsrml) for JavaScript

Most of our efforts are being put into gorml currently, so you might find the other libraries lacking for now.

Current features<a name="features"></a>
---------------------------------------

#### Prepackaging of Listeners via Setups

__You can view all available setups in the [setup](#setups) section__

The setup library for Rain provides predefined listeners you can attach to your bot. This allows
you to quickly setup listeners that you may need but do not want to bother having to go through
the sloggish process of writing them! We provide listeners for common tasks such as keeping track
of IRC channels and users.

#### As Close To the Metal as We Can Get You

Rain exports as many functions as it can to provide you as much customization as possible. As a 
matter of fact, one can recreate all the premade installs and setups by hand if you wish, giving you
access to the lower abstractions.

#### Modules

As mentioned, you can use the bot's RPC API to write your own modules using a RML. See the RMLS section
for a list of supported RML libraries. Rain also has some scaffolding generation to let you quickly
create biolerplate code for your module. If you have Rain installed, you can create a scaffold by using
the following command in a terminal or command prompt:

    rain m <rml-prefix> <module-name>

Replace the rml-prefix with your rml of choice, so if you wish to create a Golang module scaffolding,
running:

    rain m go MyGolangModule

Will do the trick.

Setups<a name="setups"></a>
--------------------------

Setups provide prepackaged listeners and commands to help in bot development of common IRC
tasks. These are bound to change rapidly within the next few version updates, we also hope
to add more soon.

#### Default

| Commands | Description                                                      |
|----------|------------------------------------------------------------------|
| m        | Provides module management tools, such as reload, info, and more |

| Listeners                       | Event             |
|---------------------------------|-------------------|
| Add/Update Topic                | RPL_TOPIC         |
| Update Users on Channel Connect | RPL_NAMREPLY      |
| Update User/Self on Join        | JOIN              |
| Update User/Self on Kick        | KICK              |
| Update User/Self on Kill        | KILL              |
| Update User/Self on Part        | PART              |
| Update User/Self on Quit        | QUIT              |
| Update User/Self on Nick        | NICK              |
| Update bot nick when in use     | ERR_NICKNAMEINUSE |

