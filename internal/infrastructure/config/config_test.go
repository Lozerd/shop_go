package config_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Lozerd/shop_go/internal/infrastructure/config"
	"github.com/andreyvit/diff"
	"github.com/gin-contrib/cors"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
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
		t.Error("Should panic on non existent file!")
	}

	filename := ".test-env"
	f, err := os.Create(filename)
	defer f.Close()
	defer os.Remove(f.Name())

	if err != nil {
		t.Error(err)
	}

	expected := "test"
	content := fmt.Sprintf("TEST_ENV=\"%s\"", expected)
	if _, err = f.Write([]byte(content)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", filename)
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	if err = config.LoadEnv(); err != nil {
		t.Error(err)
	}

	if envVar := os.Getenv("TEST_ENV"); envVar != expected {
		t.Errorf("Couldn't load TEST_ENV variable! diff: %s", diff.CharacterDiff(envVar, expected))
	}
}

func TestLoadEnv_ShouldFailAtWrongFilePath(t *testing.T) {
	t.Parallel()

	os.Setenv("APP_ENV", ".non-existent-env")
	defer os.Clearenv()

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
		t.Error(err)
	}

	dir, err := os.MkdirTemp(cwd, "TestEnvTempDir")
	if err != nil {
		t.Error(err)
	}
	defer os.RemoveAll(dir)

	testEnvFilename := ".test-env-file"
	f, err := os.CreateTemp(dir, testEnvFilename)
	defer f.Close()

	if err != nil {
		t.Error(err)
	}

	if _, err = f.Write([]byte(testEnvContent)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", f.Name())
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	config.LoadEnv()

	for _, line := range strings.Fields(testEnvContent) {
		splitRes := strings.Split(line, "=")

		if len(splitRes) < 2 {
			t.Error("Couldn't parse env variable")
		}

		if envVar := os.Getenv(splitRes[0]); envVar == "" {
			t.Errorf("Couldn't load \"%s\" env variable", splitRes[0])
		}
	}
	pgName := os.Getenv("POSTGRES_NAME")
	if pgName == "" {
		t.Error("Env variable wasn't loaded")
	}
}

func TestLoadConfig_ShouldLoadVariables(t *testing.T) {
	t.Parallel()

	expected := &config.Configuration{
		Database: config.Database{
			Name: "dev_shop",
			User: "dev_shop",
			Host: "localhost",
			Pass: "password",
			Port: "5432",
		},
		CorsConfig: cors.Config{},
		Server: config.Server{
			Port: "8080",
			Host: "0.0.0.0",
		},
		Client: config.Client{
			Port: "8000",
			Host: "0.0.0.0",
		},
		ApiVersion: "v1",
		ApiPrefix:  "api",
	}

	filename := ".test-env"
	f, err := os.Create(filename)
	defer f.Close()
	defer os.Remove(f.Name())

	if err != nil {
		t.Error(err)
	}

	if _, err = f.Write([]byte(testEnvContent)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", f.Name())
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	config.LoadConfig()
	cfg := config.GetConfig()

	if !cmp.Equal(cfg, expected, cmpopts.IgnoreUnexported(config.Configuration{})) {
		t.Errorf(
            "Configuration is not equal to expected, diff: %v", 
            cmp.Diff(cfg, expected, cmpopts.IgnoreUnexported(config.Configuration{})),
        )
	}
}

func Test_GetApiBasePath(t *testing.T) {
	filename := ".test-env"
	f, err := os.Create(filename)
	defer f.Close()
	defer os.Remove(f.Name())
	if err != nil {
		t.Error(f)
	}

	if _, err = f.Write([]byte(testEnvContent)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", f.Name())
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	config.LoadConfig()
	cfg := config.GetConfig()

	testCases := []struct {
		name     string
		actual   string
		expected string
	}{
		{name: "GetApiPrefix", actual: cfg.GetApiPrefix(), expected: fmt.Sprintf("/%s", cfg.ApiPrefix)},
		{name: "GetApiVersion", actual: cfg.GetApiVersion(), expected: fmt.Sprintf("/%s", cfg.ApiVersion)},
		{
			name:     "GetApiBasePath",
			actual:   cfg.GetApiBasePath(),
			expected: fmt.Sprintf("/%s/%s", cfg.ApiPrefix, cfg.ApiVersion),
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("test %s", tc.name), func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("%s = %q;\ndiff %q", tc.name, tc.actual, diff.CharacterDiff(tc.actual, tc.expected))
			}
		})
	}
}

func Test_GetAddr(t *testing.T) {
	filename := ".test-env"
	f, err := os.Create(filename)
	defer f.Close()
	defer os.Remove(f.Name())
	if err != nil {
		t.Error(f)
	}

	if _, err = f.Write([]byte(testEnvContent)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", f.Name())
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	config.LoadConfig()
	cfg := config.GetConfig()

	actual := cfg.GetServerAddr()
	expected := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)

	if actual != expected {
		t.Errorf("GetAddr = %q;\ndiff %q", actual, diff.CharacterDiff(actual, expected))

	}
}

func Test_GetDBUrl(t *testing.T) {
	filename := ".test-env"
	f, err := os.Create(filename)
	defer f.Close()
	defer os.Remove(f.Name())
	if err != nil {
		t.Error(f)
	}

	if _, err = f.Write([]byte(testEnvContent)); err != nil {
		t.Error(err)
	}

	err = os.Setenv("APP_ENV", f.Name())
	defer os.Clearenv()
	if err != nil {
		t.Error(err)
	}

	config.LoadConfig()
	cfg := config.GetConfig()

	actual := cfg.GetDBUrl()
	expected := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		cfg.Database.Host,
		cfg.Database.User,
		cfg.Database.Pass,
		cfg.Database.Name,
		cfg.Database.Port,
	)

	if actual != expected {
		t.Errorf("GetAddr = %q;\ndiff %q", actual, diff.CharacterDiff(actual, expected))
	}
}
