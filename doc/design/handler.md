---
title: Rain Design | Handler Package 
author: Rodolfo Castillo-Valladares
date: 2018-3-21
rain-version: 0.1.0
---

Overview
--------

The handler package is an event handler that manages all incoming and outgoing
transmissions with droplets and the Discord session. It also contains a list of
all registered droplets. It is strictly enforced that all methods of
communication with droplets and the Discord session are routed via the handler.

Dispatchers and Listeners
-------------------------

There are many problems with managing listeners from droplets and the bot itself
when just using the underlying Discord library. The biggest problem being able
to enable and disable specific listeners. This problem is handled by only using
one event dispatch handler (a function) per Discord event. Within each
dispatcher, a list of listeners for the corresponding event are fired. These are
kept by the handler in what is known as the registry, which provides convenient
access and management.

Ideology Behind the Creation and Running of Listeners
------------------------------------------------------

The way listeners are created and ran within the handler may seem convoluted at
first glance, but breaking it down into smaller pieces helps to understand why
the implementations in question exist.

### Understanding the Problem

The discord library that Rain uses does not currently allow us to remove or
disable listeners. Also, since the listeners are simply functions, it is hard to
associate metadata to the functions; data such as to whom the function belongs
to becomes difficult to implement.

With this in mind, the following section explains how a listener is created
using an interface to provide metadata to functions passed to the underlying
discord library.

### Listener Creation

As said, a `Listener` inteface is used to describe the handler's own listeners.
Because we are using an interface, we can encapsulate all types of listeners
into a singular type.

The listener interface also has a setter to allow it to recieve a function to be
ran on demand. Therefore, it is only necessary to define a function as a
listener, this function can then be passed to the handler which will encapsulate
it in a new struct that abides by the Listener interface.

### Running Listeners

Up until now, the implementation of listeners has been simple. The act of
running the listeners is the real reason for most of the code in this package.
Because we are using an interface to encapsulate different types of functions, a
new type of problem is incurred, that being how to pass specific data to this
interface from a dispatcher. 

A dispatcher function receives a very specific discord data type that has data
associated to the firing event. Because the handler's listeners use a variable
interface to store functions that will be ran by a listener, it becomes apparent
that we cannot simply pass specific discord data into this interface.

Because we need to cast our variable interface in question with an acceptable
function prototype, one which can accept our discord data type from a firing
event, a system to enable this has been put in place using an interface and
a switch at time of creation for a particular listener.

Firstly, let's define the `Action` interface. It defines a singular function to
be satisfied, the `Do` function. Within this function, an 'action handler' can
appropriately cast an interface into an appropriate discord data type. This
allows us to pass data from a discord event without ever specifying its type
right up until a listener's function needs it to run, which also happens within
the `Do` function.

Action handlers are generated along side dispatchers in dispatchers.go at build
time.

With all of this in place, there is only one problem left to solve: how to cast
a listener's function into an appropriate prototype. Our solution is to simply
cast the intial passed function, at the time of a listener's creation, into an
appropriate action handler. Action handlers are merely type'd functions with a
`Do` method to satisfy the `Action` interface. We can then store the action
handler within any listener.

To finish it all up, a listener runs by invoking it's action handler's `Do`
method which accepts an interface, and casts that interface into an
appropriate discord type.

### Optimization

Trying to find better ways to handle implementations explained in the previous
sections has been fruitless. That said, what the current implementations offer
should not be so easily dismissed. Creating listeners requires very little effort
as a function needs only be defined. Also, because said function is casted into
an action handler at creation, type casting of the function is never done during
the firing of an event, keeping casting to a minimum.
 
