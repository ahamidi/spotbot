# Description:
#   Spotbot integration plugin
#
# Dependencies:
#   "ioredis": "1.15.x"
#
# Configuration:
#   REDIS_URL
#   REDIS_CHANNEL
#
# Commands:
#   hubot stop
#   hubot play
#   hubot skip
#
# Author:
#   ahamidi

Redis = require 'ioredis'
redis = new Redis(process.env.REDIS_URL)

module.exports = (robot) ->
  robot.respond /music (.*)/i, (res) ->
    command = res.match[1]
    publishMsg command, (m) ->
      res.reply m

publishMsg = (msg, cb) ->
  jsonMsg =
    command: "#{msg}"
  redis.publish('music', JSON.stringify(jsonMsg))
  cb "sent!"
