# Price o' Matic

En Klare necesitamos añadir un nuevo microservicio que se encargue de listar los precios de cada uno de nuestos productos.
Este debe ser simple y debe considerar lo siguiente:

* Debe ser escrito en Go o Python
* Debe documentar la solución propuesta, considerar lo siguiente:
  * Debe incluir lo necesario para ejecutar y verificar los requerimientos.
  * Debe incluir un TODO, que indique que aspectos cree que puede mejorar en el futuro de la solución entregada.
* Debe modelar una tabla `Product` con los siguientes atributos:
  * Id
  * Fecha de inserción
  * Nombre
  * Precio
* Debe crear los siguientes endpoints:
    * GET `/api/v1/products` <-- Lista todos los productos
    * GET `/api/v1/products/:id` <-- Obtiene el producto :id
    * POST `/api/vi/products` <-- Inserta un producto en la base de datos
    * [BONUS] `PATCH /api/vi/products/:id` <-- Actualiza uno o mas campos del producto :id
    * [BONUS] `DELETE /api/vi/products/:id` <-- Elimina el producto :id
* Debe escribir un script que permita insertar el json incluido (`data.json`) a la base de datos, puede utilizar lo que considere necesario.
* [BONUS] Crear una ruta `GET /`, que retorne una página simple que los productos existentes en la base de datos, con todos los campos mencionados anteriormente.

## Notas

* Dudas y/o consultas por correo a ernesto.olivares@klare.cl o al whatsapp +569 9933 5874.
* Revisar los siguientes conceptos de Go: structs, structs tags, error handling.
* Se sugiere utilizar GORM (https://gorm.io/index.html) como ORM para manejar el acceso a la base de datos, pero puede ocupar lo que considere.
* El proyecto usa fiber (https://github.com/gofiber/fiber) como framework web, se recomienda revisar los ejemplos básicos.
* Si elige usar Python
  * Se recomienda utilizar Fastapi.
  * Debe reescribir el Dockerfile y añadir lo necesario para que levante el contenedor.
* Puede utilizar la base de datos que considere, el proyecto viene con un setup por defecto de `postgres`.
* El proyecto se probará en entorno dockerizado, por ende, debe levantar correctamente usando el comando `docker-compose up`.
* Una vez levantado el proyecto, se ejecutará el script para inicializar la base de datos con los datos provistos.
* [Bonus] Se permite hardcodear los strings de configuración (base de datos, etc), pero sumará puntos si carga las variables de entorno desde el archivo local.env.
  * Utilizar `postgres:5432` como dns de conexión al utilizar docker-compose

## A Evaluar

* Plazo de entrega (Avisar si requiere más tiempo)
* Correctitud de lo entregado.
* Planteamiento de la solución.
* Orden y estructura del código.
* Documentación entregada.
* Bonus.