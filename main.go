package main

import (
	"mgo-gin/app"
)


// @title Swagger Example API
// @version 1.0
// @description This is a sample swagger
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8585
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main(){
	var server app.Routes
	server.StartGin()
}