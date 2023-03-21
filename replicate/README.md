# How to setup read repication with PostgreSQL.

## Setup Requirements

- docker
- docker-compose

## Usage

With docker and docker composed installed run the following command:

```
$ docker-compose up
```

### Adding a read user to the replica

1. Access the master container

```
$ docker-compose exec postgresql-master bash
```

2. Access the psql client

```
$ psql -U postgres
```

3. Run the SQL commands, creating the user `read_user` and giving them permisisons

```
CREATE USER read_user WITH PASSWORD 'reader_password';
GRANT CONNECT ON DATABASE my_database TO read_user;
\connect my_database
GRANT SELECT ON ALL TABLES IN SCHEMA public TO read_user;
GRANT SELECT ON ALL SEQUENCES IN SCHEMA public TO read_user;
GRANT USAGE ON SCHEMA public TO read_user;
ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT ON TABLES TO read_user;
```

Restart all of the containers and possible connections.

## Acknowledgements

Based off of the work by [JosimarCamargo](https://gist.github.com/JosimarCamargo/40f8636563c6e9ececf603e94c3affa7)

