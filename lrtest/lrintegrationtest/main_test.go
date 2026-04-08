package lrintegrationtest

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := loadTestEnv(); err != nil {
		log.Printf("Skipping integration tests: configure lrtest/lrintegrationtest/config/secret.env and public.env to run them.")
		os.Exit(0)
	}

	os.Exit(m.Run())
}
