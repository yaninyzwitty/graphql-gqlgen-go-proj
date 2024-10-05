# GraphQL Todo Application

This is a simple GraphQL Todo application built using Go, gqlgen, and PostgreSQL. The application allows users to create and manage todo items, demonstrating the use of GraphQL for handling mutations and queries.

## Features

- Create todos with associated users
- Retrieve a list of todos
- User validation through foreign key constraints

## Technologies Used

- Go (Golang)
- gqlgen (GraphQL server for Go)
- PostgreSQL (Database)
- UUID (Unique identifiers for users and todos)

## Project Structure
├── graph │ ├── model │ ├── resolver │ ├── schema.graphql │ └── resolvers.go ├── database │ └── db.go ├── migrations │ └── <migration files> ├── Dockerfile ├── docker-compose.yml ├── go.mod ├── go.sum └── main.go

## Getting Started

### Prerequisites

- Go 1.22 or later
- PostgreSQL
- Docker
- Docker compose

### Setup

1. **Clone the repository:**
  

   ```bash
   git clone https://github.com/yaninyzwitty/graphql-gqlgen-go-proj.git
   graphql-gqlgen-go-proj
2. **iN THE .ENV**
  ```bash
    PORT=YOUR_PORT
    DATABASE_URL
