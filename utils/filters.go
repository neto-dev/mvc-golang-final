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

package utils

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
)

type Params struct {
	Where   string `json:"where"`
	Order   string `json:"order"`
	Limit   int    `json:"limit"`
	Page    int    `json:"page"`
	Select  string `json:"select"`
	Join    string `json:"join"`
	Include []struct {
		Schema string `json:"schema"`
		Where  string `json:"where"`
	} `json:"include"`
}

func NewFilter(ctx echo.Context, i interface{}, DB *gorm.DB) error {
	/*
		We assign the Params structure to the params variable

		Asignamos la estructura de Params a la variable params
	*/
	params := &Params{}

	// /*
	// 	We obtain the parameters and parse them in JSON
	// 	format to assign it later to the Params structure.

	// 	Obtenemos los parametros y los parseamos en formato
	// 	JSON para asignarlo posteriormente a la estructura Params
	// */
	// decoder := json.NewDecoder(c.Request.GetBody())

	// /*
	// 	We insert the values obtained from the body of the request
	// 	in the structure

	// 	Insertamos los valores obtenidos desde el body del request
	// 	en la estructura
	// */
	// err := decoder.Decode(&params)

	if err := ctx.Bind(params); err != nil {
		log.Println(err.Error())
		return errors.New("An error has occurred while processing the received data")
	}

	/*
		If within the parameters of the request the Query
		parameter is different from null it indicates that
		we will filter the results depending on the Query
		obtained.

		Si dentro de los parametros del request el parametro
		Query es diferente a null nos indica que filtraremos
		los resultados dependiendo el Query obtenido
	*/
	if params.Where != "" {
		/*
			If within the parameters of the request the
			Order parameter is different from null it
			indicates that we will order the results
			according to the obtained request.

			Si dentro de los parámetros del request el
			parámetro Order es diferente a null nos
			indica que ordenaremos los resultados según
			la petición obtenida
		*/
		if params.Order != "" {
			/*
				If within the parameters of the request
				the parameter Limit is different from null
				it indicates how many are the results that
				will be obtained according to the obtained
				request.

				Si dentro de los parámetros del request el
				parámetro Limit es diferente a null nos
				indica cuantos son los resultados que se
				van a obtener según la petición obtenida
			*/
			if params.Limit != 0 {
				/*
					If within the parameters of the request
					the Page parameter is different from null,
					it indicates that pagination is being used
					and the information will be returned according
					to the page number you want to obtain.

					Si dentro de los parámetros del request el
					parámetro Page es diferente a null nos indica
					que se esta usando paginación y se devolverá
					la información según el numero de pagina que
					quiera obtener
				*/
				if params.Page != 0 {
					if params.Page == 1 {
						/*
							If the value of Page is equal to 1 we
							will return the values from register 0
							of the query depending on the received
							limit, at this point all the filters are
							carried out since it is detected that it
							complies with all the parameters to arrive
							at this point, the results will be returned
							according to the parameters received from
							the request.

							Si el valor de Page es igual a 1 retornaremos
							los valores a partir del registro 0 de la
							consulta dependiendo del limit recibido,
							en este punto se realizan todos los filtros
							ya que se detecto que cumple con todos los
							parámetros para llegar a este punto, se
							retornara los resultados según los parámetros
							recibidos desde el request.
						*/

						if len(params.Include) > 0 {

							data := DB
							if params.Join != "" {
								data.Joins(params.Join)
							}
							for field := 0; field < len(params.Include); field++ {
								data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i)
							}
							return data.Error
						}
						if params.Join != "" {
							return DB.Joins(params.Join).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
						}
						return DB.Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
					}
					/*
						If the value of Page is equal to 2 we
						will return the values from register 0
						of the query depending on the received
						limit, at this point all the filters are
						carried out since it is detected that it
						complies with all the parameters to arrive
						at this point, the results will be returned
						according to the parameters received from
						the request.

						Si el valor de Page es igual a 2 retornaremos
						los valores a partir del registro 0 de la
						consulta dependiendo del limit recibido,
						en este punto se realizan todos los filtros
						ya que se detecto que cumple con todos los
						parámetros para llegar a este punto, se
						retornara los resultados según los parámetros
						recibidos desde el request.
					*/
					if params.Page >= 2 {
						/*
							We assign in variable the values of Page
							and Limit

							Asignamos en variable los valores de
							Page y Limit
						*/
						page := params.Page
						limit := params.Limit
						/*
							We obtain the offset based on the page
							value and limit the way in which the
							correct registers are obtained is by
							subtracting 1 from the value received
							from Page params and then multiplying
							it by the value of Limit params.
							Example:
							page: = 2
							limit: =10
							offset = (page - 1) * (limit) This returns 1

							Obtenemos el offset en base a el valor
							de page y de limit la manera en como
							se obtienen los registros correctos es
							restando 1 al valor que se recibe de
							params Page y posteriormente
							multiplicándolo por el valor de params
							Limit.
							Ejemplo:
							page := 2
							limit :=10
							offset = (page - 1) * (limit) Esto retorna 10 0
						*/
						offset := (page - 1) * (limit)
						/*
							We return the result

							Retornamos el resultado
						*/

						if len(params.Include) > 0 {

							data := DB
							if params.Join != "" {
								data.Joins(params.Join)
							}
							for field := 0; field < len(params.Include); field++ {
								data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i)
							}
							return data.Error
						}
						if params.Join != "" {
							return DB.Joins(params.Join).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
						}
						return DB.Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
					}
				}
				/*
					In case Page comes empty we return the result
					without applying that filter.

					En caso de que Page venga vacío retornamos el
					resultado sin aplicar ese filtro
				*/

				if len(params.Include) > 0 {

					data := DB
					if params.Join != "" {
						data.Joins(params.Join)
					}
					for field := 0; field < len(params.Include); field++ {
						data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i)
					}
					return data.Error
				}
				if params.Join != "" {
					return DB.Joins(params.Join).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
				}
				return DB.Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
			}

			if len(params.Include) > 0 {

				data := DB
				if params.Join != "" {
					data.Joins(params.Join)
				}
				for field := 0; field < len(params.Include); field++ {
					data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Order(params.Order).Where(params.Where).Find(i)
				}
				return data.Error
			}
			if params.Join != "" {
				return DB.Joins(params.Join).Order(params.Order).Where(params.Where).Find(i).Error
			}
			return DB.Order(params.Order).Where(params.Where).Find(i).Error
		}

		/*
			Other conditions interact in the same way validate
			the data being received and return a response as
			the case may be.

			El Resto de condiciones interactúan de la misma
			forma validan los datos que se están recibiendo
			y retorna una respuesta según sea el caso.
		*/

		if params.Limit != 0 {
			if params.Page != 0 {
				if params.Page == 1 {
					if len(params.Include) > 0 {

						data := DB
						if params.Join != "" {
							data.Joins(params.Join)
						}
						for field := 0; field < len(params.Include); field++ {
							data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i)
						}
						return data.Error
					}
					if params.Join != "" {
						return DB.Joins(params.Join).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
					}
					return DB.Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
				}
				if params.Page >= 2 {
					page := params.Page
					limit := params.Limit
					offset := (page - 1) * (limit)
					if len(params.Include) > 0 {

						data := DB
						if params.Join != "" {
							data.Joins(params.Join)
						}
						for field := 0; field < len(params.Include); field++ {
							data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i)
						}
						return data.Error
					}
					if params.Join != "" {
						return DB.Joins(params.Join).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
					}
					return DB.Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
				}
			}
			if len(params.Include) > 0 {

				data := DB
				if params.Join != "" {
					data.Joins(params.Join)
				}
				for field := 0; field < len(params.Include); field++ {
					data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Where(params.Where).Find(i)
				}
				return data.Error
			}
			if params.Join != "" {
				return DB.Joins(params.Join).Limit(params.Limit).Where(params.Where).Find(i).Error
			}
			return DB.Limit(params.Limit).Where(params.Where).Find(i).Error
		}

		if len(params.Include) > 0 {

			data := DB
			if params.Join != "" {
				data.Joins(params.Join)
			}
			for field := 0; field < len(params.Include); field++ {
				data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Where(params.Where).Find(i)
			}
			return data.Error
		}

		if params.Join != "" {
			return DB.Joins(params.Join).Where(params.Where).Find(i).Error
		}

		return DB.Where(params.Where).Find(i).Error
	}

	/*
		The page parameter is only taken into account when
		the limit parameter is also received.

		El parametro page solo se toma en cuenta cuando
		también se recibe el parámetro de limit
	*/

	if params.Order != "" {
		if params.Limit != 0 {
			if params.Page != 0 {
				if params.Page == 1 {
					if len(params.Include) > 0 {

						data := DB
						if params.Join != "" {
							data.Joins(params.Join)
						}
						for field := 0; field < len(params.Include); field++ {
							data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i)
						}
						return data.Error
					}
					if params.Join != "" {
						return DB.Joins(params.Join).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
					}
					return DB.Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
				}
				if params.Page >= 2 {
					page := params.Page
					limit := params.Limit
					offset := (page - 1) * (limit)
					if len(params.Include) > 0 {

						data := DB
						if params.Join != "" {
							data.Joins(params.Join)
						}
						for field := 0; field < len(params.Include); field++ {
							data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i)
						}
						return data.Error
					}
					if params.Join != "" {
						return DB.Joins(params.Join).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
					}
					return DB.Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
				}
			}
			if len(params.Include) > 0 {

				data := DB
				if params.Join != "" {
					data.Joins(params.Join)
				}
				for field := 0; field < len(params.Include); field++ {
					data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Order(params.Order).Find(i)
				}
				return data.Error
			}
			if params.Join != "" {
				return DB.Joins(params.Join).Limit(params.Limit).Order(params.Order).Find(i).Error
			}
			return DB.Limit(params.Limit).Order(params.Order).Find(i).Error
		}
		if len(params.Include) > 0 {

			data := DB
			if params.Join != "" {
				data.Joins(params.Join)
			}
			for field := 0; field < len(params.Include); field++ {
				data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Order(params.Order).Find(i)
			}
			return data.Error
		}
		if params.Join != "" {
			return DB.Joins(params.Join).Order(params.Order).Find(i).Error
		}
		return DB.Order(params.Order).Find(i).Error
	}

	if params.Limit != 0 {
		if params.Page != 0 {
			if params.Page == 1 {
				if len(params.Include) > 0 {

					data := DB
					if params.Join != "" {
						data.Joins(params.Join)
					}
					for field := 0; field < len(params.Include); field++ {
						data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i)
					}
					return data.Error
				}
				if params.Join != "" {
					return DB.Joins(params.Join).Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
				}
				return DB.Limit(params.Limit).Offset(0).Order(params.Order).Where(params.Where).Find(i).Error
			}
			if params.Page >= 2 {
				page := params.Page
				limit := params.Limit
				offset := (page - 1) * (limit)
				if len(params.Include) > 0 {

					data := DB
					if params.Join != "" {
						data.Joins(params.Join)
					}
					for field := 0; field < len(params.Include); field++ {
						data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i)
					}
					return data.Error
				}
				if params.Join != "" {
					return DB.Joins(params.Join).Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
				}
				return DB.Offset(offset).Limit(params.Limit).Order(params.Order).Where(params.Where).Find(i).Error
			}
		}
		if len(params.Include) > 0 {

			data := DB
			for field := 0; field < len(params.Include); field++ {
				data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Limit(params.Limit).Find(i)
			}
			return data.Error
		}
		return DB.Limit(params.Limit).Find(i).Error
	}

	if len(params.Include) > 0 {

		data := DB
		for field := 0; field < len(params.Include); field++ {
			data = data.Preload(params.Include[field].Schema, params.Include[field].Where).Find(i)
		}
		return data.Error
	}

	if params.Select != "" {
		return DB.Select(params.Select).Find(i).Error
	}

	if params.Join != "" {
		return DB.Joins(params.Join).Find(i).Error
	}

	/*
		If no data is received some returns all records.

		Si no se reciben datos algunos retorna todos los registros.
	*/

	result := DB
	if params.Where != "" {
		result = result.Where(params.Where)
	}
	if params.Order != "" {
		result = result.Order(params.Order)
	}
	if params.Limit != 0 {
		result = result.Limit(params.Limit)
	}
	if params.Select != "" {
		result = result.Select(params.Select)
	}
	if params.Join != "" {
		result = result.Joins(params.Join)
	}
	if len(params.Include) > 0 {
		for field := 0; field < len(params.Include); field++ {
			result = result.Preload(params.Include[field].Schema, params.Include[field].Where)
		}
	}
	if params.Page != 0 {
		offset := (params.Page - 1) * (params.Limit)
		result = result.Offset(offset)
	}

	return result.Find(i).Error
}
