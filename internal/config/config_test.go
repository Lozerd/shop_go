package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoad(t *testing.T) {
	os.Clearenv()
	os.Setenv("GIN_MODE", "debug")
	os.Setenv("PORT", "8000")
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "debug", cfg.GinMode)
	assert.Equal(t, 8000, cfg.Port)
	assert.Equal(t, "postgres", cfg.Name)
	assert.Equal(t, "localhost", cfg.Host)
	assert.Equal(t, 5432, cfg.DB.Port)
	assert.Equal(t, "postgres", cfg.User)
	assert.Equal(t, "postgres", cfg.Password)
}

func TestRequiredEnv(t *testing.T) {
	os.Clearenv()

	if _, err := Load(); assert.Error(t, err) {
		assert.Equal(t, "value for this field is required [DB_NAME]", err.Error())
	}

	os.Setenv("DB_NAME", "postgres")

	if _, err := Load(); assert.Error(t, err) {
		assert.Equal(t, "value for this field is required [DB_USER]", err.Error())
	}

	os.Setenv("DB_USER", "postgres")

	if _, err := Load(); assert.Error(t, err) {
		assert.Equal(t, "value for this field is required [DB_PASSWORD]", err.Error())
	}

	os.Setenv("DB_PASSWORD", "postgres")
	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "postgres", cfg.DB.Name)
	assert.Equal(t, "postgres", cfg.DB.User)
	assert.Equal(t, "postgres", cfg.DB.Password)
}

func TestDefaultValues(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_USER", "postgres")
	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "release", cfg.GinMode)
	assert.Equal(t, 8000, cfg.Port)
	assert.Equal(t, "localhost", cfg.DB.Host)
	assert.Equal(t, 5432, cfg.DB.Port)
	assert.Equal(t, "api", cfg.ApiPrefix)
}

func TestInvalidGinMode(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("GIN_MODE", "invalid")

	if _, err := Load(); assert.Error(t, err) {
		if _, ok := err.(ErrInvalidGinMode); !ok {
			t.Errorf("Expected ErrInvalidGinMode, got %T", err)
		}
		assert.Equal(t, "invalid gin mode [invalid], should be one of [debug, release].", err.Error())
	}
}

func TestTrustedProxies(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("TRUSTED_PROXIES", "127.0.0.1")

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, []string{"127.0.0.1"}, cfg.TrustedProxies)
}

func TestGetAddr(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_USER", "postgres")

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, ":8000", cfg.GetAddr())
}

func TestGetApiPrefix(t *testing.T) {
	os.Clearenv()
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("API_PREFIX", "different")

	cfg, err := Load()
	assert.NoError(t, err)
	assert.Equal(t, "/different", cfg.GetApiPrefix())
}
