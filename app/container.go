package app

import (
	"database/sql"

	"github.com/Arijeet-webonise/gameReview/app/models"
	"github.com/Arijeet-webonise/gameReview/pkg/logger"
	"github.com/Arijeet-webonise/gameReview/pkg/templates"
	"github.com/go-zoo/bone"
)

// App enscapsulates the App environment
type App struct {
	Router                     *bone.Mux
	Cfg                        *Config
	Log                        logger.ILogger
	TplParser                  templates.ITemplateParser
	DB                         *sql.DB
	ReviewService              *models.GameServiceImpl
	GenretogamerelationService *models.GenretogamerelationServiceImpl
	CommentService             *models.CommentServiceImpl
	RatingViewService          *models.RatingViewServiceImpl
	UserService                *models.UserServiceImpl
}
