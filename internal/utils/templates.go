package utils

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/lozerd/shop_go/infrastructure/logging"
)

var logger = logging.NewLogger("utils", true)

func LoadTemplates(r *gin.Engine, relpath string) ([]string, error) {
	var templateFiles []string

	err := filepath.WalkDir(relpath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".html" {
			templateFiles = append(templateFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return templateFiles, nil
}
