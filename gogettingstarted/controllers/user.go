package controllers

import (
	"net/http"
	"regexp"

	"github.com/ameyrokde/webservice/models"
)

type userController struct {
	userIDPattern *regexp.Regexp
}

func (uc userController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from User Controller"))
}

func (uc userController) getAll(w http.ResponseWriter, r *http.Request) {
	encodeResponseAsJSONWithList(models.GetUsers(), w)
}

func (uc userController) get(id int, w http.ResponseWriter) {
	u, err := models.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(u, w)
}

func (uc userController) parseRequest(r *http.Request) (models.User, error) {
	return models.User{}, nil
}

func (uc userController) post(w http.ResponseWriter, r *http.Request) {
	u, err := uc.parseRequest(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	encodeResponseAsJSON(u, w)
}

func newUserController() *userController {
	return &userController{
		userIDPattern: regexp.MustCompile("^/users/(\\d+)/?"),
	}
}

func encodeResponseAsJSON(user models.User, w http.ResponseWriter) {

}

func encodeResponseAsJSONWithList(users []*models.User, w http.ResponseWriter) {

}
