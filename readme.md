## Klare

## requerimientos

- docker >= 20
- docker-compose >= 1.27

## instrucciones para levantar

correr:
docker-compose up


## Uso

el proyecto se levanta en: 'http://localhost:8080'
respondiendo a los siguientes endpoints:

GET     "/api/v1/products"
GET     "/api/v1/products/:id"
POST    "/api/v1/products", products.AddProduct)
DELETE  "/api/v1/products/:id"

 para obtener un json con todos los productos hacer solicitud a:
 http://localhost:8080/api/v1/products
