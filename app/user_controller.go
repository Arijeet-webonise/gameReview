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

func (app *App) SignUpRender(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/User/signup.html"}

	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w, res)
}

func (app *App) CheckIfLogined(r *http.Request) error {
	cookie, err := r.Cookie("auth")

	if err != nil {
		return err
	}

	session, err := app.SessionsService.SessionByUUID(cookie.Value, cookie.Value)

	if err != nil {
		return err
	}

	_, err = app.UserService.UserByID(session.Userid, session.Userid)

	if err != nil {
		return err
	}
	return nil
}

func (app *App) SignUpSubmit(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.Log.Error(err)
		app.Router.HandleNotFound(w, r)
	}

	username := r.FormValue("username")
	password := encrypt(r.FormValue("password"))
	email := r.FormValue("email")
	fname := r.FormValue("fname")
	lname := r.FormValue("lname")

	user := &models.User{
		Email:     sql.NullString{String: email, Valid: true},
		Password:  password,
		Username:  username,
		Firstname: fname,
		Lastname:  sql.NullString{String: lname, Valid: true},
		Roles:     sql.NullString{String: "user", Valid: true},
	}

	app.UserService.InsertUser(user)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			app.Log.Error(err)
			app.Router.HandleNotFound(w, r)
		}

		username := r.FormValue("username")
		password := encrypt(r.FormValue("password"))

		user, err := app.UserService.FindByUsername(username)

		if err != nil {
			app.Log.Error(err)
			err = errors.New("Username and/or password do not match")
			app.HandlerError(w, r, err, http.StatusForbidden)
			return
		}

		if user.Password != password {
			app.Log.Error(err)
			err = errors.New("Username and/or password do not match")
			app.HandlerError(w, r, err, http.StatusForbidden)
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

		http.SetCookie(w, newCookie)

		sessionObj := &models.Session{
			UUID:   uid.String(),
			Userid: user.ID,
		}

		if err = app.SessionsService.UpsertSession(sessionObj); err != nil {
			app.Log.Error(err)
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)

	}
}

func (app *App) LoginRender(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/User/login.html"}
	// app.CheckLogin(w, r)
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w, res)
}

func (app *App) Logout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("auth")

	if err != nil {
		app.Log.Error(err)
		app.HandlerError(w, r, err, http.StatusForbidden)
		return
	}

	cookie.MaxAge = -1

	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
