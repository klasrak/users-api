<h1 align="center">
	<img
		width="450"
		alt="The Lounge"
		src="https://i.morioh.com/2019/11/29/9a4822127dc5.jpg">
</h1>

<h3 align="center">
	<strong>Users CRUD API</strong>
</h3>

<p align="center">
	<a href="https://drive.google.com/file/d/19G6_HW9ZlJPJhuXtd5OdEz7goC9l65rv/view?usp=sharing"><img
		alt="Insomnia Collection"
		src="https://img.shields.io/badge/Insomnia-collection-blueviolet"></a>
	<a href="https://golang.org/dl/go1.17.1.src.tar.gz"><img
		alt="Golang version"
		src="https://img.shields.io/badge/Go-v1.17-blue"></a>
</p>

<p align="center">
	<img src="https://i.ibb.co/kScQJKf/Screen-Shot-2021-09-20-at-10-18-18.png" width="550">
</p>

## **Overview**

- **Complete user registration with all CRUD operations.**
- **Age-based registration validation.** Only users of legal age can be registered.
- **E-mail and brazilian CPF validation.** Only users with valid data can be registered.
- **Easy to run.** Docker, docker-compose and makefile to make it easier.
- **Swagger to help make API calls.**
- **Harshly tested**. Type ```go test -v ./... -cover``` and see for yourself =).

