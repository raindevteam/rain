"use strict";

function getModule(name, modules) {
  for (let module of modules) {
    if (module.name.toLowerCase() === name ||
        module.abbrev.toLowerCase() === name) {
      return module;
    }
  }
}

const mod = {
  help: 'Lists, enables, and disables modules (rainbot.info/userguide#module)',
  ASAP: false,
  whitelist: ['MistaWolf', 'PinkDawn', 'Eventide',
              'Dustrunner', 'King_Sombra'],
  action: function(data, respond, done) {
    const args = data.args;
    if (!args[0]) return done();
    switch (args[0].toLowerCase()) {
      case 'list':
        var names = [];
        for (let module of data.bot.modules) {
          const status = module.status() ? '+' : '-';
          const name = module.abbrev ? module.abbrev : module.name;
          names.push(status + name);
        }
        respond.say(names.join(', '));
        break;
      case 'enable':
        if (data.args[1]) {
          let module = getModule(data.args[1], data.bot.modules);
          if (!module) {
            respond.say('No such module: ' + data.args[1]);
            return done();
          }
          module.enable();
          respond.say('Enabled module ' + module.name);
        }
        break;
      case 'disable':
        if (data.args[1]) {
          let module = getModule(data.args[1], data.bot.modules);
          if (!module) {
            respond.say('No such module: ' + data.args[1]);
            return done();
          }
          module.disable();
          respond.say('Disabled module ' + module.name);
        }
        break;
      case 'show':
        if (data.args[1]) {
          let module = getModule(data.args[1], data.bot.modules);
          if (!module) {
            respond.say('No such module: ' + data.args[1]);
            return done();
          }
          const status = module.status() ? '+' : '-';
          respond.say(status + module.name + ' | ' +
          'Commands: ' + module.commands.keys().length);
        }
        break;
    }
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('mod', mod);
};
