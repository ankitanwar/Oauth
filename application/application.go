package application

import (
	mongodb "github.com/ankitanwar/bookStore-OAuth/clients"
	"github.com/ankitanwar/bookStore-OAuth/http"
	"github.com/ankitanwar/bookStore-OAuth/repository/db"
	"github.com/ankitanwar/bookStore-OAuth/services"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

//StartApplication : To start the application
func StartApplication() {
	mongodb.Ping()
	dbRepository := db.NewRepository()
	atService := services.NewService(dbRepository)
	atHandler := http.NewHandler(atService)
	router.GET("/oauth/access_token/:access_token_id", atHandler.GetByID)
	router.POST("/oauth/access_token", atHandler.Create)
	router.Run(":8080")
}
