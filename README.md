[![Build Status](https://travis-ci.org/raindevteam/rain.svg?branch=develop)](https://travis-ci.org/wolfchase/rainbot)

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

1. [Installation](#installation)
2. [Using the Premade Installations](#premade)
3. [Rain Module Libraries (RMLs)](#rmls)
4. [Current Features](#features)
5. [Setups](#setups)

RainBot was essentially just another IRC bot. However, the goal now is to give a user an API, or
rather a framework, that they can leverage to quickly develop IRC bots, simple or complex. This
package already includes three predefined installs of the bot, the default install shows how one
can build a simple but very useful IRC bot that leverages RPC to enable pluggable modules. The
second install is more of an example, and shows almost the bare mininmum of code to get an IRC
bot up and running. The third install is the build off the unofficial #Neverfree bot on the 
irc.canternet.org server! With future development, we hope to provide more installs to garner to the
needs that users may have. 

Installation<a name="installation"></a>
---------------------------------------
First make sure you have properly installed go on your system (paths set and everything)

Then simply find the nearest terminal or command prompt and run:

    $ go get github.com/raindevteam/rain

Using the Premade Installations<a name="premade"></a>
-----------------------------------------------------

If you're content enough with using one of the premade installation packages, you can 
leverage Rain's commandline utilities to install. After successfully installing 
via go get, you may run:

    $ rain mk <install> <bot-name>

This will build a new bot from the specified install package to your GOPATH/bin.

Rain Module Libraries (RMLS)<a name="rmls"></a>
--------------------------------------------------

In an effort to abstract the details of RPC when building your modules, we are currently supporting
three RainBot Module Libraries called RMLS.

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
to add more starting with 0.6.0.

#### Default

| Commands | Action                                                           |
|----------|------------------------------------------------------------------|
| m        | Provides module management tools, such as reload, list, and more |

| Listeners                       | Event        |
|---------------------------------|--------------|
| Add/Update Topic                | RPL_TOPIC    |
| Update Users on Channel Connect | RPL_NAMREPLY |
| Update User/Self on Join        | JOIN         |
| Update User/Self on Kick        | KICK         |
| Update User/Self on Kill        | KILL         |
| Update User/Self on Part        | PART         |
| Update User/Self on Quit        | QUIT         |
| Update User/Self on Nick        | NICK         |

