import { Router } from "https://deno.land/x/oak/mod.ts"

import {
    getAllUsers,
    getUserByUsername,
    deleteUser,
    updateUser,
    createUser
} from "../controllers/userController.ts"

import {
    ping
} from "../controllers/statusController.ts"

const router = new Router();

router
    .post("/users", createUser)
    .get("/users", getAllUsers)
    .get("/users/:username", getUserByUsername)
    .delete("/users/:username", deleteUser)
    .patch("/users/:username", updateUser)
    .get("/ping", ping);

export default router;
