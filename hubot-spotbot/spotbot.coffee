# Description:
#   Spotbot integration plugin
#
# Dependencies:
#   "pubnub": "3.13.x"
#
# Configuration:
#   PUBNUB_PUBLISH_KEY
#   PUBNUB_SUBSCRIBE_KEY
#
# Commands:
#   hubot stop
#   hubot play
#   hubot skip
#
# Author:
#   ahamidi

PUBNUB = require 'pubnub'
pubnub = PUBNUB.init(
  publish_key: process.env.PUBNUB_PUBLISH_KEY,
  subscribe_key: process.env.PUBNUB_SUBSCRIBE_KEY,
  ssl: true)

module.exports = (robot) ->
  robot.respond /music (.*)/i, (res) ->
    command = res.match[1]
    publishMsg command, (m) ->
      res.reply "#{m[1]}"

publishMsg = (msg, cb) ->
  jsonMsg =
    command: "#{msg}"

  pubnub.publish
    channel: 'music'
    message: jsonMsg
    callback: cb
