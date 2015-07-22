## Blog For Development Updates!
http://rainbot-irc.blogspot.com/

## This README
Is incomplete. I've now however added an installation section! Now I got to
start work on API documentation!

# RainBot
Multipurpose, extensible IRC bot written with node.js

## Overview
RainBot is just another IRC bot among the many already made; it started as something for fun rather than need. That being said, plenty of development has still been going into RainBot, well at least as much as a college student can do without drowning themself in work. The master branch is for versions of the bot that are somewhat stable, but not entirely, stable enough that the bot can be used to a reasonable extent. The development branch is where I do all my work and reflects the most up-to-date changes, bugs, and features.

### A Few Things RainBot Can Do

#### Modules
The Bot has been reworked quite a few times to provide easy extensibility via modules.
The result of such efforts? Hard to judge, but see for youself!

Here's a small example showing how you can make a module sporting an echo command,

    module.exports = function(Module, load) {

      // Make a new module!
      myModule = new Module('myModule', 'mm');

      // What good is a module without hooks!?
      myModule.addCommand('echo', {
        help: "Repeats its arguments!",
        action(data, respond, done) {
          if (data.args) respond.say(data.args.join(' '));
          return done();
        }
      });

      // Shut up and take my module!
      return load(myModule);
    }


#### Command Nesting
Command nesting lets you string together commands and pass values forward as
well! Here's some brief examples:

    ;cal f(x) = x * x & ;cal f(3)
    ;roll & ;cal f({roll})

    // Hint, hint: Run these in order!

#### Command Aliasing
Command aliasing let's you take a string of commands  and group them to a singular
name. Also, you can declare parameters to pass to the alias,

    ;alias quote -> ;say "I am broot"
    ;alias enlighten x -> ;say Hey {x}, i am broot!
    ;alias kick person -> ;do kicks {person}

Currently you can only set and remove aliases, but listing aliases will be implemented soon.


## Cool, I Want One!
Back in the day, when RainBot was just a teeny-tiny little bot, growing up
all adorably in a sweetful laid back environment, RainBot wasn't meant to
be used by others. It's main purpose was to serve it's home channel #Neverfree.
But now, just recently, I've steered development in a different direction. The
direction of "maybe someone can actually find their own use for this heap of, err..
bot."

So here it is! My first attempt ever at guiding you through an installation.

#### Get io.js Installed

First thing you will need to do is download io.js and get it working on your system.
You can find it [here](https://iojs.org/en/index.html). It is basically node.js with default harmony features set (which the bot uses, so don't think you can get away easily with just node.js). You could also use the nvm manager found [here](https://github.com/creationix/nvm) (if on windows, check [here](https://github.com/coreybutler/nvm-windows))

Speaking of windows, this installation might go much smoother on linux systems, just
an fyi.

Once you have made certain that io.js is installed, you should also now have npm!
It's the package manager for node modules, and you'll be using it a bit to get RainBot
on its feet.

#### Download RainBot

Let us now download RainBot, you're most likely on the repo page already if you're
reading this, if not here it [is](https://github.com/Wolfchase/rainbot). You can go ahead and download from the releases, download a zip or tar from the master branch or just clone it and pull when you need to update, there's a thought!

#### Run npm on RainBot and Modules

Once you got the bot on your machine, your're almost ready to fly! Well actually
if you're on windows it might take a bit longer. You'll see why in a bit. Go to
where ever you put RainBot, and let's run this npm command (To clarify, go into the folder
that contains the source of the bot, you should see a package.json file),

    npm install

After running that command a few things should happen, and may vary if you're on...
yep windows!

If you're on Linux the installation probably had a build error but installed anyway. If the build error was about some optional dependency from the IRC library then there's no need to worry, it was an optional dependency afterall (the problem seems common too, but really there's no need to worry).

On Windows however, you may have some more fatal errors. If one such error happens to be
a node-gyp build error, you might need to install [MS Visual C++](https://www.visualstudio.com/en-us/downloads/download-visual-studio-vs). Yeah, sorry
about that. If you got any other problems related to node-gyp go read it's installation
section on the readme [here](https://github.com/TooTallNate/node-gyp/). You don't have
to actually install anything, just make sure you meet the requirements.

RainBot also comes with some modules that need the npm install command ran for. If you
look at the modules directory, perform what just did for the modules rcal and rbrella. Both their directories will have a package.json. And for Windows users... there should
be no gotcha this time.

#### Set up the Config

At this point you're practically done, just set up your settings in the config.js
file found in the config file. The process.env things are for those who know about node, and if you don't, you should totally read up on it a little! Otherwise feel free to get rid of them and put your own settings in there. I've left comments to guide you!

#### Fly!

Once you're done mucking about in the config, you're ready for lift off! Simply find
the nearest convenient terminal or command prompt and run

    iojs rainbot.js

    // or if you used nvm

    node rainbot.js

Yeah this was pretty long to read, but it was aimed at those who didn't know much about
node. If you've had enough experience using node, then you maybe didn't even have to read this.

Anyhow hopefully it helped, and if not, then email me at rodolfo.valladares.c@gmail.com, leave a comment on my blog or find me on the IRC server: irc.canternet.org on the #RainBot channel.

#### Whitelist Some Commands if you Wish

If you want to whitelist some commands, its not entirely that hard, but I guess
I do need to find a simpler way.

Whitelist for commands are added via properties. So you'll need to find
where the command is declared. Don't worry though, I've set my directory structures
for modules to be fairly straightforward. As an example, let's add a whitelist to the
command 'die.' The die command is found inside the rcore module, here's a directory
listing and a pointer showing the file,

    modules...
      rbrella
      rcal
      rchat
      rcore...
        hooks...
          commands...
            alias.js
            chan.js <--
            die.js
            help.js
            mod.js
            ver.js
        index.js

If you open the file you'll find an object like this

    const die = {
      help: 'Best you don\'t use it, unless the bot really needs to go bye bye',
      ASAP: true,
      action: function(data, respond, done) {
        process.exit(0);
      }
    };

Now you probably don't want everyone having access to this command so add a whitelist
property with nicks,

    const die = {
      help: 'Best you don\'t use it, unless the bot really needs to go bye bye',
      ASAP: true,
      whitelist: ['broot', 'you', 'someotherpersonwhocanusethiscommand'],
      action: function(data, respond, done) {
        process.exit(0);
      }
    };

And there you have it, now only the nicks in the whitelist array can use this command.
I understand it isn't the safest way to whitelist commands however, and I am in the process of working a better solution.

## Modules
Todo
