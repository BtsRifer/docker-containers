FROM  python:3.12-slim-bullseye

RUN apt-get update
RUN apt-get -y install libpq-dev gcc
RUN pip3 install --upgrade pip
RUN pip3 install pipenv

ENV PIPENV_VENV_IN_PROJECT=1