import falcon
from falcon import Request, Response
from router.Ping import Ping
from router.Users import Users

class LogMiddleware:
    def process_response(
        self,
        req: Request,
        resp: Response,
        resource: object,
        req_succeeded: bool
    ) -> None:
        
        print(f"Request: {req.method} {req.relative_uri} - Response: {resp.status}")

def map_routes(_app):
    _app.add_route('/ping', Ping())
    _app.add_route('/users', Users())
    _app.add_route('/users/{username}', Users())

def create_app():
    app = falcon.App(middleware=[LogMiddleware()])
    map_routes(app)
    return app


app = create_app()
