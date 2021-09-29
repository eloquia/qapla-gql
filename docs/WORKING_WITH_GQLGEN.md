# Working with GQLGEN

The steps for working with `gqlgen` follow the general pattern:

1. Define a `model` in `graph/model`
2. Define a schema/relationship in `graph/schema.graphqls`
3. Run `go run github.com/99designs/gqlgen generate` to generate code
4. Define function(s) in `graph/schema.resolvers.go`
5. [Optional] use dependency injection in `graph/schema.resolvers.go` with a service defined elsewhere
