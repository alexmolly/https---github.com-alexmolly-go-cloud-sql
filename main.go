package main

import (
	"fmt"
	"os"

	"github.com/alexmolly/gomvcboilerplate/config"
	"github.com/alexmolly/gomvcboilerplate/routers"
	"github.com/alexmolly/gomvcboilerplate/services"
	"github.com/alexmolly/gomvcboilerplate/services/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	Router *gin.Engine
}

func (server *Server) InitServer() {
	server.Router = gin.Default()
	var Router = routers.RouteLoader{}
	for _, routes := range Router.LoadRoutes() {
		routes.Route(server.Router)
	}
}

func (server *Server) InjectDBToService() {
	services.InjectDBIntoServices(server.db)
}

func (server *Server) Run() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println("Rise and shine! 🌞🌞🌞")
	fmt.Println(config.ServerConfig.AppConfig.AppName + " is listening on port : " + port)
	server.Router.Run(":" + port)

}
func main() {
	// Load Config from Env
	config.LoadConfig()

	// Uncomment this and remove Server{} to use database
	app := Server{db: db.LoadDB(config.ServerConfig.DBConfig)}
	app.InjectDBToService()

	app.InitServer()
	app.Run()
}
