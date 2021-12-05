### DOCKER SETUP
This project provides sample docker container configurations for various services.

### SERVICES
- Python service: `python_service`
    - The container contains a python virtual environment with all the dependencies described inside the `Pipfile`.
### TEST THE CONTAINERS

1. Docker desktop should be installed and running on your machine.
2. Clone the repository.
2. Open your terminal and change directory inside a service folder, e.g. : `cd .../python_service`
3. Execute on your terminal: `docker compose up`
    - To enter inside a specific container, execute on your terminal: `docker compose run <<service_name>> bash`
