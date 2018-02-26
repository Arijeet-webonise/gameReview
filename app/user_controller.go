package app

import (
	"crypto/sha1"
	"database/sql"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Arijeet-webonise/gameReview/app/models"
	"github.com/Arijeet-webonise/gameReview/pkg/framework"
	"github.com/satori/go.uuid"
)

var dbUser = map[string]*models.User{}
var dbSession = map[string]string{}

func encrypt(text string) string {
	h := sha1.New()

	h.Write([]byte(text))

	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}

func (app *App) SignUpRender(w *framework.Response, r *framework.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/User/signup.html"}

	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w.ResponseWriter, res)
}

func (app *App) GetCurrrentUser(r *http.Request) (*models.User, error) {
	cookie, err := r.Cookie("auth")
	var user *models.User

	if err != nil {
		return user, err
	}

	session, err := app.SessionsService.SessionByUUID(cookie.Value, cookie.Value)

	if err != nil {
		return user, err
	}

	user, err = app.UserService.UserByID(session.Userid, session.Userid)

	if err != nil {
		return user, err
	}
	return user, nil
}

func (app *App) SignUpSubmit(w *framework.Response, r *framework.Request) {
	if err := r.Request.ParseForm(); err != nil {
		app.Log.Error(err)
		w.NotFound(err)
	}

	username := r.Request.FormValue("username")
	password := encrypt(r.Request.FormValue("password"))
	email := r.Request.FormValue("email")
	fname := r.Request.FormValue("fname")
	lname := r.Request.FormValue("lname")

	user := &models.User{
		Email:     sql.NullString{String: email, Valid: true},
		Password:  password,
		Username:  username,
		Firstname: fname,
		Lastname:  sql.NullString{String: lname, Valid: true},
		Roles:     sql.NullString{String: "user", Valid: true},
	}

	app.UserService.InsertUser(user)
	http.Redirect(w.ResponseWriter, r.Request, "/login", http.StatusSeeOther)
}

func (app *App) Login(w *framework.Response, r *framework.Request) {

	if r.Request.Method == http.MethodPost {
		if err := r.Request.ParseForm(); err != nil {
			app.Log.Error(err)
			app.Router.HandleNotFound(w.ResponseWriter, r.Request)
		}

		username := r.Request.FormValue("username")
		password := encrypt(r.Request.FormValue("password"))

		user, err := app.UserService.FindByUsername(username)

		if err != nil {
			app.Log.Error(err)
			err = errors.New("Username and/or password do not match")
			app.HandlerError(w.ResponseWriter, r.Request, err, http.StatusForbidden)
			return
		}

		if user.Password != password {
			app.Log.Error(err)
			err = errors.New("Username and/or password do not match")
			app.HandlerError(w.ResponseWriter, r.Request, err, http.StatusForbidden)
			return
		}

		uid, err := uuid.NewV4()
		if err != nil {
			app.Log.Error(err)
		}
		expiration := time.Now().Add(time.Hour)
		newCookie := &http.Cookie{}
		newCookie.Name = "auth"
		newCookie.Value = uid.String()
		newCookie.Expires = expiration

		http.SetCookie(w.ResponseWriter, newCookie)

		sessionObj := &models.Session{
			UUID:   uid.String(),
			Userid: user.ID,
		}

		if err = app.SessionsService.UpsertSession(sessionObj); err != nil {
			app.Log.Error(err)
		}

		http.Redirect(w.ResponseWriter, r.Request, "/login", http.StatusSeeOther)

	}
}

func (app *App) LoginRender(w *framework.Response, r *framework.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/User/login.html"}
	// app.CheckLogin(w, r)
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w.ResponseWriter, res)
}

func (app *App) Logout(w *framework.Response, r *framework.Request) {
	cookie, err := r.Cookie("auth")

	if err != nil {
		app.Log.Error(err)
		app.HandlerError(w.ResponseWriter, r.Request, err, http.StatusForbidden)
		return
	}

	cookie.MaxAge = -1

	http.SetCookie(w.ResponseWriter, cookie)
	http.Redirect(w.ResponseWriter, r.Request, "/login", http.StatusSeeOther)
}
