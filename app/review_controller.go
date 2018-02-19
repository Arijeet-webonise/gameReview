package app

import (
	"database/sql"
	"io"
	"net/http"
	"os"
	"strconv"
	"sync"

	"github.com/Arijeet-webonise/gameReview/app/models"
	"github.com/go-zoo/bone"
)

var wg sync.WaitGroup

//RenderIndex renders the index page
func (app *App) RenderReviewList(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/reviews/list.html"}

	games, err := app.ReviewService.GetAllGames()

	if err != nil {
		app.Log.Error(err)
	}

	data := struct {
		Games []*models.Game
	}{games}

	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Error(err)
	}

	io.WriteString(w, res)
}

type GameChanStuct struct {
	Game *models.Game
}

type GenreChanStuct struct {
	Genre []*models.Genre
}

type CommentChanStuct struct {
	Comments []*models.Comment
}

func GetGame(id int, app *App) (GameChanStuct, error) {

	game, err := app.ReviewService.GameByID(id, id)

	if err != nil {
		return GameChanStuct{Game: game}, err
	}
	data := GameChanStuct{
		Game: game,
	}
	return data, nil
}

func GetGenre(genresChan chan GenreChanStuct, id int, app *App) {
	defer wg.Done()

	genretogames, err := app.GenretogamerelationService.GetGenreOfGame(id)
	if err != nil {
		app.Log.Error(err)
	}

	var genres []*models.Genre

	for _, genretogame := range genretogames {
		genre, err := genretogame.GetGenres(app.DB)
		if err != nil {
			app.Log.Error(err)
		}
		genres = append(genres, genre)
	}

	data := GenreChanStuct{
		Genre: genres,
	}
	genresChan <- data
}

func GetComments(commentChan chan CommentChanStuct, id int, app *App) {
	defer wg.Done()

	comments, err := app.CommentService.GetGameComments(id)

	if err != nil {
		app.Log.Error(err)
	}

	data := CommentChanStuct{
		Comments: comments,
	}
	commentChan <- data
}

func GetUserRating(userRating chan int, app *App, gameid int) {
	defer wg.Done()

	rating, err := app.RatingViewService.GetGameTotalRating(gameid)

	if err != nil {
		app.Log.Error(err)
	}
	userRating <- int(float64(rating.TotalRating.Int64) / float64(rating.Count.Int64*10) * 100)
}

func (app *App) SubmitComments(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		app.Log.Error(err)
	}

	comment := r.Form.Get("comment")
	id, err := strconv.Atoi(r.Form.Get("id"))

	if err != nil {
		app.Log.Error(err)
	}

	rating, err := strconv.Atoi(r.Form.Get("rating"))
	if err != nil {
		app.Log.Error(err)
	}

	newComment := &models.Comment{
		Comment: comment,
		Game:    sql.NullInt64{Int64: int64(id), Valid: true},
		Rating:  rating,
	}

	err = app.CommentService.InsertComment(newComment)

	if err != nil {
		app.Log.Error(err)
	}

}

func (app *App) RenderReview(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/reviews/item.html"}

	id, err := strconv.Atoi(bone.GetValue(r, "id"))
	if err != nil {
		app.Log.Error(err)
		app.Router.HandleNotFound(w, r)
		return
	}

	genresChan := make(chan GenreChanStuct, 1)
	commentChan := make(chan CommentChanStuct, 1)
	userRatingChan := make(chan int, 1)

	var genres []*models.Genre
	games, err := GetGame(id, app)
	if err != nil {
		app.Log.Error(err)
		app.Router.HandleNotFound(w, r)
		return
	}

	go GetGenre(genresChan, id, app)
	wg.Add(1)

	go GetComments(commentChan, id, app)
	wg.Add(1)

	go GetUserRating(userRatingChan, app, id)
	wg.Add(1)

	wg.Wait()
	close(genresChan)
	close(commentChan)
	close(userRatingChan)

	var comments []*models.Comment
	var rating int

	for elem := range genresChan {
		genres = elem.Genre
	}

	for elem := range commentChan {
		comments = elem.Comments
	}

	for elem := range userRatingChan {
		rating = elem
	}

	data := struct {
		Game     *models.Game
		Genres   []*models.Genre
		Comments []*models.Comment
		Rating   int
	}{games.Game, genres, comments, rating}

	res, err := app.TplParser.ParseTemplate(tmplList, data)
	if err != nil {
		app.Log.Error(err)
	}

	_, err = io.WriteString(w, res)

	if err != nil {
		app.Log.Error(err)
	}
}

func (app *App) AddReviews(w http.ResponseWriter, r *http.Request) {
	tmplList := []string{"./web/views/base.html",
		"./web/views/header.html",
		"./web/views/footer.html",
		"./web/views/reviews/edit.html"}

	res, err := app.TplParser.ParseTemplate(tmplList, nil)

	if err != nil {
		app.Log.Error(err)
	}

	_, err = io.WriteString(w, res)

	if err != nil {
		app.Log.Error(err)
	}
}

func (app *App) AddReviewsSubmit(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(800)

	if err != nil {
		app.Log.Error(err)
		r.Response.StatusCode = http.StatusNotFound
	}
	img, header, err := r.FormFile("reviewimage")

	if err != nil {
		app.Log.Error(err)
	}

	if IsInArray("image/jpeg", header.Header["Content-Type"]) || IsInArray("image/jpg", header.Header["Content-Type"]) {
		app.Log.Info("Is a Image")
	}

	f, err := os.Create("./web/assets/upload/img/" + header.Filename)
	defer f.Close()
	if err != nil {
		app.Log.Error(err)
	}
	io.Copy(f, img)

	game := &models.Game{
		Title:     r.Form.Get("title"),
		Summary:   sql.NullString{String: r.Form.Get("addReview"), Valid: true},
		Developer: sql.NullString{String: r.Form.Get("developer"), Valid: true},
		Rating:    r.Form.Get("rating"),
		ImageName: sql.NullString{String: header.Filename, Valid: true},
	}

	err = app.ReviewService.InsertGame(game)

	if err != nil {
		app.Log.Error(err)
	}
}

func IsInArray(target string, list []string) bool {
	for _, item := range list {
		if target == item {
			return true
		}
	}
	return false
}
