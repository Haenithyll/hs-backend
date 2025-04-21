# 📁 Project Structure Overview

```
├── cmd
│   └── app
├── docs
├── internal
│   ├── config
│   ├── di
│   ├── domain
│   ├── handler
│   ├── middleware
│   ├── model
│   │   ├── enum
│   │   └── json
│   ├── repository
│   ├── request
│   │   └── validator
│   ├── response
│   │   └── mapper
│   ├── route
│   ├── service
│   ├── util
│   │   ├── filter_util
│   │   └── map_util
│   └── validation
└── reports
```

### 🔹 cmd/app

Entry point of the application. Responsible for initializing and starting the service.

### 🔹 docs

Holds documentation files related to the project, API specs, and design notes.

### 🔹 internal/config

Handles application configuration: environment variables, database connection setup, and internal configuration logic.

### 🔹 internal/di

Dependency injection layer. Instantiates and wires together services, repositories, and handlers.

### 🔹 internal/domain

Defines custom application types and shared logic, such as custom errors and standardized HTTP responses.

### 🔹 internal/handler

API handlers. Receives incoming requests, binds/validates input, and delegates to the appropriate service.

### 🔹 internal/middleware

Houses middleware components such as authentication. Extendable for other cross-cutting concerns.

### 🔹 internal/model

Represents database models. Includes:
* enum: Enumerated types.
* json: Structs for JSONB fields.

### 🔹 internal/repository

Data access layer. Provides CRUD operations for each model using the database.

### 🔹 internal/request

Defines all request input types (URI params, query params, and body).
Includes:
* validator: Custom validation logic for complex request rules.

### 🔹 internal/response

Defines response types returned by the API.
Includes:
* mapper: Maps service results to response DTOs.

### 🔹 internal/route

Groups and registers all API routes, organizing them per model. Uses DI to bind routes to handlers.

### 🔹 internal/service

Business logic layer. Processes data from repositories, performs operations, and builds results.

### 🔹 internal/util

Utility functions shared across the app, including:
* filter_util: Filtering helpers.
* map_util: Mapping helpers.

### 🔹 internal/validation

Additional validation logic for requests (e.g., ownership checks, foreign key integrity).

### 🔹 reports

Auto-generated test coverage reports and related testing metrics.