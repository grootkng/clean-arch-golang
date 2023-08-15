package config

type Env struct {
	DB       string
	API_PORT string
}

func GetEnv() *Env {
	return &Env{
		DB:       "host=localhost user=postgres password=postgres dbname=db port=5432",
		API_PORT: ":8080",
	}
}
