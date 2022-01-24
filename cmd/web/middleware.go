package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page ")
		next.ServeHTTP(w, r)
	})
}

//nosurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

//load and save provides middleware which automatically loads and saves session
//data for the current request, and communicate the session token to and from
// the client in a cookie

//session load loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {

	return session.LoadAndSave(next)
}
