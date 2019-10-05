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
	"description": "Foo bar",
	"expiry": "1970-01-01T00:00:00+01:00"
}
```

Update existing task
```http
PUT /tasks/{ID}
{
	"completed": true,
	"title": "Updated title",
	"description": "Updated description",
	"expiry": "2019-01-01T00:00:00+01:00"
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
	"id": "UUID...",
	"completed": false,
	"title": "Get milk for coffe",
	"description": "Foo bar",
	"created": "1970-01-01T00:00:00+01:00",
	"updated": "1970-01-01T00:00:00+01:00",
	"expiry": "1970-01-01T00:00:00+01:00"
}
```
