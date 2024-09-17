# Go Movies App

A simple web application built using the **Gin** framework and **PostgreSQL**.

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