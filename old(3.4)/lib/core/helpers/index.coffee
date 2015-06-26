exports.isFunction = (functionToCheck) ->
  getType = {}
  functionToCheck and getType.toString
  .call(functionToCheck) == '[object Function]'