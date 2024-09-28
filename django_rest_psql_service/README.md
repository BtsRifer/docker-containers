## About

A Django application setup with a PostgreSQL database container connection.

## Run

1. Navigate to the container's folder: `cd /django_rest_psql_service`.
2. Execute in terminal: `docker compose up`

## Details

The first time that you will run the container, you need to apply the migrations and create a superuser in order to test the app:
1. While the container is running, in a new terminal session navigate to the container's folder: `cd /django_rest_psql_service`.
2. Execute: `make migration`
3. Execute: `make superuser`, and create a superuser
4. Visit on your browser the application's admin page: `localhost:8000/admin`
5. From an SQL client you can connect to the application's database by using the following URI: `postgresql://postgres:postgres@127.0.0.1:5432`
    * Or enter connect to the database directly from the database's container: `docker compose exec api_db psql -U postgres -d api`