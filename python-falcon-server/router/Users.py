import falcon
from models.user import User
import json


class Users:
    users = []

    def on_post(self, req, res):
        user = req.media
        if not User.validate_dict(user):
            res.status = falcon.HTTP_400
        else:
            self.users.append(self.body_to_user(user))
            res.status = falcon.HTTP_200

    def on_get(self, req, res, username=None):
        if username:
            user = self.find_user_by_username(username)
            if user:
                res.text = json.dumps(user, default=lambda o: o.encode(), indent=4)
            else:
                res.status = falcon.HTTP_404
        else:
            res.text = json.dumps(self.users, default=lambda o: o.encode(), indent=4)

    def on_delete(self, req, res, username=None):
        self.remove_user_by_username(username)

    def on_patch(self, req, res, username=None):
        user = req.media
        if not username or self.find_user_by_username(username) is None:
            res.status = falcon.HTTP_404
        elif not User.validate_dict(user):
            res.status = falcon.HTTP_400
        else:
            self.update_user(username, self.body_to_user(user))

    def find_user_by_username(self, username: str):
        for user in self.users:
            if user.username == username:
                return user
        return None

    def remove_user_by_username(self, username: str):
        if username is None:
            return
        for user in self.users:
            if user.username == username:
                self.users.remove(user)

    def update_user(self, username, new_user: User):
        for user in self.users:
            if user.username == username:
                user.email = new_user.email

    @staticmethod
    def body_to_user(user):
        return User(user['username'], user['email'])
