package config_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/andreyvit/diff"
)

const testEnvContent string = `
POSTGRES_NAME=dev_shop
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=dev_shop
POSTGRES_PASSWORD=password
`

func TestLoadEnv_ShouldLoadOn_APP_ENV(t *testing.T) {
	t.Parallel()

	if err := config.LoadEnv(); err == nil {
		t.Fatal("Should panic on non existent file!")
	}

	f, err := os.Create(".env")
	defer f.Close()
	defer os.Remove(f.Name())

	if err != nil {
		t.Fatal(err)
	}
	
    expected := "test"
    content := fmt.Sprintf("TEST_ENV=\"%s\"", expected)
	if _, err = f.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
    
	if err = config.LoadEnv(); err != nil {
		t.Fatal(err)
	}

    if envVar := os.Getenv("TEST_ENV"); envVar != expected {
		t.Fatal("Couldn't load TEST_ENV variable!")
    }
}

func TestLoadEnv_ShouldFailAtWrongFilePath(t *testing.T) {
	t.Parallel()

	os.Setenv("APP_ENV", ".non-existent-env")

	err := config.LoadEnv()
	if err == nil {
		t.Error("Should return error on non existent file!")
	}

	expected := fmt.Sprintf("Filepath do not exist\nOriginal: ")
	if !strings.Contains(err.Error(), expected) {
		t.Errorf("Should return proper error message!\n%v", diff.CharacterDiff(expected, err.Error()))
	}
}

func TestLoadEnv_ShouldLoadVariables(t *testing.T) {
	t.Parallel()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	dir, err := os.MkdirTemp(cwd, "TestEnvTempDir")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	testEnvFilename := ".test-env-file"
	f, err := os.CreateTemp(dir, testEnvFilename)
	defer f.Close()

	if err != nil {
		t.Fatal(err)
	}

    _, err = f.Write([]byte(testEnvContent))

	if err != nil {
		t.Fatal(err)
	}

	os.Setenv("APP_ENV", f.Name())
	config.LoadEnv()

	for _, line := range strings.Fields(testEnvContent) {
		splitRes := strings.Split(line, "=")

		if len(splitRes) < 2 {
			t.Fatal("Couldn't parse env variable")
		}

		if envVar := os.Getenv(splitRes[0]); envVar == "" {
			t.Fatalf("Couldn't load \"%s\" env variable", splitRes[0])
		}
	}
	pgName := os.Getenv("POSTGRES_NAME")
	if pgName == "" {
		t.Fatal("Env variable wasn't loaded")
	}
}

//
// 	// POSTGRES_NAME=dev_shop
// 	// POSTGRES_HOST=localhost
// 	// POSTGRES_PORT=5432
// 	// POSTGRES_USER=dev_shop
// 	// POSTGRES_PASSWORD=password
//
//     if cfg == nil {
//         t.Fatal("Config shouldn't be nil!")
//     }
//
//
// 	// if cfg.Database.Name != "dev_shop" ||
// 	// 	cfg.Database.Host != "localhost" ||
// 	// 	cfg.Database.Port != "5432" ||
// 	// 	cfg.Database.User != "dev_shop" ||
// 	// 	cfg.Database.Pass != "password" {
// 	// 	t.Fatalf("Env variables wasn't loaded correctly, should be: %s, got: %v", testEnvContent, cfg)
// 	// }
//
// }
