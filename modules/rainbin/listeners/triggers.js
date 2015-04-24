module.exports = {
  test: {
    event: 'message',
    trigger: function(message) { return true; },
    action: function(respond, done) {
      respond.say('sparks');
      return done()
    }
  }
}
