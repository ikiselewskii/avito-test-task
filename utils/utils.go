package utils

import (
	"log"
	"os"
)

func ParseEnvToDSN() (dsn string) {
	dsn = os.Getenv("DSN")
	log.Println(dsn)
	return
}
