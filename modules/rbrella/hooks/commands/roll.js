"use strict";

function randomIntFromInterval(min, max) {
  return Math.floor(Math.random() * (max - min + 1) + min);
}

const roll = {
  help: '_r _d _p | r: Repetitions, d: Size of dice, p: Pairs to roll at once (rainbot.info/userguide#roll)',
  action(data, respond, done) {
    const args = [];
    const pargs = data.args.join(' ');
    args[0] = pargs.match(/(\d+)r/);
    args[1] = pargs.match(/(\d+)d/);
    args[2] = pargs.match(/(\d+)p/);
    if (args.length > 3) {
      respond.say('Too many params');
    } else {
      if (args[0] && isNaN(args[0][1]))
        respond.say('Param 1 is not a number');
      else if (args[1] && isNaN(args[1][1]))
        respond.say('Param 2 is not a number');
      else if (args[2] && isNaN(args[2][1]))
        respond.say('Param 3 is not a number');
      else {
        const reps = args[0] ? args[0][1] : 1;
        const dice = args[1] ? args[1][1] : 6;
        const pair = args[2] ? args[2][1] : 1;

        let rolls = [];

        for (let i = 0; i < reps; i++) {
          let sum = 0;
          for (let i = 0; i < pair; i++ ) {
            sum += randomIntFromInterval(1, dice);
          }
          rolls.push(sum);
        }

        respond.say(rolls.join(', '));
      }
    }
    return done();
  }
};

module.exports = function(module) {
  module.addCommand('roll', roll);
};
