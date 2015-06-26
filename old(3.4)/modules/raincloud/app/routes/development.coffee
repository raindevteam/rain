express = require('express')
router  = express.Router()

router # Index development
.get '/', (req, res) ->
  res.render('development', title: 'Rainbot - Development')
  
module.exports = router