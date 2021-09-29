# Example Queries for GraphQL

## Create User 1

```graphql
mutation createUser {
  createUser(input: { firstName: "Dale", lastName: "Chang", goesBy: "Dale", middleName: "", email: "dale@eloquia.io", gender: "M", ethnicity: "East Asian", position: "Software Engineer", institution: "Cigna", isActive: true}) {
    firstName
    lastName
  }
}
```

## Get All Users

```graphql
query findUsers {
  users {
    firstName
    lastName
    email
  }
}
```

## Create User 2

```graphql
mutation createUser {
  createUser(input: { firstName: "John", lastName: "Doe", goesBy: "Jack", middleName: "Bradly", email: "john.doe@gmail.com", gender: "M", ethnicity: "Caucasian", position: "Undergraduate Student", institution: "Made Up University", isActive: true}) {
    firstName
    lastName
  }
}
```

## Create Project 1

```graphql
mutation createProject {
  createProject(input: { name: "Test Project 1", description: "description for test project 1" } ) {
    name
  }
}
```

## Get All Projects

```graphql
query findProjects {
  projects {
    name
    description
  }
}
```

## Create Project 2

```graphql
mutation createProject {
  createProject(input: { name: "Test Project 2", description: "" } ) {
    name
  }
}
```

## Update Project 2

```graphql
mutation updateProject {
  updateProject(input: { id: "", name: "", description: "" }) {
    id
    name
    description
  }
}
```
