// Package main is the entry point of the application
package main

import (
	"fmt"
	"hacktiv/final-project/cmd"
	"hacktiv/final-project/infrastructure/repository/postgres"
	errorsController "hacktiv/final-project/infrastructure/restapi/controllers/errors"
	"net/http"
	"os"
	"strings"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"hacktiv/final-project/infrastructure/restapi/routes"
)

func main() {

	router := gin.Default()
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())

	// postgres connection
	postgresDB, err := postgres.NewGorm()
	if err != nil {
		_ = fmt.Errorf("fatal error in postgres file: %s", err)
		panic(err)
	}

	// commands handler
	cmd.Execute(postgresDB)

	router.Use(errorsController.Handler)

	// postgres routes
	routes.ApplicationV1Router(router, postgresDB)

	// startServer(router)
	router.Run()

}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)

	}

	var s *http.Server

	environment := os.Getenv("ENV")
	if environment == "production" {
		s = &http.Server{
			Handler:        router,
			ReadTimeout:    18000 * time.Second,
			WriteTimeout:   18000 * time.Second,
			MaxHeaderBytes: 1000 << 20,
		}
	} else {
		serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
		s = &http.Server{
			Addr:           serverPort,
			Handler:        router,
			ReadTimeout:    18000 * time.Second,
			WriteTimeout:   18000 * time.Second,
			MaxHeaderBytes: 1000 << 20,
		}
	}

	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error()))
		panic(err)

	}
}
