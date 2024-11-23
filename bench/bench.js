import http from "k6/http";
import { check, sleep } from "k6";

const BASE_URL = "http://localhost:8080"; // Cambia esto a tu URL base

const user = {
  username: "user1",
  email: "test@test.com",
};

export const options = {
  stages: [
    { duration: "30s", target: 10 }, // Escala a 10 usuarios virtuales
    { duration: "1m", target: 50 }, // Mantén 50 usuarios virtuales
    { duration: "30s", target: 0 }, // Reduce gradualmente a 0 usuarios
  ],
};

function createUser(user) {
  const res = http.post(`${BASE_URL}/users`, JSON.stringify(user), {
    headers: { "Content-Type": "application/json" },
  });

  check(res, {
    "create user: status is 200": (r) => r.status === 200,
  });

  if (res.status === 200) {
    return {};
  }

  return null;
}

function modifyUser(user) {
  const payload = {
    username: `${user.username}`,
    email: `${user.email.split("@")[0]}_updated@example.com`,
  };
  const res = http.patch(
    `${BASE_URL}/users/${user.username}`,
    JSON.stringify(payload),
    {
      headers: { "Content-Type": "application/json" },
    }
  );

  check(res, {
    "modify user: status is 200": (r) => r.status === 200,
  });

  return res.status === 200;
}

function deleteUser(user) {
  const res = http.del(`${BASE_URL}/users/${user.username}`, {
    headers: { "Content-Type": "application/json" },
  });

  check(res, {
    "delete user: status is 200": (r) => r.status === 200,
  });

  return res.status === 200;
}

export default function () {
  if (Math.random() < 0.5) {
    // Crear un nuevo usuario
    createUser(user);
  } else {
    if (Math.random() < 0.5) {
      // Modificar el usuario
      modifyUser(user);
    } else {
      // Eliminar el usuario y quitarlo del array local
      deleteUser(user);
    }
  }
}
