package controllers

import (
	"fmt"
	"net/http"
)

/*
RegisterController is registered
*/
func RegisterController() {
	uc := newUserController()
	fmt.Println(uc)
	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}
