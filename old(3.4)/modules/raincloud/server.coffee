express = require 'express'
path = require 'path'
favicon = require 'serve-favicon'
logger = require 'morgan'
cookieParser = require 'cookie-parser'
bodyParser = require 'body-parser'
http = require 'http'

app = express()
server = http.createServer(app)

app.disable('x-powered-by')
app.set 'views', path.join(__dirname, 'views')
app.set 'view engine', 'jade'

app.use favicon(__dirname + '/public/favicon.ico')
app.use logger('dev')
app.use bodyParser.json()
app.use bodyParser.urlencoded(extended: false)
app.use cookieParser()
app.use express.static(path.join __dirname, 'public')

routes = require __dirname + '/app/routes/index'
userguide = require __dirname + '/app/routes/userguide'
about = require __dirname + '/app/routes/about'
development = require __dirname + '/app/routes/development'
console = require __dirname + '/app/routes/console'

app.set 'menu',
    'Home'        : '/'
    'User Guide'  : '/userguide'
    'About'  : '/about'
    'Development' : '/development'

app.use (req, res, next) ->
  app.set('selected', req.originalUrl)
  next()

app.use '/', routes
app.use '/userguide', userguide
app.use '/about', about
app.use '/development', development
app.use '/console', console

# catch 404 and forward to error handler
app.use (req, res, next) ->
  err = new Error('Not Found')
  err.status = 404
  err.message = 'Seems we couldn\'t find the page you requested!'
  next(err)

# error handlers

# development error handler
# will print stacktrace
if app.get('env') == 'development'
  app.use (err, req, res, next) ->
    res.status(err.status or 500)
    res.render 'error',
      message: err.message
      error: err

# production error handler
# no stacktraces leaked to user
app.use (err, req, res, next) ->
  res.status(err.status || 500)
  res.render 'error',
    message: err.message
    error: {}

server.listen(8080, '127.0.0.1')
