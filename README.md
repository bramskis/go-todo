# go-todo

Simple todo backend service written in Go. Connects to a Postgresql database

## Endpoints
- `GET /todo` - Gets all Todo items
- `GET /todo/:id` - Gets Todo item associated with `:id`
- `DELETE /todo/:id` - Deletes Todo item associate with `:id`
- `POST /todo` - Create new Todo item
```
Request Body:
{
  "Title":  string
  "Description" string
  "Deadline": string
  "Completed": bool
}
```
- `PUT /todo` - Update Todo item
```
Request Body:
{
  "Id": string
  "Title":  string
  "Description" string
  "Deadline": string
  "Completed": bool
}
```
