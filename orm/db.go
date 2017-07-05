package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Marquee struct {
	Seq uint `gorm:"primary_key;auto_increment"`
	Title string `gorm:"type:varchar(255);not null"`
	StartTime string `gorm:"not null"`
	EndTime string `gorm:"not null"`
}

var gdb *gorm.DB

func init() {
	var err error
	gdb, err = gorm.Open("sqlite3", "marquee.db");
	if err != nil {
		panic("failed to connect database")
	}
}

// set User's table name to be `profiles`
func (Marquee) TableName() string {
	return "marquee"
}

func GetMarquees() []Marquee {
	var marquees []Marquee
	gdb.Find(&marquees)
	return marquees
}