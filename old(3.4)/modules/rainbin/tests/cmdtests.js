expect = require('chai').expect;
Rainbin = require('../index');

describe('Rainbin', function() {
   describe('commands', function() {
      describe('pb', function() {
         it('should respond with "test text"', function() {
            response = Rainbin.testCommand('pb', 'testing', null);
            expect(response).to.equal("test text");
         });
      });
   });
});
