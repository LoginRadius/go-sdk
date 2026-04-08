package lrintegrationtest

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func SetTestEnv() {
	if err := loadTestEnv(); err != nil {
		log.Fatal("Error loading env files, please configure your secret.env and public.env.")
	}
}

func loadTestEnv() error {
	cwd, _ := os.Getwd()

	return godotenv.Load(
		filepath.Join(cwd, "./config/secret.env"),
		filepath.Join(cwd, "./config/public.env"),
	)
}

func testStubEnvSetup() {

}
