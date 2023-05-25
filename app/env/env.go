package env

import (
	"github.com/joho/godotenv"
	"github.com/jxmexdev/go-todo-app/app/commons"
	"github.com/kelseyhightower/envconfig"
	"log"
)

type Environment struct {
	AppHost    string `envconfig:"APP_HOST" default:"localhost"`
	AppPort    string `envconfig:"APP_PORT" default:"5000"`
	AppEnv     string `envconfig:"APP_ENV" default:"development"`
	JwtSecret  string `envconfig:"JWT_SECRET" required:"true"`
	DbDriver   string `envconfig:"DB_DRIVER" required:"true"`
	DbUser     string `envconfig:"DB_USER" required:"true"`
	DbPassword string `envconfig:"DB_PASSWORD" required:"true"`
	DbHost     string `envconfig:"DB_HOST" required:"true"`
	DbPort     string `envconfig:"DB_PORT" required:"true"`
	DbName     string `envconfig:"DB_NAME" required:"true"`
}

var Conf Environment

// LoadConfiguration loads env variables into the Config variable
func LoadConfiguration() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = envconfig.Process("", &Conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	hasAllowedDriver := false
	for _, driver := range commons.GetAllowedDbDrivers() {
		if driver == Conf.DbDriver {
			hasAllowedDriver = true
		}
	}
	if !hasAllowedDriver {
		log.Fatalf("DB_DRIVER must be one of %v , got %s", commons.GetAllowedDbDrivers(), Conf.DbDriver)
	}
	log.Println("EnvConfig variables loaded")
}
