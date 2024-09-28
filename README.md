## About

This project offers sample `Docker container` configurations for various services.

## Service Details

1. **Python Service**: `python_service`
   - A container with a Python virtual environment
   - Includes all dependencies listed in the `Pipfile`

2. **Golang Service**: `go_service`
   - A container with a Go installation
   - Includes all dependencies specified in the `go.mod` file

3. **PostgreSQL Service**: `psql_service`
   - A container with a PostgreSQL database connection
   - The database is initialized based on the `./db/init.sh` file

4. **Django REST with PostgreSQL Service**: `django_psql_service`
   - A container with a Django application setup
   - Includes a PostgreSQL database connection

## Run

To test any service container:

1. Ensure `Docker Desktop` is installed and running (or `docker` and `docker-compose`).
2. Clone the repository.
3. Navigate to the desired service folder in your terminal, e.g.: `cd /python_service`.
4. Run: `docker compose up`
    * To access a specific container's shell, use: `docker compose run <<service_name>> bash`