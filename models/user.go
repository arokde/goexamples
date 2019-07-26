package models

import (
	"errors"
	"fmt"
)

/*
User is struct
*/
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

/*
GetUsers function this is actually returning slice aka list of user of objects
*/
func GetUsers() []*User {
	return users
}

/*
AddUser is adding user
*/
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("New usr not include id value")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

/*GetUserByID returns the user matching the id passed
 */
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

/*UpdateUser updates the argument u
 */
func UpdateUser(u User) (User, error) {
	for i, candidate := range users {
		if candidate.ID == u.ID {
			users[i] = &u
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

/*RemoveUserByID removes the passed user from the list
 */
func RemoveUserByID(userToRemove User) error {
	for i, candidate := range users {
		if candidate.ID == userToRemove.ID {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", userToRemove.ID)
}
