package db_test

import (
	"fmt"
	// "testing"

	// . "github.com/Lozerd/shop_go/internal/infrastructure/db"
	// "gorm.io/driver/postgres"
	// "gorm.io/gorm"
)

var testDBName string = "test_dev_shop"

func getTestDbUrl() string {
    return fmt.Sprint("")
}

// func Test_SetupTestDB_InterfaceIsRight(t *testing.T) {
//     teardownSuite := SetupTestDB(t)
// 
//     db, err := gorm.Open(postgres.Open(getTestDbUrl()))
//     if err != nil {
//         t.Error(err)
//     }
// 
//     err = db.Exec("select 1 from pg_database where").Error
//     if err != nil {
//         t.Error(err)
//     }
// 
//     defer teardownSuite(t)
// }
