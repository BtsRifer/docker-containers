## About
A container with a PostgreSQL database connection. The database is initialized using the `./db/init.sh` file.

## Run
1. Execute in terminal: `docker compose up`
2. Connect to the database: `docker compose exec db psql -U postgres -d test_db`
    * Or execute queries from outside the container (while the container is running): `docker compose exec db psql -U postgres -d test_db -c "select * from hello_world;"`
    * Or connect to the database from a client using the URI: `postgresql://postgres:postgres@127.0.0.1:5432`