# place-api

Simplify version of [wdve-td-api](https://github.com/thestrayed/wdve-td-api/). But this is written in GO instead of typescript.

## Getting Started

### Prerequisite

1. Make sure you have `docker` and `docker-compose` installed.

1. Clone project.

1. Run following command at the root of project in order to get both application and database running

```bash
docker-compose up
```

### Running application without `docker`

Please be remind that application is coupled with `POSTGRES`, the application won't start unless there is `POSTGRES` instance running.

```bash
go run main.go
```

### Test

```bash
go test ./...
```

## Improvements

1. Write test (integration and unit) for all functionality

1. Use migration script rather than auto migrate from `DB` itself

1. Separate route for each domain e.g. `place` route and so on

1. Utility function to validate request

## References

- https://medium.com/@adigunhammedolalekan/build-and-deploy-a-secure-rest-api-with-go-postgresql-jwt-and-gorm-6fadf3da505b

- https://github.com/gin-gonic/gin
