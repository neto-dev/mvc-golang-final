/*
 * Ernesto Valenzuela Vargas. Internal License
 *
 * Contact: neto.dev@protonmail.com
 *
 (License Content)
*/

/*
 * Revision History:
 *     Initial:        2018/08/24        Ernesto Valenzuela V
 */

package main

import (
	"log"
	"os"
	"github.com/mvc-golang-final/config"
	"github.com/mvc-golang-final/db"
	"github.com/mvc-golang-final/middleware"
	"github.com/labstack/echo"
	mdlw "github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/color"
)

var port = ":7070"

func main() {
	var environment = ""
	if len(os.Args) > 1 {
		environment = os.Args[1]
	} else if os.Getenv("ENVIRONMENT") != "" {
		environment = os.Getenv("ENVIRONMENT")
	} else {
		environment = "Development"
	}

	print("⇨ Environment: ", color.Red(environment+"\n\n"))

	config := config.Environment[environment]

	/*
		Inicializamos la base de datos que utilizaremos

		We initialize the database that we will use
	*/
	DB := db.InitMysqlDB(config)

	/*
		Al terminar la ejecucion de la funcion, se ejecutara el metodo Close de Gorm

		When the execution of the function is finished, Gorm's Close method will be executed.
	*/
	defer DB.Close()

	/*
		Iniciamos una nueva instancia de las rutas de Echo

		We begin a new instance of Echo's routes
	*/
	router := echo.New()

	/*
		Instanciamos los middlewares internos

		We installed the internal middlewares
	*/
	middL := middleware.InitMiddleware()

	/*
		Configuramos los CORS

		We configure the CORS
	*/
	router.Use(mdlw.CORSWithConfig(mdlw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))


	/*
		Seteamos los middlewares que deseamos implementar

		Seteamos the middlewares that we want to implement
	*/
	router.Use(middL.FormatRequest)
	router.HideBanner = true

	/*
		Ejecutamos el middleware de rutas en una GOROUTINE

		We run the middleware of routes in a GOROUTINE
	*/
	go middL.Router(DB, *router)

	/*
		Ejecutamos la funcion server para ejecutar el servidor le pasamos el puerto y el puntero de router

		We execute the server function to execute the server we pass the port and the router pointer to him
	*/

	server(port, *router)
}

func server(port string, router echo.Echo) {
	log.Printf(color.Magenta("The server is running"))
	print(
		"─████████████─██████████████─\n",
		"─██░░░░░░░░██─██░░░░░░░░░░██─\n",
		"─██░░████████─██░░██████░░██─\n",
		"─██░░░░░░░░██─██░░░░░░░░░░██─\n",
		"─████████░░██─██░░████░░████─\n",
		"─██░░░░░░░░██─██░░██──░░░░██─\n",
		"─████████████─██████──██████─   ",
		color.Cyan("© v0.5.0-Alpha \n\n"))
	print(color.White("Implementation of Clean Architecture running on Echo. by. Ernesto Valenzuela & Sargento Robot\n\n"))
	print(color.Green("_____________________________________________________________O/_____________\n"))
	print(color.Green("                                                             O\\\n"))

	print("⇨ You can start testing from the route ", color.Red("http://localhost"+port+"\n\n"))

	/*
		Ejecutamos el servidor

		We run the server
	*/
	if err := router.Start(port); err != nil {
		log.Fatalf(color.Blue("⇨ error in ListenAndServe: %s"), err)
	}
}