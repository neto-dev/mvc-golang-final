package db

/*
Importamos las librerias

We import libraries
*/
import "github.com/mvc-golang-final/models"

/*
Declaramos la estructura Relationship que se encargara de guardar los datos referentes a las relaciones con las que cuenta cada entridad

We declare the Relationship structure that will be in charge of storing the data referring to the relationships with which each entry counts.
*/

type Relationship struct {
	Field    string
	Dest     string
	OnDelete string
	OnUpdate string
}

/*
Declaramos la estructura Entity que se encargara de guardar los datos referentes a las entidades y sus relaciones

We declare the Entity structure that will be in charge of storing the data relating to the entities and their relationships. */

type Entity struct {
	Model         interface{}
	Relationships []Relationship
}

/*
Declaramos la estructura Entitys la cual sera un array de Entity

We declare the structure Entitys which will be an array of Entity
*/
type Entitys []Entity

/*
Declaramos la variable entitys la cual sera un objeto de Entitys y a su vez le asignamos todos los valores correspondientes

We declare the variable entitys which will be an object of Entitys and at the same time we asicnamos all the corresponding values.
*/
var entitys = Entitys{
	Entity{
		models.Todo{},
		nil,
		//[]Relationship{
		//	{
		//		"todo_relation_id",
		//		"todo_relations(id)",
		//		"RESTRICT",
		//		"RESTRICT",
		//	},
		//},
	},
}
