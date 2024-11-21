import { User } from "./../models/user.ts";
import { Request, Response } from "https://deno.land/x/oak/mod.ts";


let users: Array<User> = [];

const getAllUsers = ({ response }: { response: any }) => {
	response.body = users;
};

const getUserByUsername = ({ params, response }: {
	params: { username: string };
	response: any;
}) => {
	const user =
		users.filter((listUser: User) => listUser.username === params.username)[0];
	if (user) {
		response.status = 200;
		response.body = user;
	} else {
		response.status = 404;
		response.body = { msg: "Not found" };
	}
};

const createUser = async ({ request, response }: {
	request: Request;
	response: Response;
}) => {
    const body = await request.body.json();
    try {
        const user: User = userFromBody(body);
        users.push(user);
        response.status = 200;
    } catch (error) {
        response.status = 400;
        response.body = { msg: error.message };
        return;
    }
};

const updateUser = async ({
	params,
	request,
	response,
}: {
	params: { username: string };
	request: any;
	response: any;
}) => {
	const user =
		users.filter((listUser: User) => listUser?.username === params.username)[0];
	if (user) {
		const newUser = await request.body.json();
		user.username = newUser.username;
		user.email = newUser.email;
		response.status = 200;
		response.body = user;
	} else {
		response.status = 404;
		response.body = { msg: "Not found" };
	}
};

const deleteUser = ({
	params,
	response,
}: {
	params: { username: string };
	response: any;
}) => {
	users = users.filter((user) => user.username !== params.username);
	response.status = 200;
};

function userFromBody(body: any): User {
    if (body.username === undefined || body.email === undefined) {
        throw new Error("Invalid user data. Please check the data and try again.");
    }

    const user: User = {
        username: body.username,
        email: body.email,
    };
    return user;
}

export { getAllUsers, getUserByUsername, updateUser, deleteUser, createUser };

