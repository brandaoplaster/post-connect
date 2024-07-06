package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL  = ""
	PORT    = 0
	HashKey []byte
	BlokKey []byte
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	PORT, erro = strconv.Atoi("APP_PORT")
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlokKey = []byte(os.Getenv("BLOCK_KEY"))
}
