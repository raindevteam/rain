// We only pass bot to modules so that they can have a reference. So
// I pass a "dummy" bot to the module. As long as the mock gets passed around
// correctly (regardless of it not being a real bot instance) everything
// should be fine.
Module = require('./../../lib/bot/module')({ iambroot: true });

module.exports = {

  makeModule() {
    return new Module('TestModule', 'tm');
  },

  addValidCommands(module) {
    module.addCommands({
      'echo': {
        ASAP: false,
        action(data, respond, done) {
          respond.say('from echo');
          return done();
        }
      },
      'say': {
        ASAP: true,
        action(data, respond, done) {
          respond.say('test command 3');
          return done();
        }
      }
    });
  }
};
