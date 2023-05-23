# Member-Crud API

Member CRUD Service API with search capabilities


## Table of Contents
- [Member-Crud API](#member-crud-api)
    - [Search members](#search-members)
    - [Health endpoint](#health-endpoint)
    - [Ping endpoint](#ping-endpoint)
    - [Get all members](#get-all-members)
    - [Add a new member](#add-a-new-member)
    - [Get a member by ID](#get-a-member-by-id)
    - [Delete a member](#delete-a-member)
    - [Update a member by ID](#update-a-member-by-id)
    - [Get all tags](#get-all-tags)
    - [Definitions](#definitions)
        - [models.Member](#modelsmember)
        - [models.Tag](#modelstag)
## Search members

Searches members based on tags and member type.

- Path: `/api/v1/members/search`
- Method: `GET`
- Produces: `application/json`
- Tags: `members`

### Query Parameters

- `tags` (optional): Tags to search for. Multiple values can be provided.
- `type` (optional): Member type to search for.

### Responses

- `200 OK`: Returns an array of member objects matching the search criteria.
#### Example:
```
/api/v1/members/search?tags=Java&tags=Python&type=EMPLOYEE
```
```json
[
  {
  "name": "John",
  "tags": ["Java","Python"],
  "contractDuration": 12,
  "role": "Developer",
  "type": "EMPLOYEE"
  }
]
```

## Health endpoint

Checks the health status of the server and database connection.

- Path: `/api/health`
- Method: `GET`
- Produces: `text/plain`
- Tags: `health`

### Responses

- `200 OK`: The health check is successful.
- `500`: Could not establish a connection to the database.

## Ping endpoint

Checks the availability of the server.

- Path: `/api/ping`
- Method: `GET`
- Produces: `text/plain`
- Tags: `ping`

### Responses

- `200 OK`: The server is available.

## Get all members

Gets all members with their associated tags.

- Path: `/api/v1/members`
- Method: `GET`
- Produces: `application/json`
- Tags: `members`

#### Responses

- `200 OK`: Returns an array of member objects with tags.

#### Example:
```json
[
  {
    "id": 1,
    "name": "John",
    "tags": ["java"],
    "createdAt": "2022-01-01T00:00:00Z",
    "deletedAt": null,
    "updatedAt": "2022-01-01T00:00:00Z",
    "role": "Developer",
    "type": "EMPLOYEE"
  },
  {
    "id": 2,
    "name": "Anna",
    "tags": ["manager"],
    "createdAt": "2022-01-01T00:00:00Z",
    "deletedAt": null,
    "updatedAt": "2022-01-01T00:00:00Z",
    "role": "Manager",
    "type": "EMPLOYEE"
  }
]
```
## Add a new member

Adds a new member.

- Path: `/api/v1/members`
- Method: `POST`
- Consumes: `application/json`
- Produces: `application/json`
- Tags: `members`

### Request Body

- Name: `member`
- Type: `object`
- Required: Yes
- Schema: [models.Member](#definitionsmodels.member)

### Responses

- `200 OK`: Returns the newly added member object.
#### Example:
```json
{
    "name": "Mike",
    "tags": ["developer"],
    "contractDuration": 12,
    "role": "Developer",
    "type": "CONTRACTOR"
}
```

## Get a member by ID

Gets a member by their ID with their associated tags.

- Path: `/api/v1/members/{id}`
- Method: `GET`
- Produces: `application/json`
- Tags: `members`

### Parameters

- Name: `id`
- Type: `string`
- Description: Member ID
- In: `path`
- Required: Yes

### Responses

- `200 OK`: Returns the member object with tags.

#### Example:
```json
{
  "id": 1,
  "name": "John",
  "tags": ["developer"],
  "contractDuration": 12,
  "createdAt": "2022-01-01T00:00:00Z",
  "deletedAt": null,
  "updatedAt": "2022-01-01T00:00:00Z",
  "role": "Developer",
  "type": "EMPLOYEE"
}
```
## Delete a member

Deletes a member by their ID.

- Path: `/api/v1/members/{id}`
- Method: `DELETE`
- Tags: `members`

### Parameters

- Name: `id`
- Type: `string`
- Description: Member ID
- In: `path`
- Required: Yes

### Responses

- `200 OK`: Member deleted successfully.

## Update a member by ID

Updates a member by their ID.

- Path: `/api/v1/members/{id}`
- Method: `PUT`
- Tags: `members`

### Parameters

- Name: `id`
- Type: `string`
- Description: Member ID
- In: `path`
- Required: Yes

### Request Body

- Name: `member`
- Type: `object`
- Required: Yes
- Schema: [models.Member](#definitionsmodels.member)

#### Request Body

- Name: `member`
- Type: `object`
- Required: Yes
- Schema: `models.Member`

#### Example:
```json
{
  "name": "Updated John",
  "tags": ["Java"],
  "contractDuration": 12,
  "role": "Developer",
  "type": "CONTRACTOR"
}
````

### Responses

- `200 OK`: Member updated successfully.


## Get all tags

Gets all tags. Ideal for populating a combo box in the UI for example or a filter.
Does not return associated members for efficiency.

- Path: `/api/v1/tags`
- Method: `GET`
- Produces: `application/json`
- Tags: `tags`

### Responses

- `200 OK`: Returns an array of tag objects.

## Definitions

### models.Member

Represents a member object.

Properties:

- `contractDuration` (integer)
- `createdAt` (string)
- `deletedAt` (string)
- `id` (integer)
- `name` (string)
- `role` (string)
- `tags` (array of [models.Tag](#definitionsmodels.tag))
- `type` (string, enum: EMPLOYEE, CONTRACTOR)
- `updatedAt` (string)

### models.Tag

Represents a tag object.

Properties:

- `createdAt` (string)
- `deletedAt` (string)
- `id` (integer)
- `members` (array of [models.Member](#definitionsmodels.member))
- `name` (string)
- `updatedAt` (string)
