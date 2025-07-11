# did-be-wp

The Workspace contains the following projects:

- [api](https://github.com/thanhtan541/did-be-wp/api): Http api
- [core](https://github.com/thanhtan541/did-be-wp/core): Core library for did specification

> **_NOTE:_** To extend this workspace, follow the [link](https://go.dev/doc/tutorial/workspaces)

## Installation

------------

### Pre-requisites

You'll need to install:

- [Golang 1.23](https://go.dev/doc/install): main language

### Init local db

```bash
# To init db and run migration
./scripts/init-db.sh

# or only run migration
SKIP_DOCKER=true ./scripts/init_db.sh
```

To add new migration
```
migrate create -ext sql -dir db/migrations -seq create_users_table
```

> **_NOTE:_** All of migrations are located in `api/migrations` directory

### Build the project

To build the project, run:

```bash
go build -o api ./api
```

### Running tests

You can run tests using the following command:

```bash
make test
```

### Control log output
```bash
# Need to create logs/app.log before running the project
export LOG_TO_FILE=true // log to file
export LOG_TO_FILE=false // log to stdout
```

### Generate code reports

You can run code reports using the following command:

```bash
./scripts/init_code_report_cov.sh
```

Send reports to sonarqube local
- [Check this](./sonarqube-local/README.md)


### F.A.Qs
1. Fix `too many open files` error on MacOS
```bash
ulimit -n X (X is the number of open files)
```
