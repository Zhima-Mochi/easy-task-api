# easy-task-api - Domain Driven Design Example

## Introduction
This is a simple example of a Domain Driven Design (DDD) project. The project is a restful task API application.

## Project Structure 
```
easy-task-api
├── app
│   ├── assembler
│   │   └── task_assembler.go
│   ├── dto
│   │   └── task_dto.go
│   ├── mocks
│   │   └── mock_task_app_service.go
│   ├── task_app_service.go
│   └── task_app_service_test.go
├── domain
│   ├── entity
│   │   └── task.go
│   ├── repo
│   │   ├── mocks
│   │   │   └── mock_task_repository.go
│   │   ├── task_repository.go
│   │   └── task_repository_test.go
│   ├── service
│   │   ├── mocks
│   │   │   └── mock_task_service.go
│   │   ├── task_service.go
│   │   └── task_service_test.go
│   └── valueobject
│       └── status.go
├── infra
│   └── persistence
│       ├── persistence.go
│       └── persistence_test.go
├── main.go
└── middleware
    └── middleware.go

```

## Run
```bash
make docker-build
make docker-run
```

## Endpoints
- GET /tasks
- POST /tasks
- GET /tasks/{id}
- PUT /tasks/{id}
- DELETE /tasks/{id}

### Swagger
- GET /swagger/index.html

