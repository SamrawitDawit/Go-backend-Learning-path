# API Documentation

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
## Get Tasks
This endpoint retrieves a list of tasks.

### Request
The request should be sent via an HTTP GET method to http://localhost:8080/tasks.

### Response
Upon a successful execution, the endpoint returns a status code of 200 and a JSON response with an array of tasks. Each task object in the array contains the following properties:
  - id (string): The ID of the task.
  - title (string): The title of the task.
  - description (string): The description of the task.
  - due_date (string): The due date of the task.
  - status (string): The status of the task.

### Possible Errors
  - 500 Internal Server Error: If there is an issue on the server side.

## Get Task by ID
This endpoint retrieves a task with a specific id.

### Request 
The request should be sent via an HTTP GET method to http://localhost:8080/tasks/:id.

### Response 
Upon a successful execution, the endpoint returns a status code of 200 and a JSON response with a task of that ID.
The task object contains the following properties:
  - id (string): The ID of the task.
  - title (string): The title of the task.
  - description (string): The description of the task.
  - due_date (string): The due date of the task.
  - status (string): The status of the task.

### Possible Errors
  - 404 Not Found: If the task with the specified ID is not found.
  - 500 Internal Server Error: If there is an issue on the server side.

## Add Task
This endpoint adds a new task.

### Request
The request should be sent via an HTTP POST method to http://localhost:8080/tasks.
The request payload should have the following parameters in the raw request body type:
  - id (string, optional): The ID of the task.
  - title (string, optional): The title of the task.
  - description (string, optional): The description of the task.
  - status (string, optional): The status of the task.

### Response 
Upon a successful execution, the endpoint returns a status code of 201 and a JSON response with a message "Task created".

### Possible Errors
  - 400 Bad Request: If the request body is invalid.
  - 500 Internal Server Error: If there is an issue on the server side.

## Remove Task 
This endpoint removes a task with a specific id.

### Request 
The request should be sent via an HTTP DELETE method to http://localhost:8080/tasks/:id.

### Response 
Upon a successful execution, the endpoint returns a status code of 200 and a JSON response with a message "Task removed".

### Possible Errors
  - 404 Not Found: If the task with the specified ID is not found.
  - 500 Internal Server Error: If there is an issue on the server side.

## Update Task
This endpoint updates a task of specific id.

### Request 
The request should be sent via an HTTP PUT method to http://localhost:8080/tasks/:id.
The request payload should have one or more of the following parameters in the raw request body type:
  - id (string, optional): The ID of the task.
  - title (string, optional): The title of the task.
  - description (string, optional): The description of the task.
  - status (string, optional): The status of the task.

### Response 
Upon a successful execution, the endpoint returns a status code of 200 and a JSON response with a message "Task updated".

### Possible Errors
  - 400 Bad Request: If the request body is invalid.
  - 404 Not Found: If the task with the specified ID is not found.
  - 500 Internal Server Error: If there is an issue on the server side.