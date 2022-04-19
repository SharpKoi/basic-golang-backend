package database

/* This file defines the client's methods to provide database services to handlers
 * so that handlers can interact with the database by services instead of access database directly.
 */

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"time"
)

/** USER SERVICE **/

func (c Client) CheckUserExists(email string) (bool, error) {
	var tmp string
	err := c.db.QueryRow("select email from users where email = $1", email).Scan(&tmp)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func (c Client) GetUsers(limit int) ([]User, error) {
	var users []User

	// set limit to keep memory safe
	rows, err := c.db.Query("select * from users limit ($1)", limit)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id   int
			user User
		)
		err := rows.Scan(&id, &user.Email, &user.Password, &user.Name, &user.Age, &user.Role, &user.CreatedAt)
		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return users, err
	}

	return users, err
}

func (c Client) GetUser(email string) (User, error) {
	row := c.db.QueryRow("select * from users where email = $1", email)

	var (
		id   int
		user User
	)
	err := row.Scan(&id, &user.Email, &user.Password, &user.Name, &user.Age, &user.Role, &user.CreatedAt)

	return user, err
}

func (c Client) CreateUser(email, password, name string, age int, role string) (sql.Result, error) {
	if role == "" {
		// default role if role is not given
		role = "default"
	}
	stmt, err := c.db.Prepare("insert into users (email, password, name, age, role, createAt) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(email, password, name, age, role, time.Now().UTC())
	return res, err
}

func (c Client) UpdateUser(email, password, name string, age int, role string) (sql.Result, error) {
	stmt, err := c.db.Prepare("update users set password = $2, name = $3, age = $4, role = $5 where email = $1")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(email, password, name, age, role)
	return res, err
}

func (c Client) DeleteUser(email string) (sql.Result, error) {
	stmt, err := c.db.Prepare("delete from users where email = $1")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(email)
	return res, err
}

/** POST SERVICE **/

func (c Client) GetPostById(id string) (Post, error) {
	row := c.db.QueryRow("select * from posts where uid = $1", id)
	var post Post
	err := row.Scan(&post.ID, &post.UserEmail, &post.Text, &post.CreatedAt)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (c Client) GetPosts(userEmail string) ([]Post, error) {
	posts := []Post{}

	rows, err := c.db.Query("select * from posts where userEmail = $1", userEmail)
	if err != nil {
		return posts, err
	}
	defer rows.Close()

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.UserEmail, &post.Text, &post.CreatedAt)
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return posts, err
	}

	return posts, err
}

func (c Client) CreatePost(userEmail, text string) (sql.Result, error) {
	exists, err := c.CheckUserExists(userEmail)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, errors.New(fmt.Sprintf("user: \"%s\" not found", userEmail))
	}

	stmt, err := c.db.Prepare("insert into public.posts (uid, useremail, text, createat) values ($1, $2, $3, $4);")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(uuid.New().String(), userEmail, text, time.Now().UTC())
	return res, err
}

func (c Client) DeletePost(id string) (sql.Result, error) {
	stmt, err := c.db.Prepare("delete from posts where uid = $1")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(id)
	return res, err
}
