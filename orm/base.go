package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

var gdb *gorm.DB

func init() {
	var err error
	gdb, err = gorm.Open("sqlite3", "marquee.db");
	if err != nil {
		panic("failed to connect database")
	}

	err = gdb.AutoMigrate(Marquee{}, User{}).Error
	if err != nil {
		panic(fmt.Sprintf("AutoMigration Failed: %s", err))
	}
}