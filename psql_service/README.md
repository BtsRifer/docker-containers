## PostgreSQL service
A container with a PostgreSQL database connection. 

## DETAILS
The database is initialized based on the `./db/init.sh` file (optional).
To connect to the database:
- Change directory to the service directory and execute: `docker compose up`
- On a different terminal session, change directory to the service directory and execute: `docker compose exec db psql -U postgres`
    - Or execute queries outside the container as follow, e.g.: `docker compose exec db psql -U postgres -d test_db -c "select * from hello_world;"`