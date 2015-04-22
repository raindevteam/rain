module.exports = {
   /* pb command */
   pb: {
      nest: true,
      ASAP: false,
      action: function(args, respond, done) {
         var randomTestText = "test text";
         respond.say randomTestText;
         return done();
      }
   }
}
