package app

import (
	"net/http"

	"github.com/Arijeet-webonise/gameReview/pkg/framework"
)

// Handle will be serving only those requests that dont need to be authed
func (app *App) handle(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
		res.Write()
	})
}

//RenderView renders a view
func (app *App) renderView(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		handler(&res, &req)
	})
}

func (app *App) renderSecureView(handler func(*framework.Response, *framework.Request)) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := framework.NewResponse(w)
		req := framework.Request{Request: r}
		_, err := app.GetCurrrentUser(r)
		if err != nil {
			app.Log.Error(err)
			res.Error(err, http.StatusForbidden)
			return
		}
		handler(&res, &req)
	})

}
