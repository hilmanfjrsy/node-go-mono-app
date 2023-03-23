const db = require("../models")
const bcrypt = require('bcrypt')
const jwt = require("jsonwebtoken")
const Ajv = require("ajv")
const ajv = new Ajv()
const phoneRegex = /^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im
class AuthController {
  login = async (req, res, next) => {
    try {
      const schema = {
        type: "object",
        properties: {
          phone: { type: "string" },
          password: { type: "string" },
        },
        required: ["phone", "password"],
        additionalProperties: false
      }
      const body = req.body
      const validSchema = ajv.validate(schema, body)
      if (!validSchema) {
        return res.status(400).send({ message: "invalid body" })
      }
      if (!phoneRegex.test(body.phone)) {
        return res.status(400).send({ message: "invalid phone number" })
      }

      let checkUsers = await db.users.findOne({ where: { phone: body.phone }, raw: true })

      if (checkUsers) {
        let checkPassword = await bcrypt.compare(body.password, checkUsers.password)
        if (checkPassword) {
          var token = jwt.sign(
            {
              "name": checkUsers["name"],
              "phone": checkUsers["phone"],
              "role": checkUsers["role"]
            },
            process.env.SECRET,
            {
              expiresIn: 7200 // 2 hours
            }
          );

          return res.status(200).send({ token })
        } else {
          let message = "invalid phone/password"
          return res.status(400).send({ message })
        }
      } else {
        let message = "invalid phone/password"
        return res.status(400).send({ message })
      }
    } catch (error) {
      return res.send(error)
    }
  }

  register = async (req, res, next) => {
    try {
      const schema = {
        type: "object",
        properties: {
          phone: { type: "string" },
          name: { type: "string" },
          role: { type: "string" }
        },
        required: ["phone", "name", "role"],
        additionalProperties: false
      }
      const body = req.body
      const validSchema = ajv.validate(schema, body)
      if (!validSchema) {
        return res.status(400).send({ message: "invalid body" })
      }
      if (!phoneRegex.test(body.phone)) {
        return res.status(400).send({ message: "invalid phone number" })
      }

      let checkUsers = await db.users.findOne({ where: { phone: body.phone } })
      if (!checkUsers) {
        let randomPassword = Math.random().toString(36).slice(2, 6);
        let hash = await bcrypt.hash(randomPassword, 10);
        body.password = hash

        let users = await db.users.create(body)
        users.password = randomPassword
        return res.status(201).send({
          data: users,
          message: "success"
        })
      } else {
        let message = "phone already exist"
        return res.status(400).send({ message })
      }
    } catch (error) {
      return res.send(error)
    }
  }

  verify = async (req, res, next) => {
    try {
      const schema = {
        type: "object",
        properties: {
          token: { type: "string" }
        },
        required: ["token"],
        additionalProperties: false
      }
      const body = req.body
      const validSchema = ajv.validate(schema, body)
      if (!validSchema) {
        return res.status(400).send({ message: "invalid body" })
      }

      return jwt.verify(body.token, process.env.SECRET, async function (err, decoded) {
        if (!err) {
          return res.status(200).send({
            data: decoded,
          })
        }
        return res.status(401).send(err)
      })
    } catch (error) {
      return res.send(error)
    }
  }
}

module.exports = new AuthController()