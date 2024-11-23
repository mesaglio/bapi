from models.user import User
from flask import request, Blueprint, jsonify

users = []

users_blueprint = Blueprint("users", __name__)


@users_blueprint.route("/", methods=["POST"])
def create_user():
    try:
        body = User(request.json)
        users.append(body)
        return "", 200
    except Exception as e:
        return "", 400


@users_blueprint.route("/", methods=["GET"])
def get_all_users():
    return jsonify([o.__dict__ for o in users]), 200


@users_blueprint.route("/<username>", methods=["GET"])
def get_user_by_username(username):
    index = get_index_of_user(username)
    if index != -1:
        return jsonify(users[index].__dict__), 200
    return "", 404


@users_blueprint.route("/<username>", methods=["PATCH"])
def update_user_by_username(username):
    body = User(request.json)
    if get_index_of_user(username) != -1:
        delete_user(username)
        users.append(body)
        return "", 200
    return "", 404


@users_blueprint.route("/<username>", methods=["DELETE"])
def delete_user_by_username(username):
    delete_user(username)
    return "", 200


def get_index_of_user(username: str):
    index = 0
    for _user in users:
        if _user.username == username:
            return index
        index += 1
    return -1


def delete_user(username: str):
    index = get_index_of_user(username)
    if index != -1:
        users.pop(index)
