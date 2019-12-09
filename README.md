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
	"title": "Get milk for coffe",
	"description": "Some of my guests don't like black coffe.",
	"expires": "1970-01-01T00:00:00+01:00"
}
```

Update existing task
```http
PUT /tasks/{ID}
{
	"completed": true,
	"title": "Updated title",
	"description": "Updated description",
	"expires": "2019-01-01T00:00:00+01:00"
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
	"title": "Get milk for coffe",
	"description": "Foo bar",
	"created": "1970-01-01T00:00:00+01:00",
	"updated": "1970-01-01T00:00:00+01:00",
	"expires": "1970-01-01T00:00:00+01:00"
}
```
