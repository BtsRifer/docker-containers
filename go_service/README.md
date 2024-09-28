### About
A container with a Go installation, including all dependencies specified in the `go.mod` file

### Run
1. Execute in terminal: `docker compose up`

The following logs should appear, indicating that the container is properly set up:

```bash
python_service-1  | ==> ALL DONE WITH SETTING UP THE VIRTUAL ENVIRONMENT
python_service-1  | ==> LOCATION: /app/python_service/.venv/bin/python
python_service-1  | ==> VERSION: 3.12.6 ...
```

To inspect the virtual environment setup:

* Enter the container: `docker compose run python_service bash`
* Activate the virtual environment inside the container: `source .venv/bin/activate`
* View the installed dependencies: `pip freeze`