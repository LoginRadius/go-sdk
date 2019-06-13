package lrintegrationtest

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SetTestEnv() {
	cwd, _ := os.Getwd()

	err := godotenv.Load(
		filepath.Join(cwd, "./config/secret.env"),
		filepath.Join(cwd, "./config/public.env"),
	)

	if err != nil {
		log.Fatal("Error loading env files, please configure your secret.env and public.env.")
	}
}

func testStubEnvSetup() {

}
