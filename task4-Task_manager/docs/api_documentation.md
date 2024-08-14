# Postman Documentation

`https://documenter.getpostman.com/view/33567770/2sA3s6FpkW`


# API Documentation

### Get Tasks

  * Endpoint: `Get/tasks`
  * Description: retrieves a list of tasks.
  * Response:
    - Status Code: `200 Ok`
    - Body: a JSON response with an array of tasks.
        ```json
           {
            "id": "string",
            "title": "string",
            "description": "string",
            "due_date": "time.Time",
            "status": "string"
            }
            ```
  * Possible Errors: 
    - `500 Internal Server Error`: If there is an issue on the server side.

### Get Task By ID

  * Endpoint: `Get/task/:id`
  * Description: retrieves a task with a specific id.
  * Response:
    - Status Code: `200 Ok`
    - Body: a JSON response a JSON response with a task of that ID.
           ```json
           {
            "id": "string",
            "title": "string",
            "description": "string",
            "due_date": "time.Time",
            "status": "string"
            }
            ```       
  * Possible Errors: 
    - `404 Not Found`: If the task with the specified ID is not found.
    - `500 Internal Server Error`: If there is an issue on the server side.

### Add Task 

  * Endpoint: `POST/tasks`
  * Description: adds a new task.
  * Request Payload:
  ```json
  {
    "id": "string",
    "title": "string",
    "description": "string",
    "due_date": "time.Time",
    "status": "string"
  }
  ```
  * Response:
    - Status Code: `201 Created`
    - Body: JSON message `"Task created"`
  * Possible Errors: 
    - `400 Bad Request`: If the request payload is invalid.
    - `500 Internal Server Error`: If there is an issue on the server side.

### Remove task

  * Endpoint: `DELETE/tasks/:id`
  * Description: removes a task with a specific id.
  * Request Parameters:
  `id`: The ID of the task to be removed
  * Response:
    - Status Code: `200 OK`
    - Body: JSON message `"Task removed"`
  * Possible Errors: 
    - `404 Not Found` If the task with the specified ID is not found.
    - `500 Internal Server Error`: If there is an issue on the server side.

### Update task

  * Endpoint: `PUT/tasks/:id`
  * Description: updates a task of specific id.
  * Request Parameters:
  `id`: The ID of the task to be updated
  * Request Payload:
  ```json
  {
    "id": "string", //optional
    "title": "string", //optional
    "description": "string", //optional
    "due_date": "time.Time", //optional
    "status": "string" //optional
  }
  ```
  * Response:
    - Status Code: `200 OK`
    - Body: JSON message `"Task updated"`
  * Possible Errors: 
    - `404 Not Found` If the task with the specified ID is not found.
    - `400 Bad Request`: If the request payload is invalid.
    - `500 Internal Server Error`: If there is an issue on the server side.