# Go Movies App

A simple web application built using the **Gin** framework and **PostgreSQL**.

## Used tools
* Gin web framework
* Swagger for API docs
* ORM -> Gorm
* Migrations -> Goose
* Swagger
* PostgreSQL

## How to Run

1. **Clone the repository:**

   ```bash
   git clone https://github.com/Olim2508/go-movies-app.git
   cd go-movies-app

2. **Configure PostgreSQL creds in .env file:**
```
DB_HOST=localhost
DB_PORT=5432
DB_USER=your_username
DB_PASSWORD=your_password
DB_NAME=movie_db
```
3. **Install deps**
```
go mod tidy
```

3. **Run the app**
```
go run main.go
```
#### It will listen to address http://localhost:8080/ or you can open Swagger docs page on http://localhost:8080/swagger/index.html


**Migrations**

For migrations goose library is being used. After adding new models in go packages, you should run the command:
```
goose -dir ./db/migrations create migration_file_name sql
```

It will create empty sql migration file in `/db/migrations` folder. You should add table creating/updating changes in this sql file like this:
```
-- +goose Up
-- +goose StatementBegin
CREATE TABLE directors (
                           id SERIAL PRIMARY KEY,
                           name VARCHAR(255) NOT NULL,
                           age INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS directors;
-- +goose StatementEnd
```


Then you can apply migrations by the following command:
```
goose -dir ./db/migrations postgres "user=postgres password=postgres dbname=movie_db sslmode=disable" up
```

To learn more about goose, you can visit its [docs](https://github.com/pressly/goose?tab=readme-ov-file)
