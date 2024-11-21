import uvicorn as uvicorn
from fastapi import FastAPI, Request
from fastapi.exceptions import RequestValidationError
from fastapi.responses import JSONResponse

from src.routes.ping_route import ping
from src.routes.users_route import users


def create_app():
    app = FastAPI()
    @app.exception_handler(RequestValidationError)
    async def validation_exception_handler(request: Request, exc: RequestValidationError):
        return JSONResponse(
            status_code=400,  # Cambiar 422 a 404
            content={},
        )
    
    app.include_router(ping)
    app.include_router(users)
    return app


if __name__ == "__main__":
    app = create_app()
    uvicorn.run("main:create_app", host="0.0.0.0", port=8080, reload=False, lifespan="on")
