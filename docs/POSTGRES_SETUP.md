# Setting Up Postgres

## Creating Database

To set up the initial database:

1. Create a database named `test_qapla`

## Initial Schema

The initial schema should be listed in the first migration file starting with `000001_*`.

## Initial Role

1. Create a user (login by default) named `qapla_user` with some kind of password
2. Run `GRANT CONNECT ON DATABASE test_qapla TO qapla_user;`
3. Run `GRANT USAGE ON SCHEMA core_qapla TO my_user;` or whatever the schema name is
4. Run `GRANT ALl PRIVILEGES ON ALL TABLES IN SCHEMA core_qapla TO qapla_user;`

## Running Migration in Dockerized Postgres via `psql`

After using `docker cp` to copy over the migration file, run `docker exec` to connect to the Postgres container and run `\i \path\TO\file_name.sql` to run the DML in the copied file.
