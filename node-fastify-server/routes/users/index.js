"use strict";
var _ = require("lodash");
var httpErrors = require("http-errors");

var users = [];

module.exports = async function (fastify, opts) {
  fastify.get("/", async function (request, reply) {
    return users;
  });

  fastify.post("/", async function (request, reply) {
    var user = getUserFromBody(request);
    if (haveAllProps(user)) {
      users.push(user);
      return reply.send();
    }
    return sendError(reply, 400);
  });

  fastify.get("/:username", async function (request, reply) {
    var username = request.params.username;
    if (username === undefined || username === "username") {
      return sendError(reply, 404);
    }
    var [exist, index] = userExist(username);
    if (!exist) {
      return sendError(reply, 404);
    }
    return reply.send(users[index]);
  });

  fastify.delete("/:username", async function (request, reply) {
    var username = request.params.username;
    deleteUser(username);
    return reply.send();
  });

  fastify.patch("/:username", async function (request, reply) {
    var username = request.params.username;
    var [exist, _] = userExist(username);

    if (!exist) return sendError(reply, 404);

    var user = getUserFromBody(request);
    if (haveAllProps(user)) {
      deleteUser(username);
      users.push(user);
      return reply.send();
    }

    return sendError(reply, 400);
  });
};

function sendError(reply, code) {
  return reply.code(code).send();
}

function userExist(username) {
  const index = getUserIndex(username);
  return [index !== -1, index];
}

function getUserIndex(username) {
  var index = _.findIndex(users, { username: username });
  return index;
}

function deleteUser(username) {
  _.remove(users, { username: username });
}

function haveAllProps(user) {
  return _.values(user).every((prop) => {
    return prop !== undefined;
  });
}

function getUserFromBody(req) {
  var { username, email } = req.body;
  return {
    username: username,
    email: email,
  };
}
