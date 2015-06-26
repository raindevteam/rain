express = require('express')
router  = express.Router()

router # Index Routes
.get '/', (req, res) ->
  res.render('index', title: 'Rainbot - Home')

module.exports = router