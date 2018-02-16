package app

import "net/http"

// InitRouter will intialise the router
func (app *App) InitRouter() {
	initialiseV1API(app)
}

func initialiseV1API(app *App) {
	//REST API
	app.Router.Get("/api/ping", app.handle(app.ping))
	// app.Router.Get("/api/todo/", app.handle(app.GetAllTodos))
	//VIEW
	app.Router.Get("/", app.renderView(app.RenderIndex))
	app.Router.Get("/reviews/games", app.renderView(app.RenderReviewList))
	app.Router.Get("/reviews/games/commentSubmit", app.renderView(app.SubmitComments))
	app.Router.Get("/reviews/games/add", app.renderView(app.AddReviews))
	app.Router.Post("/reviews/games/addSubmit", app.renderView(app.AddReviewsSubmit))
	app.Router.Get("/reviews/games/:id", app.renderView(app.RenderReview))

	//STATIC FILES
	fs := http.FileServer(http.Dir("web/assets/"))
	app.Router.Get("/static/", http.StripPrefix("/static/", fs))
}
