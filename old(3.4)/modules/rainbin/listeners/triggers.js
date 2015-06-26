module.exports = {
  test: {
    event: 'message',
    trigger: function(message) { return false; },
    action: function(respond, done) {
      respond.say('sparks');
      return done()
    }
  }
}
