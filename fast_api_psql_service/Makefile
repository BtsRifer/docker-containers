venv:
	poetry install --no-root

run:
	docker compose up

build:
	docker compose up --build 

black:
	black .

isort:
	isort . -sl
	
reformat: black isort