package main

import (
	"API3/config"
	"API3/handler"
	"API3/helper"
	"API3/model"
	"API3/respository"
	"API3/router"
	"API3/service"
	"github.com/go-playground/validator/v10"
	"github.com/rs/zerolog/log"
	"net/http"
)

func main() {
	log.Info().Msg("Starting server!")

	// Database
	db := config.DatabaseConnection()
	validate := validator.New()

	db.Table("tags").AutoMigrate(&model.Tags{})

	// Repository
	tagsRepository := respository.NewTagsRepositoryImpl(db)

	// Service
	tagsService := service.NewTagsServiceImpl(tagsRepository, validate)

	// Controller
	tagsController := handler.NewTagsHandler(tagsService)

	// Router
	routes := router.NewRouter(tagsController)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes,
	}
	err := server.ListenAndServe()
	helper.ErrorPanic(err)
}
