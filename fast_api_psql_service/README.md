## About

A FastAPI application setup with a PostgreSQL database container connection.

## Run

1. Navigate to the container's folder: `cd /fast_api_psql_service`.
2. Execute in terminal: `docker compose up` (or `make run`)
3. Try a sample request on root: `curl http://0.0.0.0:8000`
   
## Details

* Access the API Swagger Documentation at: `http://0.0.0.0:8000/docs`
* Create a local virtual environment: `poetry install --no-root` (or `make vevn`)
* Connect to the service's database from an SQL client using the URI: `postgresql://postgres:postgres@127.0.0.1:5432`
    * Or access the database directly from its container: `docker compose exec api_db psql -U postgres -d api`