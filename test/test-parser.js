/* globals __config*/
"use strict";

const globals = require('./helpers/globals');
const test = require('tape');
const parser = require('./../lib/parser');

__config.commandPrefix = ';';

const nick = __config.nick;

test("should check if text is a command (normal)", function(t) {
  t.ok(parser.isCommand(nick + ", ;say is a command?"),
        "is command (normal)");
  t.end();
});

test("should check if text is a command (minimal)", function(t) {
  t.ok(parser.isCommand(__config.commandSymbol + " ;say is a command?"),
       "is command (minimal)");
  t.end();
});

test("should check if text is a command (no denotion)", function(t) {
  t.ok(parser.isCommand(";say is a command?"),
       "is command (no denotion)");
  t.end();
});

test('parser should split commands into an array (Test #1)', function(t) {
  const commands = parser.getCommands(';say testing && ;say testing');
  t.equal(commands[0], ';say testing && ;say testing', 'Command[0]');
  t.end();
});

test('parser should split commands into an array (Test #2)', function(t) {
  const commands = parser.getCommands(
        ';say testing \\&\\&\\&\\&&&\\&& ;say commands \\& & ;say test'
      );
  t.equal(commands[0], ';say testing &&&&&&&', 'commands[0]');
  t.equal(commands[1], ';say commands &', 'commands[1]');
  t.equal(commands[2], ';say test', 'commands[2]');
  t.end();
});

test('parser should split commands into an array (Test #3)', function(t) {
  const commands = parser.getCommands(
        ';say testing &&&&&&&&&& commands&;say &&'
      );
  t.equal(commands[0], ';say testing &&&&&&&&&& commands', 'commands[0]');
  t.equal(commands[1], ';say &&', 'commands[1]');
  t.end();
});

test('parser should split commands into an array (Test #4)', function(t) {
  const commands = parser.getCommands(
        ';say testing \\&\\& commands & ;jip true && false & ;cal 4 + 5'
      );
  t.equal(commands[0], ';say testing && commands', 'commands[0]');
  t.equal(commands[1], ';jip true && false', 'commands[1]');
  t.equal(commands[2], ';cal 4 + 5', 'commands[2]');
  t.end();
});

test('parser should split commands into an array (Test #5)', function(t) {
  const commands = parser.getCommands(
        ';say testing \\&\\& commands & & ;say test'
      );
  t.equal(commands[0], ';say testing && commands', 'commands[0]');
  t.equal(commands[1], ';say test', 'commands[1]');
  t.end();
});
