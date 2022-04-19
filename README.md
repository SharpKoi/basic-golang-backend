# Golang Backend Example - Social Media

> :bulb: **Now you can see the [advanced example](https://github.com/SharpKoi/golang-backend-example/tree/advance) which implemented more ideas by switching to `advance` branch**.

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

See the new updates in [advance](https://github.com/SharpKoi/golang-backend-example/tree/advance) branch 🤩
