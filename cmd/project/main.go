package main

import (
	"test-swagger/api"

	// import file swagger docs
	_ "test-swagger/docs"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server Petstore server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		petstore.swagger.io
//	@BasePath	/v2

func main() {
	e := echo.New()

	e.GET("/testapi/get-string-by-int/", api.GetUserByID())

	// call the echoSwagger
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":1323"))
}

// --------------------------------------------------------/
// type in terminal
// create the swagger docs "swag init -g cmd/project/main.go"
