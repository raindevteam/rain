express = require('express')
router  = express.Router()

router # About Routes
.get '/', (req, res) ->
  res.render('about', title: 'Rainbot | About')

module.exports = router
