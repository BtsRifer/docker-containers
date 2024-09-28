### SERVICE RUN
1. Execute on terminal: `docker compose up`

The following logs should appear indicated that the container is properly setup:

```bash
python_service-1  | ==> ALL DONE WITH SETTING UP THE VIRTUAL ENVIRONMENT
python_service-1  | ==> LOCATION: /app/python_service/.venv/bin/python
python_service-1  | ==> VERSION: 3.12.6 (main, Sep 27 2024, 06:08:29) [GCC 10.2.1 20210110]
```

* To have a look into the virtual environment setup execute:
    * Enter the container: `docker compose run python_service bash`
    * Activate the virtual environment inside the container: `source .venv/bin/activate`
    * Take a look into the dependencies: `pip freeze`