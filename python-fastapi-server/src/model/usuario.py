from pydantic import BaseModel


class User(BaseModel):
    username: str
    email: str

    def __eq__(self, other):
        return self.username == other.username and self.email == other.email
