# Test

This directory contains resources for starting a conatinerized, local instance of the Streamflows server.
This enables manual validation of E2E scenarios on your local machine.

## Requirements

[Docker Compose] must be installed on the host machine and available on the CLI.

Verify that Docker Compose is installed, run:

```bash
docker-compose -v
```

## Containers

### `server`

The `server` container builds and runs the Streamflows server.

To start the container, run:

```bash
docker-compose up server
```

The Streamflows server is accessible on the host machine via `http://localhost:8080`

### `mongo`

This container hosts a local Mongo DB instance for use by the local Streamflows server.

To start the container, run:

```bash
docker-compose up -d mongo
```

### `mongo-express`

This container hosts a local [Mongo Express] instance.
This can be used as a web-based Admin UI to read/write collections in the local Mongo DB.

The Mongo Express instance is accessible via `http://localhost:8081` with the credentials `admin` and `pass`.

The `mongo-express` container depends on the `mongo` container.
To start these containers, run:

```bash
docker-compose up -d mongo mongo-express
```

[Docker Compose]: https://docs.docker.com/compose
[Mongo Express]: https://github.com/mongo-express/mongo-express
