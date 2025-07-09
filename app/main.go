package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/OgiDac/CompanyTask/config"
	_ "github.com/OgiDac/CompanyTask/docs"
	"github.com/OgiDac/CompanyTask/domain"
	"github.com/OgiDac/CompanyTask/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           CompanyTask API
// @version         1.0
// @description     API documentation for the CompanyTask project
// @host            localhost:8000
// @BasePath        /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	app := config.App()

	db := app.DB
	db.AutoMigrate(&domain.User{})
	defer app.CloseDatabaseConnection()

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	timeout := time.Duration(5) * time.Second

	router.Setup(timeout, app.DB, r)

	srv := &http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()
	fmt.Println("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}

	fmt.Println("shutting down")
	os.Exit(0)
}
