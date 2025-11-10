package config

type Config struct {
	DBURL string
	HTTPPort string
}

func Load() Config {
	return Config{
		DBURL:  "postgres://postgres:gaurav@localhost:5432/User-Management?sslmode=disable",
		HTTPPort: ":8080",
	}
}