module.exports = {
   /* pb command */
   '++': {
     action: function(args, respond, done) {
       respond.say(args[0] + 1);
       return done();
     }
   },

   echoArgs: {
      action: function(args, respond, done) {
         respond.say(args.join(' '));
         return done();
      }
   }
}
