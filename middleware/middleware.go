package middleware

/*
Importamos las librerias

We import libraries
*/
import (
	"log"

	"github.com/mvc-golang-final/config"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/color"
)

/*
Estructura por la cual se podran llamar algunos metodos ya que instancian de ella

Structure by which some methods can be called since they instantiate it.
*/
type goMiddleware struct {
	// another stuff , may be needed by middleware
}

/*
Funcion encargade de generar las rutas de la api

Function in charge of generating the routes of the api
*/
func (m *goMiddleware) Router(DB *gorm.DB, router echo.Echo) echo.Echo {
	/*
		Se crea un grupo de ruta para manejar la version actual del proyecto

		A path group is created to handle the current version of the project
	*/
	groupApi := router.Group("/v1")

	/*
		Se ejecuta en una GOROUTINE la funcion que contiene los metodos encargados de generar las rutas (Repository,
		UseCase, Handlers)

		It is executed in a GOROUTINE the function that contains the methods in charge of generating the routes (Repository,
		UseCase, Handlers)
	*/
	go config.Routes(DB, *groupApi)

	/*
		Retornamos las rutas generadas

		We return the generated routes
	*/
	return router
}

/*
Funcion encargada de imprimir la informacion de cada peticion que recive el servidor

Function in charge of printing the information of each request that the server receives.
*/
func (m *goMiddleware) FormatRequest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		print("\n\n")
		log.Printf(color.Red(color.Bold("-- REQUEST INFORMATION  --")))
		log.Printf(color.BlueBg("Method: %s"), ctx.Request().Method)
		log.Printf(color.Cyan("URL: %s"), ctx.Request().URL)
		log.Printf(color.Magenta("Remote IP: %s"), ctx.RealIP())
		log.Printf(color.Bold("Host: %s"), ctx.Request().Host)

		//Iterate over all header fields

		log.Printf(color.Green("RemoteAddr= %q"), ctx.Request().RemoteAddr)
		return next(ctx)
	}

}

/*
Funcion encargada de iniciar los Middlewares

Function in charge of initiating the Middlewares

*/
func InitMiddleware() *goMiddleware {
	return &goMiddleware{}
}
