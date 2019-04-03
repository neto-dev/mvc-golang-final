# Implementación MVC + Go lang + Mysql + Echo

Nota: En el próximo release implementare los unit test utilizando Moks y entorno de desarrollo Test. Fecha aproximada 15/abr/2019

Ejemplo de la implementación de MVC en un servidor go + echo.

El ejemplo cuenta con archivos de configuración para bases de datos basado en el entorno que se este ejecutando.

Tecnologias implementadas:

- Go lang
- Gorm
- Mysql
- Echo

El actual ejemplo cuenta con una implementación básica de los métodos **Get/GetByID/Create/Update/Delete**, con un método Filters para realizar filtros.

Primero tendrás que clonar el repositorio a tu computadora con el siguiente comando dentro del directorio `$GOPATH/src/github.com/`.

`git clone https://github.com/neto-dev/mvc-golang-final.git`

Una vez descargado ingresar a la carpeta.

`cd mvc-golang-final`

Para instalar las diferentes dependencias que se utilizan en el proyecto ejecutar el siguiente comando.

`go mod init`

En caso de que arrojé el siguiente error.

`go: modules disabled inside GOPATH/src by GO111MODULE=auto; see 'go help modules'`

Ejecutar el siguiente comando.  

`export GO111MODULE=on`

Seguido de:

`go mod init`

Despues ejecutamos:

`go mod vendor`

Listo en este punto ya contaremos con las dependencias descargadas en el directorio vendor y con esto ya podremos ejecutar nuestro proyecto.


Seguido crearemos el archivo variables.env en el cual crearemos las siguientes variables de entorno.

```
export DATABASE_USER='root'
export DATABASE_PASS='pass_db'
export DATABASE_HOST=''
export ENVIRONMENT='Development'
```

Estas ultimas son las que el sistema utilizara para conectar con la base de datos y para setear el entorno de producción.

En seguida ejecutamos el comando:

`source ./variables.env`

Esto para setear nuestras variables de entorno que usaremos en el proyecto.

Teniendo una vez esto ya podremos ejecutar el proyecto con el comando.

`npm run main.go`

Y listo las migraciones y toda la configuración se ejecutara automáticamente al correr el servicio.

Es una implementación básica con la cual podrán desarrollar sus servidores o para basarse según sea el caso. 

### Autor
@neto-dev Ernesto Valenzuela, apasionado a la programación así como de las tecnologías open source, amante de los lenguajes Ruby, PHP, Python y Go, puedes contactarme directamente a correo hello@netodev.me o bien seguirme en twitter [@neto_dev](https://twitter.com/neto_dev)