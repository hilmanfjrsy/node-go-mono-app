const Sequelize = require('sequelize')

var db_name = process.env["AUTH_DB_NAME"]
var db_username = process.env["AUTH_DB_USER"]
var db_password = process.env["AUTH_DB_PASSWORD"]
var db_host = process.env["AUTH_DB_HOST"]
var db_dialect = process.env["AUTH_DB_DIALECT"]
console.log("===>>>", db_name, db_username, db_password, db_host, db_dialect)
module.exports = new Sequelize(db_name, db_username, db_password, {
    host: db_host,
    dialect: db_dialect,
    timezone: 'Asia/Jakarta',
    logging: true
});