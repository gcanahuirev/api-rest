# Simple api rest in memory

Simple test of api rest without dependencies, but in memory

- No third-party `packages/dependencies`
- You can use curl or a `http`  client for requests

## Requirements

- Go stable version

## Features

- [x] `GET /shirts` returns list of shirts as JSON
- [x] `GET /shirts/{id}` returns details of specific shirt as JSON
- [x] `POST /shirts` accepts a new shirt to be added
- [x] `POST /shirts` returns status 415 if content is not `application/json`
- [x] `GET /admin` requires basic auth
- [x] `GET /shirts/random` redirects (Status 302) to a random shirt

### Data Types

A shirt object should look like this:

```json
{
  "id" : "0001",
  "material" : "Lana",
  "class" : "Manga Larga",
  "size" : 14,
}
```

## Usage

Note: If you are using pwsh, please note this

- First configure the password, then run `go run server.go`

``` sh
> $env:ADMIN_PASSWORD = 'secret'
```

- Or just run

``` sh
> ADMIN_PASSWORD='secret' go run server.go
```

- To test basic authentication

``` sh
> curl localhost:3000/admin -u admin:secret
```

- Remember that you will have to add new data at each restart

### Persistence

There is no persistence of data, a memory time is fine.
