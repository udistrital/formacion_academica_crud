# formacion_academica_crud
API de registro de formación académica

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `formacion_academica_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/udistrital/formacion_academica_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `FORMACION_ACADEMICA_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `FORMACION_ACADEMICA_CRUD__PGUSER`: Usuario de la base de datos
 - `FORMACION_ACADEMICA_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `FORMACION_ACADEMICA_CRUD__PGURLS`: Host de conexión
 - `FORMACION_ACADEMICA_CRUD__PGDB`: Nombre de la base de datos
 - `FORMACION_ACADEMICA_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
FORMACION_ACADEMICA_HTTP_PORT=8098 FORMACION_ACADEMICA_CRUD__PGUSER=user FORMACION_ACADEMICA_CRUD__PGPASS=password FORMACION_ACADEMICA_CRUD__PGURLS=localhost FORMACION_ACADEMICA_CRUD__PGDB=bd FORMACION_ACADEMICA_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![modelo_formacion_academica_crud_git](https://user-images.githubusercontent.com/14035745/66046769-0e9b8a00-e4ec-11e9-8639-c4ec76377fa8.png)
