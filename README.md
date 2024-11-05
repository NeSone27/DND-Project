# DND Service API

## Installation & Setup

### Prerequisites

- Golang 1.16+
- PostgreSQL

### Steps

1. Install dependencies
   `go mod tidy`

2. Configure PostgreSQL and create tables
   `psql -U postgres -d DND_db -f db/schema.sql`

3. Run the project
   `go run main.go`

## API Documentation

### Base URL

http://localhost:8080

### Endpoints

- User
   - POST /user
     Request Body: json
     ```json
     {
        "username": "string",
        "password": "string",
        "role": "string", // admin, user
        "status": "string" // active, inactive
     }
     ```
   - GET /user
   - GET /user/{id}
   - PATCH /user/{id}
     Request Body: json
     ```json
     {
        "password": "string",
        "role": "string", // admin, user
        "status": "string" // active, inactive
     }
   - DELETE /user/{id}

- Class
   - POST /class
     Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - GET /class
   - GET /class/{id}
   - PATCH /class/{id}
         Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - DELETE /class/{id}

- Race
   - POST /race
     Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - GET /race
   - GET /race/{id}
   - PATCH /race/{id}
         Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - DELETE /race/{id}

- Difficulty Level
   - POST /difficulty-level
      Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - GET /difficulty-level
   - GET /difficulty-level/{id}
   - PATCH /difficulty-level/{id}
      Request Body: json
     ```json
     {
        "name": "string",
        "detail": "string",
        "status": "string", // active, inactive
        "user_dnd_id": "int" // id user dnd if use authenticated can be null
     }
     ```
   - DELETE /difficulty-level/{id}

- Quest
   - POST /quest
     Request Body: json
     ```json
     {
        "title": "string",
        "description": "string",
        "difficulty_level_id": "int",
        "created_by": "int",
        "status": "string", // active, inactive
        "is_public": "boolean",
        "image": ["string"] // example: ["http://link-public-image1.com", "http://link-public-image2.com"]
     }
     ```
   - GET /quest
     Query: params
     ```
     user_dnd_id: int // can be null for public quest
     ```
   - PATCH /quest/{id}
      Request Body: json
     ```json
     {
        "title": "string",
        "description": "string",
        "difficulty_level_id": "int",
        "status": "string", // active, inactive
        "is_public": "boolean",
        "updated_by": "int", // only user creator or admin can update
        "image": ["string"] // example: ["http://link-public-image1.com", "http://link-public-image2.com"]
     }
     ```
   - DELETE /quest/{id}
       Request Body: json
     ```json
     {
        "user_dnd_id": "int" // only user creator or admin can delete
     }
     ```

- Character
   - POST /character
     Request Body: json
     ```json
     {
        "title": "string",
        "description": "string",
        "class_id": "int",
        "race_id": "int",
        "status": "string", // active, inactive
        "is_public": "boolean",
        "created_by": "int",
        "image": ["string"] // example: ["http://link-public-image1.com", "http://link-public-image2.com"]
     }
     ```
   - GET /character
     Query: params
     ```
     user_dnd_id: int // can be null for public quest
     ```
   - PATCH /character/{id}
      Request Body: json
     ```json
     {
        "title": "string",
        "description": "string",
        "class_id": "int",
        "race_id": "int",
        "status": "string", // active, inactive
        "is_public": "boolean",
        "updated_by": "int", // only user creator or admin can update
        "image": ["string"] // example: ["http://link-public-image1.com", "http://link-public-image2.com"]
     }
     ```
   - DELETE /character/{id}
       Request Body: json
     ```json
     {
        "user_dnd_id": "int" // only user creator or admin can delete
     }
     ```
