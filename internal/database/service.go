package database

/* This file defines the client's methods to provide database services to handlers
 * so that handlers can interact with the database by services instead of access database directly.
 */

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

/** USER SERVICE **/

func (c Client) GetUsers() ([]User, error) {
	var users []User

	schema, err := c.readDB()
	if err != nil {
		return []User{}, err
	}
	for _, user := range schema.Users {
		users = append(users, user)
	}

	return users, err
}

func (c Client) GetUser(email string) (User, error) {
	schema, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, exists := schema.Users[email]
	if !exists {
		return User{}, errors.New("user doesn't exist")
	}

	return user, err
}

func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	schema, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	_user, exists := schema.Users[email]
	if exists {
		return _user, errors.New("user's email has already been registered")
	}

	// create a new instance
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}

	// update schema and store it in the database
	schema.Users[email] = user
	err = c.writeDB(schema)

	return user, err
}

func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	schema, err := c.readDB()
	if err != nil {
		return User{}, err
	}

	user, exists := schema.Users[email]
	if !exists {
		return User{}, errors.New("user doesn't exist")
	}

	user.Password = password
	user.Name = name
	user.Age = age

	// set as the updated user
	schema.Users[email] = user
	// remember to save the updated schema
	err = c.writeDB(schema)

	return user, err
}

func (c Client) DeleteUser(email string) error {
	schema, err := c.readDB()
	if err != nil {
		return err
	}

	// remember to save schema after the user deleted :)
	delete(schema.Users, email)
	err = c.writeDB(schema)

	return err // actually nil
}

/** POST SERVICE **/

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	posts := make([]Post, 0, 10)

	schema, err := c.readDB()
	if err != nil {
		return posts, err
	}

	for _, post := range schema.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}

	return posts, err
}

func (c Client) CreatePost(userEmail, text string) (Post, error) {
	schema, err := c.readDB()
	if err != nil {
		return Post{}, err
	}

	_, exists := schema.Users[userEmail]
	if !exists {
		return Post{}, errors.New("user doesn't exist")
	}

	post := Post{
		ID:        uuid.New().String(),
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text:      text,
	}

	schema.Posts[post.ID] = post
	err = c.writeDB(schema)

	return post, err
}

func (c Client) DeletePost(id string) error {
	schema, err := c.readDB()
	if err != nil {
		return err
	}

	// remember to save schema after the post deleted :)
	delete(schema.Posts, id)
	err = c.writeDB(schema)

	return err
}
