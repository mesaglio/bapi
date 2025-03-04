import http from "k6/http";
import { check } from "k6";

const BASE_URL = "http://localhost:8080";

const NUM_USERS = 100;
const users = Array.from({ length: NUM_USERS }, (_, i) => ({
  username: `user_${i}`,
  email: `user_${i}@test.com`,
}));

export const options = {
  stages: [
    { duration: "10s", target: 10 },
    { duration: "20s", target: 50 },
    { duration: "10s", target: 0 },
  ],
};

function getUserForVU() {
  const vuId = __VU - 1;
  return users[vuId % NUM_USERS];
}

function createUser(user) {
  const res = http.post(`${BASE_URL}/users`, JSON.stringify(user), {
    headers: { "Content-Type": "application/json" },
  });

  check(res, {
    "User created (201)": (r) => r.status === 201,
  });

  return res.status === 201;
}

function modifyUser(user) {
  const payload = {
    username: user.username,
    email: `${user.username}@updated.com`,
  };

  const res = http.patch(
    `${BASE_URL}/users/${user.username}`,
    JSON.stringify(payload),
    { headers: { "Content-Type": "application/json" } }
  );

  check(res, {
    "User modified (200)": (r) => r.status === 200,
  });

  return res.status === 200;
}

function deleteUser(user) {
  const res = http.del(`${BASE_URL}/users/${user.username}`, {
    headers: { "Content-Type": "application/json" },
  });

  check(res, {
    "User deleted (200)": (r) => r.status === 200,
  });

  return res.status === 200;
}

export default function () {
  const user = getUserForVU();

  createUser(user);
  modifyUser(user);
  deleteUser(user);
}
