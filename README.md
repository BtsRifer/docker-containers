### DOCKER SETUP
This project provides sample docker container configurations for various services.

### SERVICES
- Python service: `python_service`
    - A container with a python virtual environment with all the dependencies listed inside the `Pipfile`.
- Golang service: `go_service`
    - A container with a Go installation with all the dependencies listed inside the `go.mod` file.
- Django REST with PostgreSQL service: `django_psql_service`
    - A container with a Django application setup with a PostgreSQL database connection.
    - Change directory to the service directory, execute `docker compose up` and visit `localhost:8000/admin` (user: admin, pass: admin) to check the application.
### TEST THE CONTAINERS

1. Docker desktop should be installed and running on your machine.
2. Clone the repository.
2. Open your terminal and change directory inside a service folder, e.g. : `cd .../python_service`
3. Execute on your terminal: `docker compose up`
    - To enter inside a specific container, execute on your terminal: `docker compose run <<service_name>> bash`
