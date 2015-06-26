express = require('express')
router  = express.Router()
bot = __core.bot

console.log(bot.chans)

router # Console Routes
.get('/', (req, res) ->
  res.render('console', title: 'Rainbot | Console')
).post('/chat', (req, res) ->
  res.render('partial/chat', { channels: bot.chans })
).post('/logs', (req, res) ->
  res.render('partial/logs')
)

module.exports = router
