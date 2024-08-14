## Overview
This API provides endpoints to manage tasks, including retrieving, adding, updating and deleting tasks. The tasks are now stored in a MongoDB database, and this documentaion provides details on how to configure and connect to MongoDB.

## MongoDB Configuration
To connect the GO application to MongoDB, you need to congigure the MongoDB connection string. This can be done through environment variables or configuration file.

### Environment Variables
Set the following environment variables
```export MONGODB_URI="mongodb://localhost:27017"
   export MONGODB_DATABASE="taskdb"
```

### Configuration File
Alternatively, you can use a configuration file to store the MonoDB connection settings.
eg. `config.json`
```{
  "mongodb_uri": "mongodb://localhost:27017",
  "mongodb_database": "taskmanager"
}
```
## MongoDB Connection Setup
The following steps describe how to establish a connection to MongoDB Go driver.

### Step 1: Install the MongoDB Go Driver
`go get go.mongodb.org/mongo-driver/mongo`

### Step 2: Connect to MongoDB
```go
clientOptions := options.Client().ApplyURI(mongodb://localhost:27017)
client, err := mongo.Connect(context.TODO(), clientOptions)
if err != nil{
  log.Fatal(err)
}
db := client.Database("taskdb")
tasksCollection := db.Collection("tasks")
userCollection := db.Collection("users")
```

### Step 3: Handle Database Operations 
For CRUD operations, use the `taskCollection` to perform operations on the task collection.

## Database Structure
Tasks are stored in the `tasks` collection within the specified MongoDB database. Each task document has the following structure
```json
{
  "_id": "<ObjectId>",
  "title": "Task title",
  "description": "Task description",
  "due_date": "YYYY-MM-DD",
  "status": "Not Started/Completed/etc."
}
```

# Postman Documentation

https://documenter.getpostman.com/view/33567770/2sA3s6FpkW


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