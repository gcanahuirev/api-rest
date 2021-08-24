# api-rest
> Ejemplo simple de api rest sin dependencias

* No third-party **packages/dependencies**
* Puede utilizar curl o algún cliente **http** para las consultas

## Requerimientos
Esta API REST debe coincidir con algunos requisitos:

* [x] `GET /shirts` returns list of shirts as JSON
* [x] `GET /shirts/{id}` returns details of specific coaster as JSON
* [x] `POST /shirts` accepts a new shirt to be added
* [x] `POST /shirts` returns status 415 if content is not `application/json`
* [x] `GET /admin` requires basic auth
* [x] `GET /shirts/random` redirects (Status 302) to a random shirt

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

Nota: Si esta usando pwsh, tenga esto en cuenta 
- Para `GET /admin`, primero deberá configurar la clave
``` pwsh
> $env:ADMIN_PASSWORD = 'secret'
```
- Para probar el authbasic
``` pwsh
> curl localhost:3000/admin -u admin:secret
```
- Recuerde que tendra que agregar nuevos datos en cada reinicio
### Persistencia
No hay persistencia de datos, un tiempo en memoria está bien.