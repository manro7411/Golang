package handler

import (
	"net/http"

	"github.com/GDSC-KMUTT/totp-session/service"
)

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) userHandler {
	return userHandler{service: service}
}

func (h userHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignUp("", "")
	w.Write([]byte("User Created!"))
}

func (h userHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	// implement me
	h.service.SignIn("", "")
	w.Write([]byte("User Sign In"))
}
