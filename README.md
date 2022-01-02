### DOCKER SETUP
This project provides sample docker container configurations for various services.

### SERVICES
- Python service: `python_service`
    - A container with a python virtual environment with all the dependencies listed inside the `Pipfile`.
- Golang service: `go_service`
    - A container with a Go installation with all the dependencies listed inside the `go.mod` file.
- PostgreSQL service: `psql_service`
    - A container with a PostgreSQL database connection. The database is initialized based on the `db/init.sh` file (optional).
    - To connect to the database:
        - Change directory to the service directory and execute: `docker compose up`
        - On a different terminal session, change directory to the service directory and execute: `docker compose exec db psql -U postgres`
            - Or execute queries outside the container as follow, e.g.: `docker compose exec db psql -U postgres -d test_db -c "select * from hello_world;"`
- Django REST with PostgreSQL service: `django_psql_service`
    - A container with a Django application setup with a PostgreSQL database connection.
    - The first time that you will run the container, you need to apply the migrations and create a superuser in order to test the app. You can do that as follow:
        - Change directory to the service directory: `cd ../django_psql_service`
        - Execute: `docker compose up`
        - On a different terminal session, change again directory to the service directory: `cd ../django_psql_service`
        - Execute: `make migration`
        - Execute: `make superuser`, and create a superuser
        - Visit on your browser: `localhost:8000/admin`
        
### TEST THE CONTAINERS

1. Docker desktop should be installed and running on your machine.
2. Clone the repository.
2. Open your terminal and change directory inside a service folder, e.g. : `cd .../python_service`
3. Execute on your terminal: `docker compose up`
    - To enter inside a specific container, execute on your terminal: `docker compose run <<service_name>> bash`
