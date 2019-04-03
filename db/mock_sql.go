/*
 * Created Ernesto Valenzuela Vargas.
 *
 * Contact: neto.dev@protonmail.com
 *
 */

package db

/*
Importamos las librerias

We import libraries
*/
import (
	"log"

	"github.com/mvc-golang-final/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*
La funcion MockSql es la encargada de generar la base de datos que se utilizara para los unit test dentro de la
aplicacion.

The MockSql function is in charge of generating the database that will be used for the unit tests within the
application.
*/
func MockSql() *gorm.DB {

	configuration := config.Environment["Test"]

	/*
		Declaramos la variable db la cual almacenara la instancia a la base de datos y la variable err que almacenara el
		error en caso de que sugiera alguno

		We declared the variable db which would store the instance to the database and the err variable which would store
		the error in case you suggest any
	*/
	db, err := gorm.Open(configuration.DataBase.Dialect, configuration.DataBase.User+":"+
		configuration.DataBase.Pass+"@tcp("+configuration.DataBase.
		DbHost+")/?charset=utf8&parseTime=True&loc=Local")

	if err := db.Exec("CREATE DATABASE IF NOT EXISTS " + configuration.DataBase.DataBase).Error; err != nil {
		log.Panicf("Failed to create to database: %v\n", err)
	}

	/*
		Asignamos la base de datos a usar

		We assign the database to use
	*/
	if err := db.Exec("USE " + configuration.DataBase.DataBase).Error; err != nil {
		log.Panicf("Failed to assign database: %v\n", err)
	}

	/*
		Validamos si al querer inicializar la base de datos retorna un error.

		Let's validate if when wanting to initialize the database returns an error.
	*/
	if err != nil {
		log.Panicf("Failed to connect to database: %v\n", err)
	}

	/*
		Activamos el log mode de GORM para imprimir en consola las consultas realizadas.

		Activate the GORM log mode to print the queries in the console.
	*/
	db.LogMode(configuration.Debug)

	/*
		DB obtiene `*sql.DB` de la conexión actual

		DB get `*sql.DB` from current connection
	*/
	db.DB()

	/*
		Declaramos las propiedades de la base de datos.

		We declare the properties of the database.
	*/
	db.Set("gorm:table_options", "ENGINE=InnoDB")

	/*
		Iteramos el array entitys

		We enter the array entities
	*/
	for _, entity := range entitys {
		/*
			Generamos la migracion a la base de batos

			We generate the migration to the base of batos
		*/
		db.AutoMigrate(entity.Model)
		for _, relationship := range entity.Relationships {
			/*
				Generamos las relaciones de cada entidad

				We generate the relationships of each entity
			*/
			db.Model(entity.Model).AddForeignKey(relationship.Field, relationship.Dest, relationship.OnDelete, relationship.OnUpdate)
		}
	}
	//for _, procedure := range procedures {
	//	db.Exec(`
	//		DELIMITER //
	//			CREATE PROCEDURE ` + procedure.Name + `(IN ` + procedure.Params + `)
	//			BEGIN ` + procedure.Query + ` END
	//		//
	//		DELIMITER ;
	//	`)
	//}

	return db

}
