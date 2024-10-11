package config

import (
	"fmt"
	"os"
)

func GetPostgresURI() string {
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	postgresURI := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname)

	return postgresURI
}

// GetGmailCredentialsPath returns the path to the Gmail credentials file (if needed for your project)
func GetGmailCredentialsPath() string {
	return "credentials.json"
}
