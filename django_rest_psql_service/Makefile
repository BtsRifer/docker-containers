black:
	black --verbose .

isort: 
	isort --verbose .

reformat: black isort

migration:
	docker compose run api pipenv run api/manage.py makemigrations &&\
	docker compose run api pipenv run api/manage.py migrate

superuser:
	docker compose run api pipenv run api/manage.py createsuperuser