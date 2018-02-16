package app

import (
	"net/http"

	"github.com/satori/go.uuid"
)

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")

	if err != nil {
		app.Log.Error(err)
	}

	id, err := uuid.NewV4()

	if err != nil {
		app.Log.Error(err)
	}

	cookie = &http.Cookie{
		Name:     "Session",
		Value:    id.String(),
		HttpOnly: true,
		// Secure: true,
	}
	http.SetCookie(w, cookie)
}
