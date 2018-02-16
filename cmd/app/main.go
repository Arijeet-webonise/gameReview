package main

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Arijeet-webonise/gameReview/app"
	"github.com/Arijeet-webonise/gameReview/pkg/database"
	"github.com/Arijeet-webonise/gameReview/pkg/logger"
	"github.com/Arijeet-webonise/gameReview/pkg/templates"
	"github.com/go-zoo/bone"
	"gopkg.in/yaml.v2"
)

func ReadDBConfig() (map[string]string, error) {
	ConfigMap := make(map[string]string, 0)

	ymlFile, err := ioutil.ReadFile("./db/dbconf.yml")
	if err != nil {
		return ConfigMap, err
	}
	type DBOpen struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Dbname   string `yaml:"dbname"`
	}

	type DBConfig struct {
		Driver string `yaml:"driver"`
		Open   DBOpen `yaml:"open"`
	}

	type DBConfigRoot struct {
		Test        DBConfig `yaml:"test"`
		Development DBConfig `yaml:"development"`
	}

	var dbConf DBConfigRoot

	err = yaml.Unmarshal(ymlFile, &dbConf)
	if err != nil {
		return ConfigMap, err
	}

	ConfigMap["driver"] = dbConf.Development.Driver
	ConfigMap["host"] = dbConf.Development.Open.Host
	ConfigMap["db"] = dbConf.Development.Open.Dbname
	ConfigMap["user"] = dbConf.Development.Open.User
	ConfigMap["password"] = dbConf.Development.Open.Password

	return ConfigMap, err
}

func main() {

	//initialise the router
	router := bone.New()

	//initialise logger
	log := &logger.RealLogger{}
	log.Initialise()

	// need to constrcut an instance of the AppConfig from various environment vars
	cfg, cfgErr := constructAppConfig()
	//hydrate the map of DB connection params and send it
	dbConnectionParams, err := ReadDBConfig()
	db := &database.DatabaseWrapper{}

	dbConn, dbErr := db.Initialise(dbConnectionParams)
	if dbErr != nil || dbConn == nil || err != nil {
		panic(errors.New("could not initialise the DB"))
	}

	//if the configuration is not loaded then exit before startup
	if cfgErr != nil {
		panic("the configuration wasnt enabled")
	}

	a := &app.App{
		Router:    router,
		Cfg:       cfg,
		Log:       log,
		TplParser: &templates.TemplateParser{},
		DB:        dbConn,
	}

	a.InitRouter()
	err = http.ListenAndServe(cfg.Port, router)
	if err != nil {
		panic(err)
	}
}

func constructAppConfig() (*app.Config, error) {
	cfg := &app.Config{
		Port: ":9999",
	}
	return cfg, nil
}
