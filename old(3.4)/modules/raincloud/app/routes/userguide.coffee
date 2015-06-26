express = require('express')
router  = express.Router()

router # Index userguide
.get '/', (req, res) ->
  res.render('userguide', title: 'Rainbot - User Guide')
  
module.exports = router