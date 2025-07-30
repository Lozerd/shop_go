package postgres

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnection(t *testing.T) {
	os.Setenv("DB_NAME", "postgres")
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	dsn := GetDSN()
	assert.Equal(t, "postgres://postgres:postgres@localhost:5432/postgres?charset=utf8&parseTime=True&loc=Local", dsn)
}
