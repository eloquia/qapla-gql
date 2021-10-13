# Qapla-GQL

Qapla-GQL is the companion GraphQL API for Qapla-Ui. It will eventually replace the cev-api REST API.

The Qapla project is moving from REST to GraphQL because the supported data:

* Is highly circular:
  * People are part of Projects
  * Projects have people
  * Meetings can be about Projects or People
  * Meeting Notes are written by a Person and can be about a Project or a Person
* A variety of screens that all need different data shape

Moving from a conventional REST API will reduce the number of data shapes that the back-end needs to handle.

---

## Getting Started

1. Clone this project
2. Install dependencies
3. Run the server with `go run server.go`

---

## Contributing

To add code

1. Update the `Query` or `Mutation` object in `schema.graphqls` with the GraphQL query you'd like to support
2. Run `go run github.com/99designs/gqlgen generate` to generate boilerplate code for the interface
3. Update the corresponding service interface in `services`
4. Add a function that implements the business logic provided by the GraphQL query
5. Write a unit test for testing the business logic
6. Run `go run server.go` to test the newly added business logic for functionality
