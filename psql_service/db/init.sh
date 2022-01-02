# Create databse if not exists.
psql -U postgres -tc "SELECT 1 FROM pg_database WHERE datname = 'test_db'" | grep -q 1 || psql -U postgres -c "CREATE DATABASE test_db"

# Create schema.
psql -U postgres -d test_db -a -f opt/schema_queries/create_schema.sql

# Insert.
psql -U postgres -d test_db -a -f opt/schema_queries/insert.sql
