package main

import "github.com/grootkng/clean-arch-golang/internal/pkg/repository/database"

func main() {
	db, err := database.Db()

	if err != nil {
		panic(err.Error())
	}

	db.Exec("CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(100), gender VARCHAR(100), age INTEGER);")
}
