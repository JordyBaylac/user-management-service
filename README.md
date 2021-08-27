# user-management-service


Simple HTTP server that manages user data.

## Functionalities
1. __Create new user__. Given user email and name, it will return a created user with a unique system assigned user ID.
2. __Get user data__. Given a system assigned user ID, it will returns saved user data.
3. __Update user data__. Given a system assigned user ID, it will allow to change user attributes.

## Architecture
<!-- TODO: diagram -->

### Folder structure
- _api/_ package is the interface of this service. It's where the http server is configured (routes, middlewares, etc)
- _user/_ it is a domain specific package that contains business regarding the management of users.

### Dependencies
- [https://gofiber.io/](https://gofiber.io/): web framework used, an express like for Go.

## Run locally
### (Option 1) Using __air__ for hot reloading
```sh
# install air
go get -u github.com/cosmtrek/air

# run
air
```

### (Option 2) Dockerize the service
```sh
```

## Test
### Usage
### cURL commands
