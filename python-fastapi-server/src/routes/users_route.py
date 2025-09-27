from typing import List, Optional

from fastapi import APIRouter, Response

from src.model.usuario import User

users_db: List[User] = []

users = APIRouter()


@users.get("/users")
async def get_user():
    return users_db


@users.post("/users", status_code=201)
async def create_user(user: User):
    users_db.append(user)
    return user


@users.get("/users/{username}", status_code=200)
async def get_user_by_username(username: str, response: Response):
    user = find_user(username)
    if user:
        return user
    response.status_code = 404
    return response


@users.delete("/users/{username}", status_code=204)
def delete_user(username: str, response: Response):
    user = find_user(username)
    if user:
        users_db.remove(user)
        return Response(status_code=204)
    response.status_code = 404
    return response


@users.patch("/users/{username}")
def update_user(username: str, req_user: User, response: Response):
    user = find_user(username)
    if user:
        users_db.remove(user)
        users_db.append(req_user)
        response.status_code = 200
    else:
        response.status_code = 404
    return response


def find_user(username: str) -> Optional[User]:
    for u in users_db:
        if u.username == username:
            return u
    return None
