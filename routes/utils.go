package routes

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DEBUG_MODE = os.Getenv("DEBUG_MODE")

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	os.Getenv("POSTGRES_HOST"),
	os.Getenv("POSTGRES_PORT"),
	os.Getenv("POSTGRES_USER"),
	os.Getenv("POSTGRES_PASSWORD"),
	os.Getenv("POSTGRES_NAME"),
)
