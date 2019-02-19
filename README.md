# Golang API + Mongo database

 

## Signup
User signup  

Retrieve user credentials from the body and validate against database.
For invalid email or password, `send 400 - Bad Request` response.
For valid email and password, save user in database and send `201 - Created` response.  

Request  

```sh
curl \
  -X POST \
  http://localhost:8000/signup \
  -H "Content-Type: application/json" \
  -d '{"email":"jon@doe.com","password":"secret"}'
```
Response  

`201 - Created`  

```json
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "jon@doe.com",
  "password": "secret"
}
```


## Login
User login  

Retrieve user credentials from the body and validate against database.
For invalid credentials, send 401 - Unauthorized response.
For valid credentials, send 200 - OK response:
Generate JWT for the user and send it as response.
Each subsequent request must include JWT in the Authorization header.
Method: `POST`  
Path: `/login`  

Request  

```sh
curl \
  -X POST \
  http://localhost:8000/login \
  -H "Content-Type: application/json" \
  -d '{"email":"jon@doe.com","password":"secret"}'
```
Response  

`200 - OK`

```json
{
  "id": "58465b4ea6fe886d3215c6df",
  "email": "jon@doe.com",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
}
```
Client should store the token, for browsers, you may use local storage.  



## Create Task


For invalid token, send 400 - Bad Request response.
For valid token:
If user is not found, send 404 - Not Found response.
Add a follower to the specified user in the path parameter and send 200 - OK response.
Method: POST 
Path: /task/

Request

```sh
curl \
  -X POST \
  http://localhost:8000/task/
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"\
  -H "Content-Type: application/json" \
  -d '{ "TaskName":"Second task in Golang","Description":"Store data in mongo db and fetch", "Status":0}'

```
Response


`200 - OK`
```json
{
    "id": "5c6a84372bdf275b8ca78e51",
    "UserId": "5c6a73c62bdf2736badce882",
    "TaskName": "Second task in Golang",
    "Description": "Store data in mongo db and fetch",
    "Status": 0
}
```

## Update Task


For invalid request payload, send 400 - Bad Request response.
If user is not found, send 404 - Not Found response.
Otherwise save post in the database and return it via 201 - Created response.
Method: `POST`  
Path: `/updateTask/:id`  

Request 

```sh
curl \
  -X POST \
  http://localhost:8000//updateTask/5c6a7b492bdf2748d56e50a9 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0" \
 -H "Content-Type: application/json" \
  -d '{ "TaskName":"Edited task in Golang","Description":"Store data in mongo db and fetch", "Status":0}'
```
Response  

`200 -OK`




## Complete Task


Method: `Complete Task`  
Path: `/completeTask/:id`  

Request  

```sh
curl \
  -X GET \
  http://localhost:8000/completeTask/5c6a7b492bdf2748d56e50a9 \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
```
Response  

`200 - OK`
```
"Task completed sucessfully"
```

## Get All Task


Method: `fetchTasks`  
Path: `/fetchTasks/`  

Request  

```sh
curl \
  -X GET \
  http://localhost:8000/fetchTasks/ \
  -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0ODEyNjUxMjgsImlkIjoiNTg0NjViNGVhNmZlODg2ZDMyMTVjNmRmIn0.1IsGGxko1qMCsKkJDQ1NfmrZ945XVC9uZpcvDnKwpL0"
  -H "Content-Type: application/json" \
```
Response  

`200 - OK`

```json
[
    {
        "id": "5c6a7b492bdf2748d56e50a9",
        "UserId": "5c6a73c62bdf2736badce882",
        "TaskName": "First task in Golang",
        "Description": "Store data in mongo db",
        "Status": 0
    },
    {
        "id": "5c6a84372bdf275b8ca78e51",
        "UserId": "5c6a73c62bdf2736badce882",
        "TaskName": "Second task in Golang",
        "Description": "Store data in mongo db and fetch",
        "Status": 0
    }
]
```

