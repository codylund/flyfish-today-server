# [flyfish.today] server

[flyfish.today] is a customizable online dashboard for visualizing [USGS streamflow data]. 
This repository contains the backend component of the dashboard.
The frontend component, built with React, can be found [here][flyfish-today-client].

## Technology

The flyfish.today server is written in Golang and uses the [Gin] HTTP web framework.
The server offers a variety of REST APIs for managing user accounts. 
User data is stored in a MongoDB instance.

## Build

The server binary is built from the repository root using:

```bash
go build -o ./bin/ ./src/
```

## Test

The Docker Compose file in `/test` defines the following services:
- `server`: hosts the flyfish.today server on <http://localhost:8080>
- `mongo`: hosts a MongoDB instance on <http://localhost:27017>
- `mongo-express`: hosts an admin user interface for interacting with the MongoDB on <http://localhost:8081>

See [`/test/README.md`](./test/README.md) for complete details.

Together with the Dockerfile defined in the [client repo][flyfish-today-client],
these containers provide a local E2E test environment. 

## Deploy

The flyfish.today server is deployed as a standalone Docker container.
See the Dockerfile in the repository root.

This container must be deployed with the following environment variables:
- `MONGODB_URL`: the URL for the MongoDB instance.
- `ORIGIN_URL`: the expected origin URL for client requests (e.g. `https://flyfish.today`).
- `DOMAIN`: the client domain for setting cookies (e.g. `flyfish.today`).

## API reference

### Register new user 

`POST /v1/register`

Register a new flyfish.today user.

#### Request body

```json
{
  "username": "johnsmith@flyfish.today",
  "display_name": "John Smith", 
  "password": "abc123"
}
```

#### Response

On success (`200 OK`), registers the user and sets a fresh `session` cookie.

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Sign in existing user

`POST /v1/signin`

Authenticate an existing flyfish.today user,
setting a `session` cookie for future requests.

> [!NOTE]
> This API is designed for use by the [flyfish.today client][flyfish-today-client] in a web browser.
> Thus, session cookies are used for authentication in lieu of `Authorization` headers.

#### Request body

```json
{
  "username": "johnsmith@flyfish.today",
  "password": "abc123"
}
```

#### Response

On success (`200 OK`), sets a valid `session` cookie.

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Get user details

`GET /v1/me`

Get details about the authenticated user.

#### Response

On success (`200 OK`):

```json
{
  "username": "johnsmith@flyfish.today",
  "display_name": "John Smith", 
}
```

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Sign out curent user

`POST /v1/signout`

Sign out the authenticated user.

#### Response

On success (`200 OK`), the `session` cookie is cleared.

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Get sites for current user

`GET /v1/sites`

Returns the authenticated user's custom list of site entries.

#### Response Body

On success:

```json
[
  {
    "_id": "507f1f77bcf86cd799439011",
    "site_id": "06719505",
    "is_favorite": false,
    "tags": [
      "Front Range",
      "After Work",
      "Clear Creek"
    ]
  },
  {
    "_id": "64b7f171c1d509779fc0893f",
    "site_id": "06710605",
    "is_favorite": true,
    "tags": [
      "Front Range",
      "Bear Creek"
    ]
  }
]
```

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Add site

`POST /v1/sites/add`

Adds a USGS site entry for the authenticated user.

#### Request body

```json
{
  "site_id": "06719505",
  "is_favorite": false,
  "tags": [
    "Front Range",
    "After Work",
    "Clear Creek"
  ]
}
```

#### Response Body

On success, the response body echos the site entry.
It also includes a unique OID (see `_id`) for the USGS site entry.

```json
{
  "_id": "507f1f77bcf86cd799439011",
  "site_id": "06719505",
  "is_favorite": false,
  "tags": [
    "Front Range",
    "After Work",
    "Clear Creek"
  ]
}
```

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Update site

`PATCH /v1/sites/{id}`

Updates an existing site entry for the authenticated user.

| Parameter | Description |
|-----------|-------------|
| `id`      | The unique OID for the site entry. |

#### Request Body

The request body is a JSON blob of properties to patch on the site entry.

To update `tags` for a site:

```json
{
  "tags": [
    "Front Range",
    "After Work",
    "Clear Creek"
  ]
}
```

To update `is_favorite` for a site:

```json
{
  "is_favorite": false,
}
```

#### Response

On success (`200 OK`), the response body echos the complete updated site details.

```json
{
  "_id": "507f1f77bcf86cd799439011",
  "site_id": "06719505",
  "is_favorite": false,
  "tags": [
    "Front Range",
    "After Work",
    "Clear Creek"
  ]
}
```

On failure, returns an appropriate HTTP status code
and JSON body with additional details.
See [Errors by Status Code](#errors-by-status-code) for more info.

### Remove site

`DELETE /v1/sites/{id}`

Removes an existing site entry for the authenticated user.

| Parameter | Description |
|-----------|-------------|
| `id`      | The unique OID for the site entry. |

### Errors by status code

Responses returned by this server follow standard HTTP status codes.

The response body may contain additional information about the error,
formatted as follows:

```json
{
  "message": "An account with username johnsmith@flyfish.today already exists."
}
```

[flyfish.today]: https://flyfish.today
[flyfish-today-client]: https://github.com/codylund/flyfish-today-client

[Gin]: https://gin-gonic.com/docs/
[USGS streamflow data]: https://waterdata.usgs.gov/nwis/rt