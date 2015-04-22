// As an example, modules don't need to be written in coffeescript, one can
// still use pure javascript.

module.exports = function(Module) {
   Rainbin = new Module("Rainbin");
   Rainbin = addListeners(commands);
   return Rainbin;
}
