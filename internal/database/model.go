package database

/* This file declares database models.
 */

import "time"

// User the user data
type User struct {
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"createdAt"`
}

// Post store data of a post
type Post struct {
	ID        string    `json:"id"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdAt"`
}
