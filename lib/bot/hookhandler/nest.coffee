resultNest = {}

module.exports =
  fillNests: (text) ->
    nests = text.match(/\{(.*?)\}/g)
    if nests?.length > 0
      for nest in nests
        nest = nest.replace('{', '').replace('}', '')
        if resultNest[nest]
          text = text.replace('{'+nest+'}', resultNest[nest].response)
    return text

  nest: (name, respondData) ->
    thisNest = resultNest[name] = {}
    if respondData.output then thisNest.response = respondData.output
    else
      thisNest.response = ''
      for response in respondData.responses
        thisNest.response += " " + response.data
        thisNest.response = thisNest.response.trim()

  reset: () ->
    resultNest = {}
