package routes

import (
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DEBUG_MODE = os.Getenv("DEBUG_MODE")

var psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=disable",
	os.Getenv("DATABASE_HOST"),
	os.Getenv("DATABASE_PORT"),
	os.Getenv("DATABASE_USER"),
	os.Getenv("DATABASE_PASSWORD"),
	os.Getenv("DATABASE_NAME"),
)
