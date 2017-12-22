---
title: Rain Design | Parser Package 
author: Rodolfo Castillo-Valladares
date: 2017-10-30
rain-version: 0.1.0
---

Overview
========

The parser package is suited with various string parsing functions for user
and bot data manipulation. It can parse commands and do other string validation
duties.

Command Parsing
===============

The parser can take appropriate commands from Discord users in the form of
strings. It breaks these strings down into a corresponding struct filled with
information that the handler can read. There are primarily three types of
command examples which the following sections cover.

Internal Command
----------------

The following example is an internal command.

        !!info

The two exclamation marks identify the command as internal. The following struct
is created after successful parsing.

        { owner: "__INTERNAL__", command: "info" }

Droplet Command
---------------

The following example is a droplet command.

        !roll

It produces the following struct.

        { owner: "", command: "roll" }

The struct's owner field is left empty signifying to the handler that it should
use the first droplet it can find possesing the command "roll."

Droplet Specified Command
-------------------------

The previous example suffers from the problem of not being able to specify which
droplet to use for a command. In case of multiple droplets using the same name
for a command, the user can instead specify the droplet's name in the input
string. The following example shows the format.

        !game:roll

This would produce the following struct.

        { owner: "game", command: "roll" }

Command Struct Processing
-------------------------

For information on how the struct data is used, please read the <section> in the
handler package's design document.