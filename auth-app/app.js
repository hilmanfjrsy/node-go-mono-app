var createError = require('http-errors');
var express = require('express');
var cookieParser = require('cookie-parser');
var logger = require('morgan');
var path = require('path');
const dotenv = require('dotenv')
dotenv.config()
var port = process.env.PORT || 3000;

var indexRouter = require('./routes/index');
var authRouter = require('./routes/auth');
const db = require('./models');

db.sequelize.sync({ alter: true }).then(() => {
  console.log("success: sync all models");
}).catch(err => {
  console.log(err);
});

var app = express();

// view engine setup
app.set('views', path.join(__dirname, 'views'));
app.set('view engine', 'jade');

app.use(logger('dev'));
app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cookieParser());
app.use(express.static(path.join(__dirname, 'public')));

app.use('/', indexRouter);
app.use('/auth', authRouter);

// catch 404 and forward to error handler
app.use(function (req, res, next) {
  next(createError(404));
});

// error handler
app.use(function (err, req, res, next) {
  // set locals, only providing error in development
  res.locals.message = err.message;
  res.locals.error = err;

  // render the error page
  res.status(err.status || 500);
  res.render('error');
});

app.listen(port, () => {
  console.log(`App listening at http://localhost:${port}`);
});
module.exports = app;
