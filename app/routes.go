package app

import "net/http"

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {

	app.Router.NotFoundFunc(app.renderView(app.Handler404Error))

	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	// app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/reviews/games", app.renderView(app.RenderReviewList))
	app.Router.Get("/reviews/games/commentSubmit", app.renderView(app.SubmitComments))
	app.Router.Get("/reviews/games/add", app.renderView(app.AddReviews))
	app.Router.Post("/reviews/games/addSubmit", app.renderView(app.AddReviewsSubmit))
	//app.Router.Get("/reviews/games/:id", app.renderView(app.RenderReview))
	app.Router.Get("/reviews/games/:id", app.renderSecureView(app.RenderReview))

	app.Router.Get("/login", app.renderView(app.LoginRender))
	app.Router.Post("/loginSubmit", app.renderView(app.Login))
	app.Router.Get("/signup", app.renderView(app.SignUpRender))
	app.Router.Post("/signupSubmit", app.renderView(app.SignUpSubmit))
	app.Router.Get("/logout", app.renderView(app.Logout))

	//STATIC FILES
	fs := http.FileServer(http.Dir("web/assets/"))
	app.Router.Get("/static/", http.StripPrefix("/static/", fs))
}
