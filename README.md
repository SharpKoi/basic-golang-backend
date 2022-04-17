# Advanced Golang Backend - Social Media

> :bulb: **Info: You're in advanced example which implemented more ideas. To see the original example please switch to `master` branch.**

This project is written in[ Go](https://go.dev/). After learning from [Boot.dev](https://boot.dev/courses/cs-track) Golang courses, trying to implement a simple RESTful backend API with golang standard library `net/http`, and using json as the database.

To run this project, type the command below:

```sh
git clone https://github.com/SharpKoi/golang-backend-example.git
cd golang-backend-example
go run .
```

And it will create a localhost server listening on port 8080.

You can use Postman or curl or any else to test this API.

## API

| Method | URL                  | Description                                                 |
| :----- | :------------------- | :---------------------------------------------------------- |
| GET    | /api/users           | Get all user accounts from database                         |
| GET    | /api/users/${email}  | Get the user account by user's email                        |
| POST   | /api/users           | Create an user account by given request body                |
| PUT    | /api/users/${email}  | Update the user account by the given email and request body |
| DELETE | /api/users/t${email} | Delete user account by the given email                      |
| GET    | /api/posts/${email}  | Get all the posts created by the user with the given email  |
| POST   | /api/posts           | Create a post by given request body                         |
| DELETE | /api/posts/${uuid}   | Delete a post by the given uuid                             |

Here's also a [full testing example](https://www.postman.com/science-architect-49213412/workspace/go-backend-examples/collection/17316452-4ed311e2-369b-46d9-aac2-cd8137b67a97?action=share&creator=17316452) created at Postman. You can use the given examples to test this API.

## TODO

These Ideas below are from [the final section](https://boot.dev/project/709a2e74-eb45-46ea-ac26-4b8e6a3ce3e6/ec5c7007-8ed2-4e17-a9c9-c54007d0e0fb) of the course, for extending the backend program.

- [x] Use PostgreSQL instead of a JSON file for the database layer
  - Use [pgx](https://github.com/jackc/pgx) as PostgreSQL Driver
  - Use [database/sql](https://pkg.go.dev/database/sql) to implement SQL operations
  - Alternatively there are more powerful tools like [migrate](https://github.com/golang-migrate/migrate), [gorm](https://github.com/go-gorm/gorm) to implement SQL database layer.

- [ ] Add proper authentication to each request, may use the [password validator](https://github.com/wagslane/go-password-validator) designed by Lane

- [ ] Allow users to save other data with their posts
- [ ] Add more unit tests
- [ ] Deploy the API on AWS, GCP, or Digital Ocean
- [ ] Dockerize it
- [ ] Add documentation using markdown files
- [ ] Write a frontend that interacts with the API, maybe a webpage or a mobile app

 And these ideas below are mine

- [ ] Write a full guide for the backend API designing
- [ ] Use [gin](https://github.com/gin-gonic/gin) to implement backend API, [go-oauth2](https://github.com/golang/oauth2) to implement authentications, and [gorm](https://github.com/go-gorm/gorm) to implement SQL database layer
- [ ] Integrate with discord bot
- [ ] Use sentimental analysis to detect the emotion of each post
