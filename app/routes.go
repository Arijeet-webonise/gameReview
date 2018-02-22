package app

import "net/http"

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {

	app.Router.NotFoundFunc(app.renderView(app.Handler404Error, false))

	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping, false))
	// app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex, false))
	app.Router.Get("/reviews/games", app.renderView(app.RenderReviewList, true))
	app.Router.Get("/reviews/games/commentSubmit", app.renderView(app.SubmitComments, true))
	app.Router.Get("/reviews/games/add", app.renderView(app.AddReviews, true))
	app.Router.Post("/reviews/games/addSubmit", app.renderView(app.AddReviewsSubmit, true))
	app.Router.Get("/reviews/games/:id", app.renderView(app.RenderReview, true))

	app.Router.Get("/login", app.renderView(app.LoginRender, false))
	app.Router.Post("/loginSubmit", app.renderView(app.Login, false))
	app.Router.Get("/signup", app.renderView(app.SignUpRender, false))
	app.Router.Post("/signupSubmit", app.renderView(app.SignUpSubmit, false))
	app.Router.Get("/logout", app.renderView(app.Logout, true))

	//STATIC FILES
	fs := http.FileServer(http.Dir("web/assets/"))
	app.Router.Get("/static/", http.StripPrefix("/static/", fs))
}
