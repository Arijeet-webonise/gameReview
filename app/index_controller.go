package app

import (
	"io"

	"github.com/Arijeet-webonise/gameReview/pkg/framework"
)

//RenderIndex renders the index page
func (app *App) RenderIndex(w *framework.Response, r *framework.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/index.html"}
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w.ResponseWriter, res)
}
