var Sequelize = require('sequelize')
var sequelize = require('../config/db.config')
var DataType = Sequelize.DataTypes

var users = require('./users')(sequelize, DataType)
var db = {
    sequelize,
    Sequelize,
    users,
}
module.exports = db