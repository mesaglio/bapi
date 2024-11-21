import os
import requests
import pytest
import json
import random

port = os.getenv('PORT', 8080)
base_url = f'http://localhost:{port}'
ping_url = base_url + '/ping'
users_url = base_url + '/users'
users = [
    {
        "username": "juan",
        "email": "juan@gmail.com"
    },
    {
        "username": "test",
        "email": "test@gmail.com"
    }
]


def eq_obj(obj1, obj2):
    items_obj1 = obj1.items()
    items_obj2 = obj2.items()
    return all(item in items_obj1 for item in items_obj2)


def change_email(obj):
    obj['email'] = obj['email'].replace("@",f"{random.randint(30000000,40000000)}@")
    return obj


def test_ping():
    response = requests.get(ping_url)
    assert response.status_code == 200
    assert 'Pong' in response.text

def test_empyt_users():
    empty_users = requests.get(users_url)
    assert empty_users.status_code == 200
    assert empty_users.content == b'[]'


def test_bad_requests():
    post_bad_request = requests.post(users_url, json={})
    assert post_bad_request.status_code == 400


def test_not_found():
    url = f"{users_url}/{random.choice(users).get('username')}"
    print(url)
    patch_bad_request = requests.patch(url, json={})
    assert patch_bad_request.status_code == 404


def test_delete_not_found():
    delete_request = requests.delete(f"{users_url}/{random.choice(users).get('username')}")
    assert delete_request.status_code == 200


def test_add_users():
    for user in users:
        response = requests.post(users_url, json=user)
        assert response.status_code == 200


def test_get_users():
    for user in users:
        response = requests.get(f"{users_url}/{user.get('username')}")
        assert response.status_code == 200
        assert eq_obj(user, json.loads(response.content))

def test_update_users():
    new_users = list(map(lambda o: change_email(o),users))
    for user in new_users:
        response = requests.patch(f"{users_url}/{user.get('username')}", json=user)
        assert response.status_code == 200
        response_get = requests.get(f"{users_url}/{user.get('username')}")
        assert response_get.status_code == 200
        assert eq_obj(user,json.loads(response_get.content))


def test_delete_all_users():
    for user in users:
        request = requests.delete(f"{users_url}/{user.get('username')}")
        assert request.status_code == 200
    all_users = requests.get(users_url)
    assert all_users.status_code == 200
    assert b'[]' in all_users.content
