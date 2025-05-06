package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/refresh-token", handleRefreshToken)
	http.HandleFunc("/get-server-cookie", handleGetServerCookie)

	fmt.Println("Server starting on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Only set cookie if it's not present in request
	cookie, err := r.Cookie("x-b-sess")
	serverCookieValue := "not set"
	if err == http.ErrNoCookie {
		expiration := time.Date(2025, 10, 6, 13, 28, 0, 0, time.UTC)
		cookie = &http.Cookie{
			Name:     "x-b-sess",
			Value:    "123",
			Expires:  expiration,
			HttpOnly: false,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, cookie)
		serverCookieValue = "123"
	} else {
		serverCookieValue = cookie.Value
	}

	// Parse and execute template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, map[string]string{
		"serverCookieValue": serverCookieValue,
	})
}

func handleRefreshToken(w http.ResponseWriter, r *http.Request) {
	// Set new cookie
	expiration := time.Date(2025, 10, 6, 13, 30, 0, 0, time.UTC)
	cookie := &http.Cookie{
		Name:     "x-b-sess",
		Value:    "234",
		Expires:  expiration,
		HttpOnly: false,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, cookie)

	// Return new value in response body
	w.Write([]byte("234"))
}

func handleGetServerCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("x-b-sess")
	if err == http.ErrNoCookie {
		w.Write([]byte("not set"))
		return
	}
	w.Write([]byte(cookie.Value))
}
