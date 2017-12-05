package readwrite_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testConfig struct {
	TimeoutScale                  float64 `json:"timeout_scale"`
	MongoHost                     string `json:"mongo_host"`
	MongoPort                     string `json:"mongo_port"`
	MongoRoot                     string `json:"mongo_root"`
	MongoRootPassword             string `json:"mongo_root_password"`
}

func loadConfig(path string) (cfg testConfig) {
	configFile, err := os.Open(path)
	if err != nil {
		fatal(err)
	}

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(&cfg); err != nil {
		fatal(err)
	}

	return
}

var (
	config = loadConfig(os.Getenv("CONFIG_PATH"))
)

func fatal(err error) {
	fmt.Printf("ERROR: %s\n", err.Error())
	os.Exit(1)
}

func TestReadwrite(t *testing.T) {

	RegisterFailHandler(Fail)

	RunSpecs(t, "MongoDB Acceptance Tests")
}
