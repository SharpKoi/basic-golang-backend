# Advanced Golang Backend - Social Media

> :bulb: **Info: You're in advanced example which implemented more ideas. To see the original example please switch to `master` branch.**

This project is written in[ Go](https://go.dev/). After learning from [Boot.dev](https://boot.dev/courses/cs-track) Golang courses, trying to implement a simple RESTful backend API with golang standard library **`net/http`**, and using ~~json~~ **PostgreSQL** as the database.

To run this project, please clone this project to your local machine first:

```sh
git clone https://github.com/SharpKoi/basic-golang-backend.git
cd basic-golang-backend
```

 And you need to set environment variables `DATABASE_URL` and `JWTSecret`:

```sh
export database_url=${your postgreSQL database URL}
export jwtsecret=${your jwt key}
```

Type the command below to run the project (make sure you have [docker](https://docs.docker.com/engine/install/) and [docker compose](https://docs.docker.com/compose/install/) installed):

```sh
docker-compose up --build
```

It will create a server binded on docker default bridge network address, either `127.17.0.0` or `0.0.0.0`, with port `8080`.

You can use Postman or curl or any else to test this API by sending requests to server address.

## API

| Method | URL                  | Description                                                 | Permission |
| :----- | :------------------- | :---------------------------------------------------------- | :--------: |
| POST   | /api/users/login     | Login to retrieve your json web token                       |  everyone  |
| GET    | /api/users           | Get all user accounts from database                         |   admin    |
| GET    | /api/users/${email}  | Get the user account by user's email                        |   owner+   |
| POST   | /api/users           | Create an user account by given request body                |   admin    |
| PUT    | /api/users/${email}  | Update the user account by the given email and request body |   owner+   |
| DELETE | /api/users/t${email} | Delete user account by the given email                      |   owner+   |
| GET    | /api/posts/${email}  | Get all the posts created by the user with the given email  |  everyone  |
| POST   | /api/posts           | Create a post by given request body                         |   owner    |
| DELETE | /api/posts/${uuid}   | Delete a post by the given uuid                             |   owner+   |

Here's also a [full testing example](https://www.postman.com/science-architect-49213412/workspace/go-backend-examples/collection/17316452-4ed311e2-369b-46d9-aac2-cd8137b67a97?action=share&creator=17316452) created at Postman. You can use the given examples to test this API.

## TODO

These Ideas below are from [the final section](https://boot.dev/project/709a2e74-eb45-46ea-ac26-4b8e6a3ce3e6/ec5c7007-8ed2-4e17-a9c9-c54007d0e0fb) of the course, for extending the backend program.

- [x] Use PostgreSQL instead of a JSON file for the database layer
  - Used [pgx](https://github.com/jackc/pgx) as PostgreSQL Driver
  - Used [database/sql](https://pkg.go.dev/database/sql) to implement SQL operations
  - Alternatively there are more powerful tools like [migrate](https://github.com/golang-migrate/migrate), [gorm](https://github.com/go-gorm/gorm) to implement SQL database layer.
- [x] Add proper authentication to each request, may use the [password validator](https://github.com/wagslane/go-password-validator) designed by Lane
  - Used [jwt-go](https://github.com/dgrijalva/jwt-go) to implement JWT authorization layer
  - TODO: use [password validator](https://github.com/wagslane/go-password-validator) to validate password strength
- [ ] Allow users to save other data with their posts
- [ ] Add more unit tests
- [ ] Deploy the API on AWS, GCP, or Digital Ocean
- [x] Dockerize it
  - Used [docker-compose](https://docs.docker.com/compose/) to build **golang** and **PostgreSQL** containers
    - **docker-compose** is a tool that help us to build multiple containers conveniently.
  - Used [bridge network](https://docs.docker.com/network/bridge/) to connect the two containers so that they  can communicate with each other 
- [ ] Add documentation using markdown files
- [ ] Write a frontend that interacts with the API, maybe a webpage or a mobile app

 And these ideas below are mine

- [ ] Write a full guide for the backend API designing
- [ ] Use [gin](https://github.com/gin-gonic/gin) to implement backend API, [go-oauth2](https://github.com/golang/oauth2) to implement authentications, and [gorm](https://github.com/go-gorm/gorm) to implement SQL database layer
- [ ] Integrate with discord bot
- [ ] Use sentimental analysis to detect the emotion of each post
