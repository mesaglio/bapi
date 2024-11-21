"use strict";
var _ = require("lodash");
var httpErrors = require("http-errors");

module.exports = async function (fastify, opts) {
  fastify.get("/", async function (request, reply) {
    return "pong";
  });
};