## **Pre-requisites**
- [Go](https://golang.org/dl/go1.17.1.src.tar.gz), latest version preferred;
- [Docker](https://docs.docker.com/engine/install/) and [docker-compose](https://docs.docker.com/compose/install/);
## **Instructions**

All the dependencies for running the project are installed and configured via commands in the [Makefile](https://github.com/klasrak/users-api/blob/master/Makefile).

First, we need to add the project's .env:
```sh
$ cp .env.example .env
```

Next we need to install the dependencies for running the migrations, and run the migrations on the database.

Run the command:
```sh
$ make prepare
```

Attention: the script will install the dependencies (it may ask you to enter your user password), and WAIT for postgres to become available to receive connections before running the migrations. **Pay attention to the console logs, you should see something like this:**
```
go mod download && \
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest && \
	docker-compose up -d postgres && \
	./docker/entrypoint.sh 127.0.0.1:5432 && \
	/Library/Developer/CommandLineTools/usr/bin/make migrate-up N= && \
	sudo chown -R <user> ./.dbdata && \
	docker-compose down
Creating network "users-api_backend" with driver "bridge"
Pulling postgres (postgres:alpine)...
alpine: Pulling from library/postgres
a0d0a0d46f8b: Pull complete
5034a66b99e6: Pull complete
82e9eb77798b: Pull complete
314b9347faf5: Pull complete
2625be9fae82: Pull complete
5ec8358e2a99: Pull complete
2e9ccfc29d86: Pull complete
2a4d94e5dde0: Pull complete
Digest: sha256:a70babcd0e8f86272c35d6efcf8070c597c1f31b3d19727eece213a09929dd55
Status: Downloaded newer image for postgres:alpine
Creating postgres ... done
+ arg=127.0.0.1:5432
+ HOST=127.0.0.1
+ PORT=5432
+ docker-compose exec postgres sh -c pg_isready
/var/run/postgresql:5432 - no response
+ echo 'Postgres is unavailable - sleeping'
Postgres is unavailable - sleeping
+ sleep 1
+ docker-compose exec postgres sh -c pg_isready
/var/run/postgresql:5432 - no response
+ echo 'Postgres is unavailable - sleeping'
Postgres is unavailable - sleeping
+ sleep 1
+ docker-compose exec postgres sh -c pg_isready
/var/run/postgresql:5432 - no response
+ echo 'Postgres is unavailable - sleeping'
Postgres is unavailable - sleeping
+ sleep 1
+ docker-compose exec postgres sh -c pg_isready
/var/run/postgresql:5432 - accepting connections
+ sleep 2
+ echo 'Postgres is up - executing command'
Postgres is up - executing command
migrate -source file:///Users/<user>/<folder>/users-api/migrations -database postgres://postgres:123456@localhost:5432/users-api?sslmode=disable up ;
1/u add_users_table (185.868686ms)
Stopping postgres ... done
Removing postgres ... done
Removing network users-api_backend
```

Finally, run the command ```make init``` and this will start the application:

```
users-api    | running...
users-api    | API server listening at: [::]:2345
users-api    | 2021-09-20T14:05:24Z info layer=debugger launching process with args: [./tmp/main]
users-api    | 2021-09-20T14:05:25Z debug layer=debugger continuing
users-api    | 2021/09/20 14:05:25 Starting server...
users-api    | 2021/09/20 14:05:25 Connecting to database
users-api    | 2021/09/20 14:05:25 Starting postgres connection
users-api    | 2021/09/20 14:05:25 Injecting dependencies
users-api    | [GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.
users-api    |
users-api    | [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
users-api    |  - using env:	export GIN_MODE=release
users-api    |  - using code:	gin.SetMode(gin.ReleaseMode)
users-api    |
users-api    | [GIN-debug] GET    /api/v1/users             --> github.com/klasrak/users-api/handlers.(*Handler).GetAll-fm (4 handlers)
users-api    | [GIN-debug] GET    /api/v1/users/:id         --> github.com/klasrak/users-api/handlers.(*Handler).GetByID-fm (4 handlers)
users-api    | [GIN-debug] POST   /api/v1/users             --> github.com/klasrak/users-api/handlers.(*Handler).Create-fm (4 handlers)
users-api    | [GIN-debug] PUT    /api/v1/users/:id         --> github.com/klasrak/users-api/handlers.(*Handler).Update-fm (4 handlers)
users-api    | [GIN-debug] DELETE /api/v1/users/:id         --> github.com/klasrak/users-api/handlers.(*Handler).Delete-fm (4 handlers)
users-api    | [GIN-debug] GET    /docs/*any                --> github.com/swaggo/gin-swagger.CustomWrapHandler.func1 (4 handlers)
users-api    | 2021/09/20 14:05:25 Listening on port :8080
```

## **How it works**

If you use [Insomnia](https://insomnia.rest/download), download the Collection [here](https://drive.google.com/file/d/19G6_HW9ZlJPJhuXtd5OdEz7goC9l65rv/view?usp=sharing). This project also uses **Swagger**, so at any time you can open your browser at http://localhost:8080/docs/index.html and you can consume the API there.

### **First, let's add a few users:**

**POST** ```/users```
```sh
curl --request POST \
  --url http://localhost:8080/api/v1/users \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "John Doe da Siva",
	"email": "johndoe@mail.com",
	"cpf": "182.345.015-69",
	"birthdate": "1987-06-21T15:04:05Z"
}'
```
**RESPONSE** 201 CREATED
```json
{
  "id": "653565ef-6000-4021-8804-91f3369b3190",
  "name": "John Doe da Siva",
  "email": "johndoe@mail.com",
  "cpf": "182.345.015-69",
  "birthdate": "1987-06-21T00:00:00Z"
}
```

**Keep an eye on the information**. You must send valid **e-mail**, **brazilian CPF** and **date of birth**!

The date must be in **RFC3339** format (2006-01-02T15:04:05Z)

To help generate valid brazilian CPF, I recommend this [website](https://www.geradordecpf.org/).

If you send some of this invalid information, you will receive the following errors:

**RESPONSE** 400 BADREQUEST, Invalid e-mail:
```json
{
  "error": {
    "type": "BADREQUEST",
    "message": "Bad request. Reason: Invalid request parameters. See invalidArgs"
  },
  "invalidArgs": [
    {
      "field": "Email",
      "value": "invalid_email",
      "tag": "email",
      "param": ""
    }
  ]
}
```
**RESPONSE** 400 BADREQUEST, Invalid cpf:
```json
{
  "error": {
    "type": "BADREQUEST",
    "message": "Bad request. Reason: cpf invalid"
  }
}
```
And if the person you are trying to add is under the age of 18:
<br/>
<br/>
**RESPONSE** 400 BADREQUEST, underage:
```json
{
  "error": {
    "type": "BADREQUEST",
    "message": "Bad request. Reason: underage"
  }
}
```
Remember that the e-mail and the cpf are constraints, that is, they cannot be repeated. If you try to add someone with this repeated data, you will get the following error:
<br/>
<br/>
**RESPONSE** 409 CONFLICT, unique violation:
```json
{
  "error": {
    "type": "CONFLICT",
    "message": "resource: user not created: Key (email)=(johndoe@mail.com) already exists."
  }
}
```
```json
{
  "error": {
    "type": "CONFLICT",
    "message": "resource: user not created: Key (cpf)=(182.345.015-69) already exists."
  }
}
```
<br/>

### **Now that we have users in our database, we can call the resource:**

**GET** ```/users```:
```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/users'
```
**RESPONSE** 200 OK:
```json
[
  {
    "id": "653565ef-6000-4021-8804-91f3369b3190",
    "name": "John Doe da Siva",
    "email": "johndoe@mail.com",
    "cpf": "182.345.015-69",
    "birthdate": "1987-06-21T00:00:00Z"
  },
  {
    "id": "10285ad5-63c5-4ddd-9250-d86476566b80",
    "name": "Jane Doe Pereira",
    "email": "janedoe@mail.com",
    "cpf": "774.186.357-61",
    "birthdate": "2001-06-21T00:00:00Z"
  }
]
```
We can also filter our results and search by the name of the registered user:
```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/users?name=John'
```
**RESPONSE** 200 OK:
```json
[
  {
    "id": "653565ef-6000-4021-8804-91f3369b3190",
    "name": "John Doe da Siva",
    "email": "johndoe@mail.com",
    "cpf": "182.345.015-69",
    "birthdate": "1987-06-21T00:00:00Z"
  }
]
```

### **And how do we update the users' information?**

With the exception of ```id```, all other parameters can be updated **one at a time** or **all at once**. Let's see:

**PUT** ```/users/:id```
```sh
curl --request PUT \
  --url http://localhost:8080/api/v1/users/653565ef-6000-4021-8804-91f3369b3190 \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "John Doe da Siva Sauro",
	"email": "johndoe_novo_email@mail.com"
}'
```
**RESPONSE** 200 OK:
```json
{
  "id": "653565ef-6000-4021-8804-91f3369b3190",
  "name": "John Doe da Siva Sauro",
  "email": "johndoe_novo_email@mail.com",
  "cpf": "182.345.015-69",
  "birthdate": "1987-06-21T00:00:00Z"
}
```
<br/>

**PUT** ```/users/:id```
```sh
curl --request PUT \
  --url http://localhost:8080/api/v1/users/653565ef-6000-4021-8804-91f3369b3190 \
  --header 'Content-Type: application/json' \
  --data '{
	"email": "another_valid_email@mail.com"
}'
```
**RESPONSE** 200 OK:
```json
{
  "id": "653565ef-6000-4021-8804-91f3369b3190",
  "name": "John Doe da Siva Sauro",
  "email": "another_valid_email@mail.com",
  "cpf": "182.345.015-69",
  "birthdate": "1987-06-21T00:00:00Z"
}
```
Just a reminder that the validations for e-mail, age, cpf and unique constraints are still valid in the **PUT** ```/users/:id```.

### **What if we want to delete a user from our database?**
<br>
To complete our CRUD, we are going to delete a user from our database.
<br>
<br>

**DELETE** ```/users/:id```:
```sh
curl --request DELETE \
  --url 'http://localhost:8080/api/v1/users/653565ef-6000-4021-8804-91f3369b3190'
```
**RESPONSE** 204 NOCONTENT

That's it. User deleted. Our friend **John Doe** is no longer in our database.

**GET** ```/users```
```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/users'
````

**RESPONSE** 200 OK:
```json
[
  {
    "id": "10285ad5-63c5-4ddd-9250-d86476566b80",
    "name": "Jane Doe Pereira",
    "email": "janedoe@mail.com",
    "cpf": "774.186.357-61",
    "birthdate": "2001-06-21T00:00:00Z"
  }
]
```
<br/>
<br/>

## **Tests**

To run the tests, use the command ```go test -v ./... -cover```:
```sh
?   	github.com/klasrak/users-api	[no test files]
?   	github.com/klasrak/users-api/docs	[no test files]
=== RUN   TestUserHandler
=== RUN   TestUserHandler/GetAll
=== RUN   TestUserHandler/GetAll/Success_without_name_filter
[GIN] 2021/09/20 - 12:05:29 | 200 |     695.055µs |                 | GET      "/api/v1/users"
    user_handler_test.go:105: PASS:	GetAll(*context.emptyCtx,string)
=== RUN   TestUserHandler/GetAll/Success_with_name_filter
[GIN] 2021/09/20 - 12:05:29 | 200 |      66.457µs |                 | GET      "/api/v1/users?name=John+Doe"
    user_handler_test.go:158: PASS:	GetAll(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/GetAll/Success_no_content
[GIN] 2021/09/20 - 12:05:29 | 204 |      51.414µs |                 | GET      "/api/v1/users"
    user_handler_test.go:192: PASS:	GetAll(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/GetAll/Error
2021/09/20 12:05:29 Failed to get all users: Internal server error.
[GIN] 2021/09/20 - 12:05:29 | 500 |      79.271µs |                 | GET      "/api/v1/users"
    user_handler_test.go:228: PASS:	GetAll(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/GetByID
=== RUN   TestUserHandler/GetByID/Success
[GIN] 2021/09/20 - 12:05:29 | 200 |        79.2µs |                 | GET      "/api/v1/users/a51d7348-880e-47fb-a93e-a72e359b732a"
    user_handler_test.go:275: PASS:	GetByID(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/GetByID/Invalid_ID
2021/09/20 12:05:29 Failed to get user: Bad request. Reason: invalid id
[GIN] 2021/09/20 - 12:05:29 | 400 |      62.615µs |                 | GET      "/api/v1/users/invalid_id"
    user_handler_test.go:308: PASS:	GetByID(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/GetByID/Error_ID_not_found
2021/09/20 12:05:29 Failed to get user: resource: id with value: 6502e01f-270a-45a4-8d74-af5138be7a4f not found
[GIN] 2021/09/20 - 12:05:29 | 404 |      57.483µs |                 | GET      "/api/v1/users/invalid_id"
    user_handler_test.go:344: PASS:	GetByID(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/Create
=== RUN   TestUserHandler/Create/Success
[GIN] 2021/09/20 - 12:05:29 | 201 |     188.796µs |                 | POST     "/api/v1/users"
    user_handler_test.go:408: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserHandler/Create/Error_underage
2021/09/20 12:05:29 failed to create user: Bad request. Reason: underage
[GIN] 2021/09/20 - 12:05:29 | 400 |      76.039µs |                 | POST     "/api/v1/users"
    user_handler_test.go:458: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserHandler/Create/Error_invalid_cpf
2021/09/20 12:05:29 failed to create user: Bad request. Reason: cpf invalid
[GIN] 2021/09/20 - 12:05:29 | 400 |     652.527µs |                 | POST     "/api/v1/users"
    user_handler_test.go:508: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserHandler/Update
=== RUN   TestUserHandler/Update/Success
[GIN] 2021/09/20 - 12:05:29 | 200 |     137.573µs |                 | PUT      "/api/v1/users/91dfccff-c327-4705-978b-cc6b7c45aad8"
    user_handler_test.go:574: PASS:	Update(mock.AnythingOfTypeArgument,string,*model.User)
=== RUN   TestUserHandler/Update/Error_underage
2021/09/20 12:05:29 failed to update user: Bad request. Reason: underage
[GIN] 2021/09/20 - 12:05:29 | 400 |      89.214µs |                 | PUT      "/api/v1/users/28894b86-9092-4797-a79c-2fedc6a138d9"
    user_handler_test.go:627: PASS:	Update(mock.AnythingOfTypeArgument,string,*model.User)
=== RUN   TestUserHandler/Update/Error_invalid_email
2021/09/20 12:05:29 failed to update user: Bad request. Reason: invalid email
[GIN] 2021/09/20 - 12:05:29 | 400 |      86.734µs |                 | PUT      "/api/v1/users/2c77151d-60a8-4ff2-a5cd-f08c69200915"
    user_handler_test.go:680: PASS:	Update(mock.AnythingOfTypeArgument,string,*model.User)
=== RUN   TestUserHandler/Update/Error_invalid_cpf
2021/09/20 12:05:29 failed to update user: Bad request. Reason: cpf invalid
[GIN] 2021/09/20 - 12:05:29 | 400 |      87.424µs |                 | PUT      "/api/v1/users/15191535-36df-4d68-9a86-12471fb44126"
    user_handler_test.go:733: PASS:	Update(mock.AnythingOfTypeArgument,string,*model.User)
=== RUN   TestUserHandler/Delete
=== RUN   TestUserHandler/Delete/Success
[GIN] 2021/09/20 - 12:05:29 | 204 |      52.804µs |                 | DELETE   "/api/v1/users/e16640d3-86c6-4dca-b96f-d220e513e74b"
    user_handler_test.go:770: PASS:	Delete(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserHandler/Delete/Error
2021/09/20 12:05:29 failed to delete user: Internal server error.
[GIN] 2021/09/20 - 12:05:29 | 500 |      53.495µs |                 | DELETE   "/api/v1/users/0fe59d49-2007-439a-a87b-9cc7f705ad93"
    user_handler_test.go:806: PASS:	Delete(mock.AnythingOfTypeArgument,string)
--- PASS: TestUserHandler (0.01s)
    --- PASS: TestUserHandler/GetAll (0.00s)
        --- PASS: TestUserHandler/GetAll/Success_without_name_filter (0.00s)
        --- PASS: TestUserHandler/GetAll/Success_with_name_filter (0.00s)
        --- PASS: TestUserHandler/GetAll/Success_no_content (0.00s)
        --- PASS: TestUserHandler/GetAll/Error (0.00s)
    --- PASS: TestUserHandler/GetByID (0.00s)
        --- PASS: TestUserHandler/GetByID/Success (0.00s)
        --- PASS: TestUserHandler/GetByID/Invalid_ID (0.00s)
        --- PASS: TestUserHandler/GetByID/Error_ID_not_found (0.00s)
    --- PASS: TestUserHandler/Create (0.00s)
        --- PASS: TestUserHandler/Create/Success (0.00s)
        --- PASS: TestUserHandler/Create/Error_underage (0.00s)
        --- PASS: TestUserHandler/Create/Error_invalid_cpf (0.00s)
    --- PASS: TestUserHandler/Update (0.00s)
        --- PASS: TestUserHandler/Update/Success (0.00s)
        --- PASS: TestUserHandler/Update/Error_underage (0.00s)
        --- PASS: TestUserHandler/Update/Error_invalid_email (0.00s)
        --- PASS: TestUserHandler/Update/Error_invalid_cpf (0.00s)
    --- PASS: TestUserHandler/Delete (0.00s)
        --- PASS: TestUserHandler/Delete/Success (0.00s)
        --- PASS: TestUserHandler/Delete/Error (0.00s)
PASS
coverage: 65.9% of statements
ok  	github.com/klasrak/users-api/handlers	0.807s	coverage: 65.9% of statements
?   	github.com/klasrak/users-api/mocks	[no test files]
?   	github.com/klasrak/users-api/models	[no test files]
=== RUN   TestUserRepository
=== RUN   TestUserRepository/GetAll
=== RUN   TestUserRepository/GetAll/Success_without_name_filter
=== RUN   TestUserRepository/GetAll/Success_with_name_filter
=== RUN   TestUserRepository/GetAll/Error
=== RUN   TestUserRepository/GetByID
=== RUN   TestUserRepository/GetByID/Success
=== RUN   TestUserRepository/GetByID/Error
=== RUN   TestUserRepository/Create
=== RUN   TestUserRepository/Create/Success
=== RUN   TestUserRepository/Create/Error_unique_validation
2021/09/20 12:05:29 could not create user. Reason: pq:
=== RUN   TestUserRepository/Create/Internal_Server_Error
2021/09/20 12:05:29 failed to create user. Reason: error
=== RUN   TestUserRepository/Update
=== RUN   TestUserRepository/Update/Success
    pg_user_repository_test.go:291:
=== RUN   TestUserRepository/Delete
=== RUN   TestUserRepository/Delete/Success
=== RUN   TestUserRepository/Delete/Error_not_found
=== RUN   TestUserRepository/Delete/Internal_Server_Error
2021/09/20 12:05:29 failed to delete user. Reason: error
--- PASS: TestUserRepository (0.00s)
    --- PASS: TestUserRepository/GetAll (0.00s)
        --- PASS: TestUserRepository/GetAll/Success_without_name_filter (0.00s)
        --- PASS: TestUserRepository/GetAll/Success_with_name_filter (0.00s)
        --- PASS: TestUserRepository/GetAll/Error (0.00s)
    --- PASS: TestUserRepository/GetByID (0.00s)
        --- PASS: TestUserRepository/GetByID/Success (0.00s)
        --- PASS: TestUserRepository/GetByID/Error (0.00s)
    --- PASS: TestUserRepository/Create (0.00s)
        --- PASS: TestUserRepository/Create/Success (0.00s)
        --- PASS: TestUserRepository/Create/Error_unique_validation (0.00s)
        --- PASS: TestUserRepository/Create/Internal_Server_Error (0.00s)
    --- PASS: TestUserRepository/Update (0.00s)
        --- SKIP: TestUserRepository/Update/Success (0.00s)
    --- PASS: TestUserRepository/Delete (0.00s)
        --- PASS: TestUserRepository/Delete/Success (0.00s)
        --- PASS: TestUserRepository/Delete/Error_not_found (0.00s)
        --- PASS: TestUserRepository/Delete/Internal_Server_Error (0.00s)
PASS
coverage: 63.0% of statements
ok  	github.com/klasrak/users-api/repository	0.249s	coverage: 63.0% of statements
?   	github.com/klasrak/users-api/rerrors	[no test files]
=== RUN   TestUserService
=== RUN   TestUserService/GetAll
=== RUN   TestUserService/GetAll/Success_without_name_filter
    user_service_test.go:46: PASS:	GetAll(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserService/GetAll/Success_with_name_filter
    user_service_test.go:79: PASS:	GetAll(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserService/GetAll/Internal_Server_Error
    user_service_test.go:104: PASS:	GetAll(string,string)
=== RUN   TestUserService/GetByID
=== RUN   TestUserService/GetByID/Success
=== RUN   TestUserService/GetByID/Error_invalid_id
=== RUN   TestUserService/GetByID/Error_not_found
=== RUN   TestUserService/Create
=== RUN   TestUserService/Create/Success
    user_service_test.go:214: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Create/Error_unique_violation_email
    user_service_test.go:245: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Create/Error_unique_violation_cpf
    user_service_test.go:276: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Create/Internal_Server_Error
    user_service_test.go:307: PASS:	Create(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Create/Bad_request_underage
=== RUN   TestUserService/Create/Bad_request_invalid_cpf
=== RUN   TestUserService/Update
=== RUN   TestUserService/Update/Success
=== RUN   TestUserService/Update/Error_unique_violation_email
    user_service_test.go:439: PASS:	Update(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Update/Error_unique_violation_cpf
    user_service_test.go:472: PASS:	Update(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Update/Internal_Server_Error
    user_service_test.go:505: PASS:	Update(mock.AnythingOfTypeArgument,*model.User)
=== RUN   TestUserService/Update/Error_not_found
=== RUN   TestUserService/Update/Bad_request_underage
=== RUN   TestUserService/Update/Bad_request_invalid_cpf
=== RUN   TestUserService/Update/Bad_request_invalid_email
=== RUN   TestUserService/Delete
=== RUN   TestUserService/Delete/Success
    user_service_test.go:651: PASS:	Delete(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserService/Delete/Error_not_found
    user_service_test.go:675: PASS:	Delete(mock.AnythingOfTypeArgument,string)
=== RUN   TestUserService/Delete/Internal_Server_Error
    user_service_test.go:699: PASS:	Delete(mock.AnythingOfTypeArgument,string)
--- PASS: TestUserService (0.00s)
    --- PASS: TestUserService/GetAll (0.00s)
        --- PASS: TestUserService/GetAll/Success_without_name_filter (0.00s)
        --- PASS: TestUserService/GetAll/Success_with_name_filter (0.00s)
        --- PASS: TestUserService/GetAll/Internal_Server_Error (0.00s)
    --- PASS: TestUserService/GetByID (0.00s)
        --- PASS: TestUserService/GetByID/Success (0.00s)
        --- PASS: TestUserService/GetByID/Error_invalid_id (0.00s)
        --- PASS: TestUserService/GetByID/Error_not_found (0.00s)
    --- PASS: TestUserService/Create (0.00s)
        --- PASS: TestUserService/Create/Success (0.00s)
        --- PASS: TestUserService/Create/Error_unique_violation_email (0.00s)
        --- PASS: TestUserService/Create/Error_unique_violation_cpf (0.00s)
        --- PASS: TestUserService/Create/Internal_Server_Error (0.00s)
        --- PASS: TestUserService/Create/Bad_request_underage (0.00s)
        --- PASS: TestUserService/Create/Bad_request_invalid_cpf (0.00s)
    --- PASS: TestUserService/Update (0.00s)
        --- PASS: TestUserService/Update/Success (0.00s)
        --- PASS: TestUserService/Update/Error_unique_violation_email (0.00s)
        --- PASS: TestUserService/Update/Error_unique_violation_cpf (0.00s)
        --- PASS: TestUserService/Update/Internal_Server_Error (0.00s)
        --- PASS: TestUserService/Update/Error_not_found (0.00s)
        --- PASS: TestUserService/Update/Bad_request_underage (0.00s)
        --- PASS: TestUserService/Update/Bad_request_invalid_cpf (0.00s)
        --- PASS: TestUserService/Update/Bad_request_invalid_email (0.00s)
    --- PASS: TestUserService/Delete (0.00s)
        --- PASS: TestUserService/Delete/Success (0.00s)
        --- PASS: TestUserService/Delete/Error_not_found (0.00s)
        --- PASS: TestUserService/Delete/Internal_Server_Error (0.00s)
PASS
coverage: 96.2% of statements
ok  	github.com/klasrak/users-api/service	0.605s	coverage: 96.2% of statements
=== RUN   TestIsBrazilianCPFValid
=== RUN   TestIsBrazilianCPFValid/Success_with_valid_masked_cpf
=== RUN   TestIsBrazilianCPFValid/Success_with_valid_unmasked_cpf
=== RUN   TestIsBrazilianCPFValid/Invalid_with_masked_cpf
=== RUN   TestIsBrazilianCPFValid/Invalid_with_unmasked_cpf
--- PASS: TestIsBrazilianCPFValid (0.00s)
    --- PASS: TestIsBrazilianCPFValid/Success_with_valid_masked_cpf (0.00s)
    --- PASS: TestIsBrazilianCPFValid/Success_with_valid_unmasked_cpf (0.00s)
    --- PASS: TestIsBrazilianCPFValid/Invalid_with_masked_cpf (0.00s)
    --- PASS: TestIsBrazilianCPFValid/Invalid_with_unmasked_cpf (0.00s)
=== RUN   TestSanitizeUpdateParams
=== RUN   TestSanitizeUpdateParams/Should_parse_empty_string_to_sql.NullString
=== RUN   TestSanitizeUpdateParams/Should_parse_time.Time{}_to_sql.NullString
=== RUN   TestSanitizeUpdateParams/Do_nothing_with_valid_string
=== RUN   TestSanitizeUpdateParams/Parse_valid_time.Time
=== RUN   TestSanitizeUpdateParams/Default
=== RUN   TestSanitizeUpdateParams/Internal_Server_Error
--- PASS: TestSanitizeUpdateParams (0.00s)
    --- PASS: TestSanitizeUpdateParams/Should_parse_empty_string_to_sql.NullString (0.00s)
    --- PASS: TestSanitizeUpdateParams/Should_parse_time.Time{}_to_sql.NullString (0.00s)
    --- PASS: TestSanitizeUpdateParams/Do_nothing_with_valid_string (0.00s)
    --- PASS: TestSanitizeUpdateParams/Parse_valid_time.Time (0.00s)
    --- PASS: TestSanitizeUpdateParams/Default (0.00s)
    --- PASS: TestSanitizeUpdateParams/Internal_Server_Error (0.00s)
=== RUN   TestTimeBetween
--- PASS: TestTimeBetween (0.00s)
=== RUN   TestIsUnderage
=== RUN   TestIsUnderage/True
=== RUN   TestIsUnderage/False
--- PASS: TestIsUnderage (0.00s)
    --- PASS: TestIsUnderage/True (0.00s)
    --- PASS: TestIsUnderage/False (0.00s)
PASS
coverage: 78.1% of statements
ok  	github.com/klasrak/users-api/utils	0.389s	coverage: 78.1% of statements
```


## **Project purpose**

This project was (and will continue to be) developed for the purpose of studies. Any feedback, positive or not, will be very welcome and will be taken into consideration.

## **License**

This project follows [MIT License](https://github.com/klasrak/users-api/blob/master/LICENSE).
