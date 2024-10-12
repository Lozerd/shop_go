# SHOP GO

## Project goals:
- [ ] Learn Golang
- [ ] Learn Kafka
- [ ] Learn DDD
- [ ] Practice and gain new skills

# Project's main dependencies
- [GORM](https://gorm.io/) - Golang ORM
- [PGX](https://github.com/jackc/pgx) - Golang PostgreSQL driver, used by GORM
- [GIN](https://gin-gonic.com/) - Gin Web Framework
- [KAFKA-GO](https://github.com/segmentio/kafka-go) - Golang implementation of [Apache Kafka](https://kafka.apache.org/)

# Requirements
- [Go](https://go.dev/) `>=1.23.2` (minimum); [Install instructions](https://go.dev/doc/install)
- [AIR](https://github.com/air-verse/air) - Live reload for Go apps

# Development

### Via `go install` (Recommended)

With go 1.23 or higher:

```bash
go install github.com/air-verse/air@latest
```

### Via install.sh

```shell
# binary will be $(go env GOPATH)/bin/air
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# or install it into ./bin/
curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s

air -v
```

## Clone the project
```console
git clone https://github.com/Lozerd/shop_go
```

## Install dependencies
```console
cd shop_go
go mod download
```

## Install go-swagger3
```console
go install github.com/parvez3019/go-swagger3@latest
```

## Generate swagger docs
```console
go-swagger3 --module-path . --main-file-path ./cmd/server/main.go --output oas.json --schema-without-pkg --generate-yaml true
```

## Run application using air
```console
air
```

## Run application using docker compose
```console
docker compose up --build
```

# License
[The MIT License (MIT)](License)
