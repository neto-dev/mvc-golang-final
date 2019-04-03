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
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	//"fmt"
	"github.com/mvc-golang-final/utils"
	//"net/http"
	//"strconv"
	//"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"gopkg.in/validator.v2"
	"net/http"
	"strconv"
)


type BaseController struct {
	ctx   echo.Context
	//Utils utils.Base
}

//
func Get(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {
	/*
		We get list.

		Obtenemos los registros
	*/
	if err := _db.Find(_model).Order("id asc").Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Records Find Failure. Please inform your service representative about this error.", err)
	}
	/*
		We return the values

		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}

func GetByID(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {

	/*
		Filter within the database searching for the record corresponding
		to the received id. The preload works to bring up the relationship
		data

		Filtramos dentro de la base de datos buscando el registro que
		corresponda con el id recibido. El preload funciona para traer
		los datos de la relacion
	*/

	idP, err := strconv.Atoi(_ctx.Param("id"))

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	if err := _db.First(_model, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record not found", err)
	}

	/*
		We return the values

		Retornamos los valores
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}

func Filters(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {
	/*
		We sent to call the NewFilter method to get the
		results according to the parameters we received
		in the body of the request, in case of error return
		a query error returned error 500 with the message
		Record Find Failure so that the error can be handled
		in the client. We send the corresponding model to the
		controller that is being consulted within the parameters
		of the method.

		Mandamos a llamar el metodo NewFilter para porder
		obtener los resultados segun los parametros que
		recivimos en el body del request, en caso de que
		retorne un error la consulta retornamos error 500
		con el mensaje Record Find Failure para que se pueda
		manejar el error en el cliente. Enviamos el modelo
		correspondiente al controlador que se esta consultando
		dentro de los parametros del metodo.
	*/

	if err := utils.NewFilter(_ctx, _model, _db); err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record find Failure", err)
	}

	/*
		Return values in json format

		Retornamos los valores en formato json
	*/
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}

func Create(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {
	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _db.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Error.Error(), tx.Error)
	}

	/*
		We recover the parameters we received in the request

		Recuperamos los parametros que recibimos en el request
	*/
	if err := _ctx.Bind(_model); err != nil {
		fmt.Println(err)
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure

		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(_model); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We created the new record in the database

		Creamos el nuevo registro en la base de datos
	*/
	if err := tx.Create(_model).Error; err != nil {
		tx.Rollback()
		myerr, ok := err.(*mysql.MySQLError)
		if !ok {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}
		if myerr.Number == 1062 {
			return utils.ReturnErrorJSON(_ctx, "Duplicate Record", err)
		} else {
			return utils.ReturnErrorJSON(_ctx, "Record Create Failure", err)
		}
	}
	/*
		We return the values

		Retornamos los valores
	*/
	if err := tx.Commit().Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Commit().Error.Error(), tx.Commit().Error)
	}
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}

func Update(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {

	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _db.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Error.Error(), tx.Error)
	}


	idP, err := strconv.Atoi(_ctx.Param("id"))

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Recover the record you wish to edit

		Recuperamos el registro que se desea editar
	*/
	if err := tx.First(_model, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record Find Failure", err)
	}
	/*
		We recover the parameters we received in the request

		Recuperamos los parametros que recibimos en el request
	*/

	if err := _ctx.Bind(_model); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		We validate that the parameters have the correct structure

		Validamos que los parametros tengan la estructura correcta
	*/

	if err := validator.Validate(_model); err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}

	/*
		We save the edited information in the database

		Guardamos la información editada en la base de datos
	*/

	if err := tx.Save(_model).Error; err != nil {
		tx.Rollback()
		myerr, ok := err.(*mysql.MySQLError)
		if !ok {
			return utils.ReturnErrorJSON(_ctx, "Record Update Failure", err)
		}
		if myerr.Number == 1062 {
			return utils.ReturnErrorJSON(_ctx, "Duplicate Record", err)
		} else {
			return utils.ReturnErrorJSON(_ctx, "Record Update Failure", err)
		}
	}

	/*
		We return the values

		Retornamos los valores
	*/
	if err := tx.Commit().Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Commit().Error.Error(), tx.Commit().Error)
	}
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}

func Delete(_ctx echo.Context, _db *gorm.DB, _model interface{}) error {

	/*
		Inicializamos la variable tx la cual almacenara una transaccion de la base de datos

		We initialize the tx variable which will store a database transaction.
	*/
	tx := _db.Begin()
	// Note the use of tx as the database handle once you are within a transaction

	/*
		Validamos que se haya asignado correctamente la transaccion

		We validate that the transaction has been assigned correctly
	*/
	if tx.Error != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Error.Error(), tx.Error)
	}

	idP, err := strconv.Atoi(_ctx.Param("id"))

	if err != nil {
		return utils.ReturnErrorJSON(_ctx, err.Error(), err)
	}
	/*
		Seteamos la variable idP en la variable id transformandola en uint

		We set the variable idP in the variable id transforming it into uint
	*/
	id := uint(idP)

	/*
		Recover the record you wish to delete

		Recuperamos el registro que se desea eliminar
	*/
	if err := tx.First(_model, id).Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, "Record Find Failure", err)
	}


	if err := tx.Delete(_model).Error; err != nil {
		tx.Rollback()
		return utils.ReturnErrorJSON(_ctx, "Record Delete Failure", err)
	}
	/*
		We return a message to inform you that the registration has been deleted

		Retornamos un mensaje para informar que el registro ha sido eliminado
	*/

	if err := tx.Commit().Error; err != nil {
		return utils.ReturnErrorJSON(_ctx, tx.Commit().Error.Error(), tx.Commit().Error)
	}
	return utils.ReturnRespondJSON(_ctx, http.StatusOK, _model)
}
