helpers = __core.helpers
defined = __core.defined

class Trigger
  constructor: () ->

  cmd: (text, trigger) ->
    if helpers.isFunction(trigger)
      if text.lower().hasAt(defined.CMD_TRIGGER, 0)
        return trigger(text.after('!', true).nospace())
    match = text.before('!').after(',').trim()
    if trigger
      return match.lower() == trigger
    else if match
      return match[1]

  hasAnywhere: (text, trigger) ->
    return text.has(trigger)

module.exports = Trigger
