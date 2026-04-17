package main

import (
	"go-grpc-elk-postgres-microservice/account"
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/tinrab/retry"
)

type Config struct {
	DatabaseURI string `envconfig:"DATABASE_URL"`
}

func main() {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	var r account.Repository
	retry.ForeverSleep(2*time.Second, func(_ int) (err error) {
		r, err = account.NewPostgresRepository(cfg.DatabaseURI)
		if err != nil {
			log.Println(err)
		}
		return
	})

	defer r.Close()

	log.Println("Listening on port 8080.")
	s := account.NewService(r)
	log.Fatal(account.ListenGRPC(s, 8080))
}
