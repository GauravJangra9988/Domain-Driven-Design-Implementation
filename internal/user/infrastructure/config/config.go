package config

type Config struct {
	DBURL string
	HTTPPort string
}

func Load() Config {
	return Config{
		DBURL:  "postgres://user:dbpassword@db:5432/appdb?sslmode=disable",
		HTTPPort: ":8080",
	}
}