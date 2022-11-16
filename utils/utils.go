package utils

import (
	"fmt"
	"log"
	"os"
)


//DSN = "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$POSTGRES_HOTNAME:5432/db?sslmode=disable"
func SerializeDSN() (dsn string) {
	
	dsn = fmt.Sprintf("postgres://%s:%s@%s:5432/db&sslmode=disable",
					os.Getenv("POSTGRES_USER"),
	 				os.Getenv("POSTGRES_PASSWORD"),
					os.Getenv("POSTGRES_HOSTNAME"))
	log.Println(dsn)
	return
}
