package app

import (
	"fmt"
	"log/slog"
	"mgtest/config"
	"mgtest/internal/app/handlers"
	"mgtest/internal/repository"
	"mgtest/internal/service"
	"mgtest/internal/storage/mongo"
	"mgtest/internal/storage/redis"
	"mgtest/pkg/logger"
	"net/http"
)

type App struct {
	Server  http.Server
	S       *service.Service
	Log     *slog.Logger
	Handler handlers.Handler
}

func New() *App {
	cfg := config.MustLoad()

	mongoDB := mongo.New(cfg.MongoUrl, cfg.MongoUser, cfg.MongoPass)
	mongoClient := repository.NewMongo(mongoDB)

	redis := redis.NewRedis(cfg.RedisAddress)
	redisClient := repository.NewRedis(redis)

	logger := logger.SetupLogger(cfg.LoggerLvl)
	logger.Info("starting api service", slog.String("env", cfg.LoggerLvl))
	logger.Debug("debug messages enabled")

	service := service.New(mongoClient, redisClient)
	handler := handlers.New(service, logger)
	app := &App{
		Log:     logger,
		S:       service,
		Handler: *handler,
		Server: http.Server{
			Addr: fmt.Sprintf(":%s", cfg.Port),
		},
	}
	app.Server.Handler = app.NewRouter()
	logger.Info("Git it")
	return app
}
