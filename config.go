package bot

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

const (
	Development = "dev"
	Production  = "prod"
)

type Config struct {
	Env    string `env:"ENV,default=dev"`
	Sock   string `env:"SOCK"`
	Port   string `env:"PORT"`
	DbFile string `env:"DB_FILE,required"`
	// Host     string
	// Port     int    `default:"8080"`

	MediatorWorkers   int `ENV:"MEDIATOR_WORKERS,default=0"`
	MediatorQueueSize int `ENV:"MEDIATOR_QUEUESIZE,default=0"`
}

func NewConfig() *Config {
	// load .env file
	env, exists := os.LookupEnv("ENV")
	if !exists {
		env = "dev"
	}
	file := ".env." + env
	err := godotenv.Load(file)

	if err != nil {
		panic(err)
	}
	// if env == "prod" {
	// 	Param.Env = Production
	// } else {
	// 	Param.Env = Development
	// }
	fmt.Println("Successfully loaded env file: " + file)
	ctx := context.Background()

	c := &Config{}
	if err := envconfig.Process(ctx, c); err != nil {
		panic(err)
	}
	return c
}
