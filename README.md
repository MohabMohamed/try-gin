# try-gin
A todo website just for testing golang web framework gin

## installation

### sqlite

Debian based(Ubuntu, Debian etc.) :

  - `$ sudo apt-get update `

  - `$ sudo apt-get install sqlite3 ` 
RPM based(RHEL, CentOS, Fedora etc.) :

  - `$ sudo yum update` 

  - `$ sudo yum install sqlite` 

### try-gin

  - `$ cp .env.example .env`
  
  - update  the .env file with your DB url.

  - `$ go get`

  - `$ pip install pre-commit`

  - `$ pre-commit install`


## Run
  - `$ make dev`
## Testing
  - `$ make test`

## Routes

### `GET /todos`

Gets all the todos. 
Response:
```js
[{
  id: int,
  title: String,
  description: String,
  done: Boolean
}]
```

### `POST /todo`

Creates a new todo with default state (done : false). If the todo is sent without title or description the endpoint will return error message with status code `400 bad request`.
Request:
```js
{
  title: String,
  description: String,
}
```

Response:
```js
{
  id: int,
  title: String,
  description: String,
  done: false
}
```

### `GET /todo/:id`

Gets a specific todo by it's `id` and returns `404 not found` if the todo doesn't exist.
Response:
```js
{
  id: int,
  title: String,
  description: String,
  done: Boolean
}
```

### `PATCH /todo/:id`

Update specific todo `Done` status by `Id` and returns `404 not found` if the todo doesn't exist.

Request:
```js
{
  done: Boolean
}
```

Response:
```js
{
  id: int,
  title: String,
  description: String,
  done: Boolean
}
```

### `DELETE /todo/:id`

Delete specific todo by `id` and returns the deleted todo with Status code `204 No Content` or `404 not found` if the todo doesn't exist.

Response:
```js
{
  id: int,
  title: String,
  description: String,
  done: Boolean
}
