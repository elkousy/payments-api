# Payments API
[![wercker status](https://app.wercker.com/status/fe77fc3db8a9c8cf04e858b791b65a72/s/master "wercker status")](https://app.wercker.com/project/byKey/fe77fc3db8a9c8cf04e858b791b65a72)
[![Coverage Status](https://coveralls.io/repos/github/elkousy/payments-api/badge.svg?branch=master)](https://coveralls.io/github/elkousy/payments-api?branch=master)

A home exercise requested by Form3. The purpose is to build a REST API with CRUD operations on a the `payments` resource. The representation schema is defined [here](http://mockbin.org/bin/41ca3269-d8c4-4063-9fd5-f306814ff03f)

Below are core features of the implementation:

* The API is built using Golang and various frameworks to do not duplicate efforts and make the API simple and pragmatic.
* The structure of code and implementation is optimized for Correctness, TDD (interfaces), Readibility and Extensibility.
* The API exposes Prometheus metrics for observability purpose.
* `Postgres` is the database used to save payments
* `newman` is used to perform integration tests

## Running the API

### Prerequisites

Prior running the API, make sure you have GO and Docker installed in your machine.

To able to perform integration tests, make sure to install `newman`

### Terminal

1. `config.toml` is used to provide container-like environment variables. The application will load your custom settings when you export the variable `ENVIRONMENT=dev` on your terminal:

```
export ENVIRONMENT=dev
```

2. The dependencies are managed using `dep`. Make sure to load all necessary to your `$GOPATH`:

```
make install
```

3. The application uses a `Postgres` database, make sure to have a running db instance (e.g. `docker-compose up db`) with the right connection parameters updated in `config.toml` before running t:

```
make run
```

### Docker

Please note the the `Dockerfile` uses the binary built locally. `Docker-compose` defines services for the Payments API and Postgres:

```
make docker-compose-build
make docker-compose-up
```

You can interact with the API on port `:8080`. The observabilty metrics and the heath checks are exposed on port `:8081`

## Testing

Unit tests and code coverage could be performaed using these commands:

```
make unit-test
make test-cover
```

> Please note that any changes to interfaces need to regenerate new mocks using the command `make update-mocks`.

Run the integration tests using:

```
make integration-tests
```
