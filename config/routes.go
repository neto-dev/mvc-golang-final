/*
 * Developer: Ernesto Valenzuela Vargas.
 * Created by: netodev
 * Contact: neto.dev@protonmail.com
 *
 */

/*
 * Revision History:
 *     Initial:      10/18/18  |  1:56 PM     Ernesto Valenzuela V
 */

package config

/*
Importamos las librerias

We import libraries
*/
import (
	"github.com/mvc-golang-final/controllers/v1"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

/*
Metodo encargado de crear las rutas de cada Controlador y de inyectar las dependencias de cada capa perteneciente a cada
entidad

Method in charge of creating the routes of each Controller and injecting the dependencies of each layer belonging to each Handler.
entity
*/
func Routes(DB *gorm.DB, group echo.Group) {
	/*
		Pasamos al controlador de la entidad el grupo de las rutas y le inyectamos la configuracion de base de datos

		We pass to the heandler of the entity the group of the routes and we inject the database configuration
	*/
	v1.NewTodoController(group, DB)

}
