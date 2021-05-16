package handler

import (
	"github.com/TeamKun/mcapi/pkg/minecraft"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("mcapi\n"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.FormValue("user")
	password := r.FormValue("password")
	go minecraft.Login("localhost", user, password)
	w.WriteHeader(http.StatusOK)
}
