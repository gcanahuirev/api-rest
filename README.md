# api-rest
> Ejemplo simple de api rest sin dependencias

* No third-party **packages/dependencies**
* Puede utilizar curl o algún cliente **http** 

## Requerimientos
Esta API REST debe coincidir con algunas
requisitos:

* [x] `GET /shirts` returns list of shirts as JSON
* [x] `GET /shirts/{id}` returns details of specific coaster as JSON
* [x] `POST /shirts` accepts a new shirt to be added
* [x] `POST /shirts` returns status 415 if content is not `application/json`
* [ ] `GET /admin` requires basic auth
* [ ] `GET /shirts/random` redirects (Status 302) to a random shirt

### Data Types
Un objeto de camisa debería verse así:
```json
{
  "class" : "Manga Larga",
  "material" : "Lana",
  "id" : "0001",
  "size" : 14,
}
```

### Persistencia
No hay persistencia, un tiempo en memoria está bien.