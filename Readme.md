<h1 align="center">Samespace Backend Task🧭</h1>

## 📚 | Problem Statement

-Develop a TODO API using Golang and ScyllaDB that supports basic CRUD operations and
includes pagination functionality for the list endpoint.
Requirements
● Set up a Golang project and integrate ScyllaDB as the database for storing TODO items.
Ensure that items in the database are stored user-wise.
● Implement endpoints for creating, reading, updating, and deleting TODO items for a
single user at a time. Each TODO item should have at least the following properties: id,
user_id, title, description, status, created, updated.
● Implement a paginated list endpoint to retrieve TODO items.
● Provide support for filtering based on TODO item status (e.g., pending, completed).


## 🎯 | Sample Queries

### POST /v1/todo

Create a new todo item.

-   **URL:** `http://localhost:8080/v1/todo`
-   **Method:** POST
-   **Body:**
    
    jsonCopy code
    
    `{
      "id": "123e4567-e89b-12d3-a456-426614174001",
      "user_id": "456e7890-e12c-34f5-b678-536724285000",
      "title": "Complete project tasks",
      "description": "Finish coding the backend and write documentation",
      "status": "completed",
      "created": "2024-02-05T12:00:00Z",
      "updated": "2024-02-05T12:00:00Z"
    }` 
    

### GET /v1/todo

Retrieve todo items based on filters.

-   **URL:** `http://localhost:8080/v1/todo`
-   **Method:** GET
-   **Query Parameters:**
    -   `status`: Filter by status (`pending` or `completed`)
    -   `size`: Number of items to retrieve (default 10)
    -   `offset`: Offset for pagination (default 0)

### PUT /v1/todo/:id

Update a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** PUT
-   **URL Parameters:**
    -   `id`: ID of the todo item to update

### GET /v1/todo/:id

Retrieve details of a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** GET
-   **URL Parameters:**
    -   `id`: ID of the todo item to retrieve

### DELETE /v1/todo/:id

Delete a specific todo item.

-   **URL:** `http://localhost:8080/v1/todo/:id`
-   **Method:** DELETE
-   **URL Parameters:**
    -   `id`: ID of the todo item to delete



## 🌐 | Test Project

- Clone this repository.
- Install Docker Desktop.
- Run `docker-compose build` in the root directory of the project. This process takes a a minute or two to complete, for the first time.
- Run `docker-compose up ` to initialize a Scylla-DB instance running on port 9042.
- Go from the main directory to `\cmd\todo-api` and then run `go run main.go` to start the service.

<br/>

### Features Implemented:

- Implemented CRUD routes for interaction b/w server and ScyllaDB.
- The API's are paginated for easy data retrieval.
- The application's DB part is Dockerized and is stateful through volumes.
- Support for filtering based on TODO item status (e.g., pending, completed).

### Current Architecture:

- Containerized approach to solving the problem statement.
- Two Interfaces one for the server and one for db are interacting b/w each other for the backend application.

### Future Scope:

- The current architecture is a very basic implementation of the problem statement.
- Depending upon the scale, the entire architecture can be **scaled horizontally** using nginx load balancing.
- Web can use a queueing mechanism like Rabbit or BullMQ to introduce pub-sub architecture to improve performance.
- The Go-Server could be containerized to improve deployment.
- Introduction to goroutines would increase the overall throughput of the service.  

## 🧑🏽 | Author

**Aditya Sinha**

- Linkedin: [@Aditya Sinha](https://www.linkedin.com/in/aditya-s-a07a54121/)
- Github: [@BREACH1247](https://github.com/BREACH1247)

<br/>

---