### __Important!__

__Rain is no longer an IRC library. We are now only supporting Discord!__ 
__Also we're in very early development and things might not be tidy!__

----------------------------------------------------------------------------------------------------
<img align="right" src="https://avatars1.githubusercontent.com/u/20262521?v=3&s=200">

Rain, a Discord Bot
===================
__With User Plugins (Droplets)__

### Overview

1. [What is Rain](#rain)
2. [Development](#development)
2. [Droplets](#premade)

What is Rain<a name="rain"></a>
-------------------------------

Rain is a discord bot aimed at providing user plugin capabilities. Basically the
idea is to have a bot that can accept user created "droplets", which are
essentially plugins that connect to the bot via MessagePack RPC.

Development<a name="development"></a>
-------------------------------------
[![codebeat badge](https://codebeat.co/badges/fc8f57aa-7ccb-4502-adf1-5dd90af9d08c)](https://codebeat.co/projects/github-com-raindevteam-rain-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/raindevteam/rain)](https://goreportcard.com/report/github.com/raindevteam/rain)
[![Code Climate](https://codeclimate.com/github/raindevteam/rain/badges/gpa.svg)](https://codeclimate.com/github/raindevteam/rain)
[![codecov](https://codecov.io/gh/raindevteam/rain/branch/master/graph/badge.svg)](https://codecov.io/gh/raindevteam/rain)

We're currently in early pre-0.1.0 development, but getting there! You can
follow along with development on Rain's Trello [board](https://trello.com/b/ozUA2ATm/rain). It may be a bit empty
right now but I'll soon start to update it.

Overall, we currently have the following guidelines to meet before getting to
0.1.0:

1. Get a solid foundation going.

    We don't want to regret any decisions made now, in the future. I am doing my best
    to lay groundwork so that maintenance in the future is less of a chore.

2. Ensure room for proper debugging.

    With the hail and rain sub packages, I hope to achieve a more approahable
    way to debugging the bot when things go awry live.

3. Consistency.

    We need consistency so that in the future I don't have to worry about
    whether I should do something the x way or y way.

That's all I got for now, these overlap a bit with each other but I believe they are still specific enough.

> **Note:** If you want to contribute, you'll have to wait a bit. Once I feel I have a good 0.1.0 version, I'll consider creating a CONTRIBUTING.md

Droplets<a name="droplets"></a>
-------------------------------

Currently the droplet system is not implemented yet. We are still working on 
making sure internal listeners and commands work. After all that is done
and tested, I'll move on to making the commands. It's a one man army here so 
work will be a little slow. I've got college too...
