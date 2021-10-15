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

---

## User Details

```graphql
query findUserListItems {
  userDetails {
    id
    firstName
    lastName
    assignedProjects {
      id
      name
    }
  }
}
```

---

## Create Personnel

```graphql
mutation createPersonnel {
  createPersonnel(input: { firstName: "John", lastName: "Doe", goesBy: "Jack", middleName: "Bradly", email: "john.doe@gmail.com", gender: "M", ethnicity: "Caucasian", position: "Undergraduate Student", institution: "Made Up University", isActive: true}) {
    id
    firstName
    lastName
  }
}
```

---

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

---

## Sign In

```graphql
mutation signIn {
  signIn(email: "dale@eloquia.io", password: "ehh") {
    id
    firstName
    lastName
  }
}
```

---

## Create Meeting

```graphql
mutation createMeeting {
  createMeeting(input: {name: "Test Meeting", durationMinutes: 30, startTime: "2021-10-13T17:30:15+05:30"}) {
    id
  }
}
```

## Get Meeting by ID

```graphql
query meetingById {
  getMeetingById(id: "5577006791947779410") {
    id
    name
    startTime
    durationMinutes
  }
}
```
