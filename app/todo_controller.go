package app

import (
	"errors"
	"io"
	"net/http"

	"github.com/Arijeet-webonise/gameReview/app/models"
	"github.com/Arijeet-webonise/gameReview/pkg/framework"
)

func (a App) GetAllTodos(w *framework.Response, r *framework.Request) {

	role, err := models.GetAllTodos(a.DB)
	if err != nil {
		a.Log.Info(err.Error())
		w.NotFound(errors.New("could not find the role"))
		return
	}
	a.Log.Info(role)
	if err != nil {
		panic(err)
		return
	}
	w.PutInData("roles", role)
}

//RenderIndex renders the index page
func (app *App) DisplayTodos(w http.ResponseWriter, r *http.Request) {
	roles, err := models.GetAllTodos(app.DB)
	tmplList := []string{"./web/views/base.html",
		"./web/views/todos/todo.html"}
	data := struct {
		Roles []*models.Todo
	}{roles}
	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Info(err)
	}
	io.WriteString(w, res)
}
