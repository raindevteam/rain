---
title: Rain | Design Document 
author: Rodolfo Castillo-Valladares
date: 2017-10-30
fontfamily: utopia
fontsize: 12pt 
---

\maketitle
\newpage
\tableofcontents
\newpage

Workflow
--------

The rbot Package
----------------

The handler Package
-------------------

The handler package is an event handler that manages all incoming and outgoing
transmissions with droplets and the Discord session. It also contains a list of
all registered droplets. It is strictly enforced that all methods of
communication with droplets and the Discord session are routed via the handler.

## Dispatchers and Listeners

There are many problems with managing listeners from droplets and the bot itself
when just using the underlying Discord library. The biggest problem being able
to enable and disable specific listeners. This problem is handled by only using
one event dispatch handler per Discord event. Within each dispatcher, a list of
listeners for the corresponding event are fired. These are kept by the handler
in what is known as the registry, providing convenient access and management.

## Actions and Action Handlers


The droplet Package
-------------------

The internal Package
--------------------

The rain Package
----------------

The hail Package
----------------

The seer Package
----------------
