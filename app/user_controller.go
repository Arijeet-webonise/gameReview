package app

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/Arijeet-webonise/gameReview/app/models"
	"github.com/satori/go.uuid"
)

var dbUser = map[string]models.User{}
var dbSession = map[string]string{}

func (app *App) GetCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")

	if err != nil {
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

	return cookie
}

func encrypt(text string) string {
	h := sha1.New()

	h.Write([]byte(text))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {
	// cookie := app.GetCookie(w, r)
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			app.Log.Error(err)
		}

		username := r.FormValue("username")
		password := encrypt(r.FormValue("password"))

		user, err := app.UserService.FindByUsername(username)

		if err != nil {
			app.Log.Error(err)
		}

		if user.Password == password {
			app.Log.Info(user.Username, "Successfully Login")
		} else {
			err = errors.New("Invalid User")
			app.Log.Error(err)
			app.Router.HandleNotFound(w, r)
		}
	}
}

func (app *App) LoginRender(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/User/login.html"}

	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w, res)
}
