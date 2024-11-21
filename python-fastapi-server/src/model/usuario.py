from pydantic import BaseModel


class User(BaseModel):
    username: str
    email: str

    def __eq__(self, other):
        self.__dict__.get('username') == other.get('username') and self.__dict__.get('email') == other.get('email')
