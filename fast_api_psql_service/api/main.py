from fastapi import FastAPI
from fastapi import Response
from http import HTTPStatus

app = FastAPI()


@app.get("/")
def get_root():
    return Response(content="OK", status_code=HTTPStatus.OK)
