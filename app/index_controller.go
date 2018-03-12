package app

import (
	"io"
	"net/http"

	"github.com/Arijeet-webonise/gameReview/pkg/PDFWritter"
)

//RenderIndex renders the index page
func (app *App) RenderIndex(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/index.html"}
	res, err := app.TplParser.ParseTemplate(tmplList, nil)
	if err != nil {
		app.Log.Error(err)
	}
	io.WriteString(w, res)
}

func (app *App) RenderPDFIndex(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/pdf.html"}
	// "./web/views/User/Login.html"}

	body := "my body is ready"

	data := &struct {
		Body string
	}{body}

	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Info(err)
	}

	hpdf := PDFWritter.HTMLToPDF{}

	hpdf.AddPage(res)

	if err := hpdf.Generate2(w); err != nil {
		app.Log.Info(err)
	}
}
