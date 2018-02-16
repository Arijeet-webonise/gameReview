package app

import (
	"database/sql"
	"fmt"
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

	reviewService := models.GameServiceImpl{
		DB: app.DB,
	}
	games, err := reviewService.GetAllGames()

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

func GetGame(gameChan chan GameChanStuct, id int, app *App) {
	defer wg.Done()
	reviewService := models.GameServiceImpl{
		DB: app.DB,
	}

	game, err := reviewService.GameByID(id, id)

	if err != nil {
		app.Log.Error(err)
	}
	data := GameChanStuct{
		Game: game,
	}
	gameChan <- data
}

func GetGenre(genresChan chan GenreChanStuct, id int, app *App) {
	defer wg.Done()

	genreService := models.GenretogamerelationServiceImpl{
		DB: app.DB,
	}
	genretogames, err := genreService.GetGenreOfGame(id)
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
	commentService := models.CommentServiceImpl{
		DB: app.DB,
	}

	comments, err := commentService.GetGameComments(id)

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
	ratingService := models.RatingViewServiceImpl{
		DB: app.DB,
	}

	rating, err := ratingService.GetGameTotalRating(gameid)

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

	fmt.Println(newComment)
	commentService := models.CommentServiceImpl{
		DB: app.DB,
	}

	err = commentService.InsertComment(newComment)

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
	}

	gameChan := make(chan GameChanStuct, 1)
	genresChan := make(chan GenreChanStuct, 1)
	commentChan := make(chan CommentChanStuct, 1)
	userRatingChan := make(chan int, 1)

	var genres []*models.Genre
	go GetGame(gameChan, id, app)
	wg.Add(1)

	go GetGenre(genresChan, id, app)
	wg.Add(1)

	go GetComments(commentChan, id, app)
	wg.Add(1)

	go GetUserRating(userRatingChan, app, id)
	wg.Add(1)

	wg.Wait()
	close(gameChan)
	close(genresChan)
	close(commentChan)
	close(userRatingChan)

	var game *models.Game
	var comments []*models.Comment
	var rating int

	for elem := range gameChan {
		game = elem.Game
	}

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
	}{game, genres, comments, rating}

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
	}
	img, header, err := r.FormFile("reviewimage")

	if err != nil {
		app.Log.Error(err)
	}

	fmt.Println(img)
	fmt.Println(header)

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

	gameService := models.GameServiceImpl{
		DB: app.DB,
	}

	err = gameService.InsertGame(game)

	if err != nil {
		app.Log.Error(err)
	}
}
