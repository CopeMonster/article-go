# Article-Go

## Description
*Article-Go* is a backend service built with Go for publishing and reading articles. This project aims to provide API for creating/reading/updating/deleting articles.

## Installation
To install *Article-Go*, you need to have Go installed on your machine. Once you have Go installed, you can clone this repository and build the project:

```bash
git clone https://github.com/CopeMonster/article-go.git
cd article-go
go build ./cmd/article-go/main.go
```

## Usage
After building the project, you can run the API server with:

```bash
go run ./cmd/article-go/main.go
```
The server will start on the default port 3000. You can then send HTTP requests to the server to interact with the API.

## API Endpoints

### Auth
* `POST /auth/sign-up` Use this endpoint to register a new user to the application.
* `POST /auth/sign-in` Use this endpoint to authenticate a user and log them into the application

### Todos
* `GET /article/` Use this endpoint to retrieve all articles. (Works without log in into system)
* `GET /article/:{id}` Use this endpoint to retrieve article by its id. Replace `{id}` with the actual ID of the article. (Works without log in into system)
* `POST /article/` Use this endpoint to add a new article. (Works only with log in into system)
* `PUT /article/:{id}` Use this endpoint to update the details of a specific article. Replace `{id}` with the actual ID of the article. (Works only with log in into system)
* `DELETE /article/:{id}` Use this endpoint to remove a specific article. Replace `{id}` with the actual ID of the article. (Works only with log in into system)

## Database
This project uses Postgresql as its primary database.

## Contributing
Contributions are welcome! Please feel free to submit a Pull Request.

## License
This project is licensed under the [MIT License](https://github.com/CopeMonster/article-go/blob/master/LICENSE).
