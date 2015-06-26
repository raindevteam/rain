config = require(__config)
name = config.nick.toLowerCase()

module.exports =
  'IDENTIFY_MESSAGE': 'You have 30 seconds to identify to your nickname before it is changed.'
  'CMD_TRIGGER': name + '!'
  'MSG_TRIGGER': name + ','
