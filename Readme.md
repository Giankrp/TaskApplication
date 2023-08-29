
# Ejemplo de Uso de Gorilla Mux en Go

Este repositorio contiene un ejemplo simple de cómo crear un servidor web en Go utilizando el enrutador Gorilla Mux.

## Requisitos Previos

- Go (debe estar instalado en su sistema, consulte https://golang.org/doc/install)
- Git (opcional, pero recomendado para clonar este repositorio)
- postgreSQL o docker ya que se usa una base de datos para guardar los usuarios y las tareas relacionadas a ellos (hay que agregar las variables de entorno)
- Chi, un enrutador para generar un servidor facilmente
- Air (un live reload de go), ve a https://github.com/cosmtrek/air para saber sobre la instalacion

## Cómo Ejecutar

1. Clona este repositorio en tu máquina local (si no tienes Git, puedes descargar el código fuente en formato ZIP):

    ```bash
     https://github.com/Giankrp/TaskApplication.git
    ```

2. Cambia al directorio del proyecto:

    ```bash
    cd TaskApplication
    ```
4. Para generar una configuracion de air y poder ejecutar el live reload:

    ```bash
    air init
    ```


5. Ejecuta el programa:

    ```bash
    air
    ```

6. Abre tu navegador web y visita http://localhost:8000 para ver la aplicación en funcionamiento.

## En caso de usar docker

- En caso de usar docker compose ejecutar el siguiente comando:
  ```bash
    docker compose up -d
   ```
- En caso de usar docker:
 ```bash
    docker compose up -d
  ```
- Para hacer las consultas ejecutar el siguiente comando:
  ```bash
    docker exec -it some-postgres bash
   ```

## Detalles del Proyecto

- `main.go`: Este es el archivo principal que contiene la lógica para configurar y ejecutar el servidor web utilizando el enrutador Chi.

- `routes`: En esta carpeta se definen las rutas y se configuran los manejadores correspondientes.

- `db`: Aqui se define la conexion de la base de datos (postgreSQL)

- `models`: En esta carpte se definen los modelos de las tablas "tasks" y "users"


## Recursos Adicionales

- Documentación de [Chi](https://go-chi.io/#/)

- Documentación de [air]( https://github.com/cosmtrek/air) (live Reload utilizado)

- Documentación de  [Go](https://golang.org/doc/install)

- Documentación del contenedor de postgreSQL asi como su archivo docker-compose: [postgreSQL](https://hub.docker.com/_/postgres)

- [Docker](https://docs.docker.com/get-docker/)

- [Docker Compose](https://docs.docker.com/compose/install/)
