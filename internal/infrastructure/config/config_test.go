package config_test

import (
	"fmt"
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

func TestLoadEnv_WillPanicOnEmptyString(t* testing.T) {
    t.Parallel()

    err := config.LoadEnv("")
    if err == nil {
        t.Fatal("Should panic on empty filepath!")
    }

    expected := "filepath argument must not be empty"
    if err.Error() != expected {
        t.Fatalf("Should return proper error message!\n%v", diff.CharacterDiff(expected, err.Error()))
    }
}

func TestLoadEnv_ShouldFailAtWrongFilePath(t* testing.T) {
    t.Parallel()

    filepath := ".non-existent-env"
    err := config.LoadEnv(filepath)
    if err == nil {
        t.Error("Should return error on non existent file!")
    }

    expected := fmt.Sprintf("Couldn't load \"%s\" environment file!", filepath)
    if err.Error() != expected {
        t.Errorf("Should return proper error message!\n%v", diff.CharacterDiff(expected, err.Error()))
    }
}

func TestGetConfig_IsOfTypeConfiguration(t* testing.T) {
    t.Parallel()

    _, ok := config.GetConfig().(*config.Configuration)

    if ok == false {
        t.Fatal("GetConfig() should return type *Configuration!")
    }
}

func TestLoadEnv_ShouldLoadVariables(t* testing.T) {
    t.Parallel()

}
