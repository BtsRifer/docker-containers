### DOCKER SETUP
This project provides sample docker container configurations for various services.

### TEST THE CONTAINERS
Any of the service containers can be tested as follow:

1. `Docker Desktop` should be installed and running on your machine (or `Docker` and `Docker Compose`).
2. Clone the repository.
3. Open your terminal and change directory inside a service folder.
4. Execute on your terminal: `docker compose up`
    - To enter inside a specific container, execute on your terminal: `docker compose run <<service_name>> bash`

### SERVICES DETAILS
1. Python service: `python_service`
    - A container with a python virtual environment with all the dependencies listed inside the `Pipfile`.
2. Golang service: `go_service`
    - A container with a Go installation with all the dependencies listed inside the `go.mod` file.
3. PostgreSQL service: `psql_service`
    - A container with a PostgreSQL database connection. The database is initialized based on the `./db/init.sh` file (optional).
4. Django REST with PostgreSQL service: `django_psql_service`
    - A container with a Django application setup with a PostgreSQL database connection.