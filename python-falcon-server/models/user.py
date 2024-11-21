class User:
    def __init__(self, _username, _email):
        self.username = _username
        self.email = _email

    @staticmethod
    def validate_dict(obj):
        return all(['username' in obj, 'email' in obj])

    def encode(self):
        return self.__dict__
