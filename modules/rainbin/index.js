// As an example, modules don't need to be written in coffeescript, one can
// still use pure javascript.
commands = require('./listeners/commands');
triggers = require('./listeners/triggers');

module.exports = function(Module) {
   Rainbin = new Module("Rainbin");
   console.log(Module)
   Rainbin.addCommands(commands);
   Rainbin.addTriggers(triggers)
   console.log(Rainbin);
   return Rainbin;
}
