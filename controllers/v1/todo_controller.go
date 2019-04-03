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

package v1

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/mvc-golang-final/models"
	//validator "gopkg.in/validator.v2"
)

type ControllerTodo struct {
	/*
		We inherited from the base structure

		Heredamos de la estructura base
	*/
	BaseController
	DB *gorm.DB
}

func (_controller_ ControllerTodo) Get(_ctx echo.Context) error {
	return Get(_ctx, _controller_.DB, &[]models.Todo{})
}

func (_controller_ ControllerTodo) GetByID(_ctx echo.Context) error {
	return GetByID(_ctx, _controller_.DB, &models.Todo{})
}

func (_controller_ ControllerTodo) Filters(_ctx echo.Context) error {
	return Filters(_ctx, _controller_.DB, &[]models.Todo{})
}

func (_controller_ ControllerTodo) Create(_ctx echo.Context) error {
	return Create(_ctx, _controller_.DB, &models.Todo{})
}

func (_controller_ ControllerTodo) Update(_ctx echo.Context) error {
	return Update(_ctx, _controller_.DB, &models.Todo{})
}

func (_controller_ ControllerTodo) Delete(_ctx echo.Context) error {
	return Delete(_ctx, _controller_.DB, &models.Todo{})
}

func NewTodoController(_e echo.Group, _DB *gorm.DB) echo.Group {
	/*
		Se declara la variable controller la cual contendra toda la estructura de nuestros controller methods.

		The controller variable is declared which will contain all the structure of our controller methods.
	*/
	controller := &ControllerTodo{
		DB: _DB,
	}

	/*
		Se declaran las rutas correspondientes a cada metodo

		The routes corresponding to each method are declared.
	*/

	_e.GET("/todos", controller.Get)
	_e.GET("/todos/:id", controller.GetByID)
	_e.POST("/todos/filters", controller.Filters)
	_e.POST("/todos", controller.Create)
	_e.PUT("/todos/:id", controller.Update)
	_e.PATCH("/todos/:id", controller.Update)
	_e.DELETE("/todos/:id", controller.Delete)

	return _e
}
