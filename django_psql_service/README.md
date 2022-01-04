## Django REST with PostgreSQL service

A container with a Django application setup with a PostgreSQL database connection.

## Details

The first time that you will run the container, you need to apply the migrations and create a superuser in order to test the app. You can do that as follow:
- Change directory to the service directory: `cd ../django_psql_service`
- Execute: `docker compose up`
- On a different terminal session, change again directory to the service directory: `cd ../django_psql_service`
- Execute: `make migration`
- Execute: `make superuser`, and create a superuser
- Visit on your browser: `localhost:8000/admin`