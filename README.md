# User and File Service

A Go-based REST API for managing users and their files with secure JWT authentication.

## Functionality

### User Management

- **Register** (`POST /public/api/users`): Create a new user. Returns **access** and **refresh tokens** for authentication.
- **Login** (`POST /public/api/users/login`): Authenticate and receive tokens.
- **Get All Users** (`GET /public/api/users`): Publicly available list of all users.
- **Update User** (`PUT /private/api/users`): Update user name and email. *(Requires Authorization)*
- **Delete User** (`DELETE /private/api/users/{id}`): Delete a user by ID. *(Requires Authorization)*

### File Management

- **Upload File** (`POST /public/api/files/{id}`): Upload a file for a user ID.
- **Download File** (`GET /public/api/files/{id}`): Download a file by its ID.
- **Get User's Files** (`GET /public/api/files/user/{id}`): List all files for a user.
- **Delete User's Files** (`DELETE /public/api/files/user/{id}`): Delete all files for a user.

## Routes

- **Public Routes:**
  - All `/public/` endpoints.
  - Do **not** require Authorization.
  - Includes registration, login, user listing, file upload/download/list/delete.

- **Protected Routes:**
  - All `/private/` endpoints.
  - Require `Authorization` header:
    ```
    Authorization: Bearer <token>
    ```

## Data Storage

- **MySQL:** Stores user data.
- **MongoDB:** Stores file metadata and contents (GridFS).
- **RabbitMQ:** Handles background events for file processing.

## How to Run

1. Clone the repository.
2. From the project root, run:
    ```bash
    docker-compose up --build -d
    ```
3. View logs to see the consumer handling RabbitMQ events:
    ```bash
    docker-compose logs -f consumer
    ```
4. Access the API documentation:
    - **Swagger UI:** [http://localhost:8081/swagger/index.html](http://localhost:8081/swagger/index.html)

## Architecture

- **Clean Architecture:** Separation of handlers, use-cases, repositories.
- **Context-aware:** Supports request timeouts and cancellation.
- **MongoDB GridFS:** Efficient storage and retrieval for large files.
