package infra

import (
	"project/config"
	"project/database"
	"project/handler"
	"project/log"
	"project/middleware"
	"project/repository"
	"project/service"

	"go.uber.org/zap"
)

type ServiceContext struct {
	Cacher     database.Cacher
	Cfg        config.Config
	Ctl        handler.Handler
	Log        *zap.Logger
	Middleware middleware.Middleware
}

func NewServiceContext(migrateDb bool, seedDb bool) (*ServiceContext, error) {

	handlerError := func(err error) (*ServiceContext, error) {
		return nil, err
	}

	// instance config
	config, err := config.LoadConfig(migrateDb, seedDb)
	if err != nil {
		handlerError(err)
	}

	// instance looger
	log, err := log.InitZapLogger(config)
	if err != nil {
		handlerError(err)
	}

	// instance database
	db, err := database.ConnectDB(config)
	if err != nil {
		handlerError(err)
	}

	rdb := database.NewCacher(config, 60*60)

	// instance repository
	repository := repository.NewRepository(db, rdb)

	// instance service
	service := service.NewService(repository)

	// instance controller
	Ctl := handler.NewHandler(service, log)

	mw := middleware.NewMiddleware(rdb)

	return &ServiceContext{Cacher: rdb, Cfg: config, Ctl: *Ctl, Log: log, Middleware: mw}, nil
}
