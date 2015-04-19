msg = require('./listeners/msg')

module.exports = (Module) ->
  X = new Module('X') # Create the module
  X.addListeners(msg) # Add msg listeners
  return X            # Return the module
