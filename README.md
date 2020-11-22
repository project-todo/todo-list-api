[![Go Report Card](https://goreportcard.com/badge/github.com/project-todo/todo-list-api)](https://goreportcard.com/report/github.com/project-todo/todo-list-api)

# todo-list-api
CRUD API for TODO lists

## TODO List
A TODO list is a collection of tasks which can be in the state of completed or not completed.

### REST Interface

Get all avilable tasks
```http
GET /tasks
```

Get specific task by ID
```http
GET /tasks/{ID}
```

Create new task
```http
POST /tasks
{
	"description": "Get milk for coffe"
}
```

Update existing task
```http
PUT /tasks/{ID}
{
	"completed": true,
	"description": "Updated description"
}
```

Delete existing task
```http
DELETE /tasks/{ID}
```

### Data Structures

**Tasks**
```json
{
	"completed": false,
	"description": "Foo bar"
}
```
