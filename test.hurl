# List shirts (the store starts empty after each server restart)
GET {{host}}/shirts
HTTP 200
[Asserts]
jsonpath "$" isCollection

# An unsupported content type is rejected
POST {{host}}/shirts
Content-Type: text/plain
{
  "material": "Algodón",
  "class": "Manga Corta",
  "size": 12
}
HTTP 415

# Create a new shirt
POST {{host}}/shirts
Content-Type: application/json
{
  "size": 14,
  "material": "Lana",
  "class": "Manga Larga"
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
[Captures]
id: jsonpath "$.id"

# Get details of specific shirt
GET {{host}}/shirts/{{id}}
HTTP 200
[Asserts]
jsonpath "$.id" exists

# A missing shirt returns not found
GET {{host}}/shirts/0
HTTP 404

# A random shirt redirects to one of the stored shirts
GET {{host}}/shirts/random
HTTP 302
[Asserts]
header "location" exists

# Invalid admin credentials are rejected
GET {{host}}/admin
[BasicAuth]
admin: none
HTTP 401

# The admin portal requires valid basic authentication.
# Start the server with ADMIN_PASSWORD=secret for this request.
GET {{host}}/admin
[BasicAuth]
admin: secret
HTTP 200
[Asserts]
body contains "secret admin portal"
