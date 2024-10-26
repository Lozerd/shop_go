package logging

import (
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateLogger(t *testing.T) {
	// Cleanup: remove logs folder after test
	defer os.RemoveAll("logs")

	t.Run("creates log file successfully", func(t *testing.T) {
		logger := CreateLogger("test.log")
		assert.NotNil(t, logger)

		// Check if file is created
		_, err := os.Stat("logs/test.log")
		assert.NoError(t, err, "Log file should be created")
	})

	t.Run("panics on empty filename", func(t *testing.T) {
		assert.Panics(t, func() {
			CreateLogger("")
		}, "Should panic when filename is empty")
	})
}

func TestInitLogging(t *testing.T) {
	// Cleanup: remove logs folder after test
	defer os.RemoveAll("logs")

	t.Run("initializes logging map and creates files", func(t *testing.T) {
		InitLogging()

		assert.NotNil(t, Logger)
		assert.NotNil(t, Logger["main"])
		assert.NotNil(t, Logger["request"])

		// Check if the log files are created
		_, err := os.Stat("logs/main.log")
		assert.NoError(t, err, "Main log file should be created")

		_, err = os.Stat("logs/request.log")
		assert.NoError(t, err, "Request log file should be created")
	})
}

