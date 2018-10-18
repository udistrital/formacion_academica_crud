# formacion_academica_crud

--Api de formación academica con CI--
CI deploy with lambda - S3
Drone 0.8 
formacion_academica_crud master/develop

## Requirements
Go version >= 1.8.

## Preparación:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/formacion_academica_crud

## Run

Definir los valores de las siguientes variables de entorno:

 - `FORMACION_ACADEMICA_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `FORMACION_ACADEMICA_CRUD__PGUSER`: Usuario de la base de datos
 - `FORMACION_ACADEMICA_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `FORMACION_ACADEMICA_CRUD__PGURLS`: Host de conexión
 - `FORMACION_ACADEMICA_CRUD__PGDB`: Nombre de la base de datos
 - `FORMACION_ACADEMICA_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

Ejemplo: ORMACION_ACADEMICA_HTTP_PORT=8095 FORMACION_ACADEMICA_CRUD__PGUSER=postgres FORMACION_ACADEMICA_CRUD__PGPASS=password FORMACION_ACADEMICA_CRUD__PGURLS=localhost FORMACION_ACADEMICA_CRUD__PGDB=local FORMACION_ACADEMICA_CRUD__SCHEMA=core_new bee run

## MODELO
![image](https://user-images.githubusercontent.com/14035745/42359402-4cd653f4-80a7-11e8-8b90-61e30a20bbaf.png).
