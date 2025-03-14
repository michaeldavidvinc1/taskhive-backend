# Api Spec

## Auth Api

***
> Login (POST)
***

```code
api/v1/auth/login
```

### Request

```json
{
  "email": "michael@gmail.com",
  "password" : "12345678" 
}
```

### Response

```json
{
  "success": true,
  "message": "Successfully login!",
  "data": {
    "email": "michael@gmail.com",
    "name": "Michael David Vinc",
    "token": "TOKEN JWT"
  }
}
```

***
> Register (POST)
***

```code
api/v1/auth/register
```

### Request

```json
{
  "name": "string name",
  "position": "string position",
  "email": "michael@gmail.com",
  "password" : "12345678" 
}
```

### Response

```json
{
  "success": true,
  "message": "Successfully login!",
  "data": {
    "email": "michael@gmail.com",
    "name": "Michael David Vinc",
    "token": "TOKEN JWT"
  }
}
```

## Task Api

***
> Create Task (POST)
***

```code
api/v1/task/create
```

### Role
- admin

### Request

```json
{
  "title": "Title task",
  "date" : "date now" ,
  "priority" : "string enum",
  "stage": "default todo",
  "teams": [
    "userId_1",
    "userId_2"
  ]
}
```

### Response

```json
{
  "success": true,
  "message": "Successfully create task!",
  "data": {
    "title": "Title task",
    "date" : "date now" ,
    "priority" : "string enum",
    "stage": "default todo",
    "teams": [
      "userId_1",
      "userId_2"
    ],
    "isDelete": false
  }
}
```

***
> Get All Task (GET)
***

```code
api/v1/task
```

### Params

- priority
- stage
- limit


### Response

```json
{
  "success": true,
  "message": "Successfully get all task!",
  "data": [
    {
      "title": "Title task",
      "date" : "date now" ,
      "priority" : "string enum",
      "stage": "default todo",
      "teams": [
        "userId_1",
        "userId_2"
      ],
      "isDelete": false
    }
  ]
}
```

***
> Get Single Task (GET)
***

```code
api/v1/task/:id
```

### Response

```json
{
  "success": true,
  "message": "Successfully get single task!",
  "data": {
    "title": "Title task",
    "date" : "date now" ,
    "priority" : "string enum",
    "stage": "default todo",
    "teams": [
      "userId_1",
      "userId_2"
    ],
    "isDelete": false
  }
}
```

***
> Update Task (PUT)
***

```code
api/v1/task/:id
```

### Role
- admin

### Request

```json
{
  "title": "Title task",
  "date" : "date now" ,
  "priority" : "string enum",
  "stage": "default todo",
  "teams": [
    "userId_1",
    "userId_2"
  ]
}
```

### Response

```json
{
  "success": true,
  "message": "Successfully update task!",
  "data": {
    "title": "Title task",
    "date" : "date now" ,
    "priority" : "string enum",
    "stage": "default todo",
    "teams": [
      "userId_1",
      "userId_2"
    ],
    "isDelete": false
  }
}
```

***
> Get Single Task (GET)
***

```code
api/v1/task/:id
```

### Role
- admin

### Response

```json
{
  "success": true,
  "message": "Successfully delete task!",
  "data": {
    "title": "Title task",
    "date" : "date now" ,
    "priority" : "string enum",
    "stage": "default todo",
    "teams": [
      "userId_1",
      "userId_2"
    ],
    "isDelete": true
  }
}
```

- login/register
- kalau admin dia bisa crud user, create task, delete task
- kalau user paling update task, create activity, update activity, delete activity