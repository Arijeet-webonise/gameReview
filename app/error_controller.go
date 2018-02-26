package app

//
// func (app *App) Handler404Error(w *framework.Response, r *framework.Request) {
// 	tmplList := []string{"./web/views/base.html",
// 		"./web/views/header.html",
// 		"./web/views/footer.html",
// 		"./web/views/error/404error.html"}
//
// 	w.ResponseWriter.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.ResponseWriter.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.ResponseWriter.WriteHeader(404)
//
// 	res, err := app.TplParser.ParseTemplate(tmplList, nil)
// 	if err != nil {
// 		app.Log.Error(err)
// 	}
// 	io.WriteString(w, res)
// }

// func (app *App) HandlerError(w http.ResponseWriter, r *http.Request, err error, code int) {
// 	tmplList := []string{"./web/views/base.html",
// 		"./web/views/header.html",
// 		"./web/views/footer.html",
// 		"./web/views/error/error.html"}
//
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	w.Header().Set("X-Content-Type-Options", "nosniff")
// 	w.WriteHeader(code)
//
// 	data := struct {
// 		Msg  string
// 		Code int
// 	}{err.Error(), code}
//
// 	res, err := app.TplParser.ParseTemplate(tmplList, data)
// 	if err != nil {
// 		app.Log.Error(err)
// 	}
// 	io.WriteString(w, res)
// }
