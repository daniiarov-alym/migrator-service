# migrator-service

## About

Simple service designed to run your migrations in PostgreSQL database.

## Usage

Service relies on existense of folder `/migrations` containing sql script files. Their names should be valid names for migration files
(for example `0001_foo.up.sql`, `0001_foo.down.sql`). Refer to [documenation](https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md)
for additional information.

# Docker container

```
$ git clone https://github.com/daniiarov-alym/migrator-service.git
$ cd migrator-service
$ docker build . -t migrator-service
$ docker run -e PG_HOST=address -e PG_PORT=port -e PG_USER=username -e PG_PASSWORD=password -e PG_DATABASE=database -v path_to_migrations:/migrations migrator-service
```

Adjust flags and command for your needs.
