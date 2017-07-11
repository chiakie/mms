package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type Marquee struct {
	Seq int `gorm:"primary_key" json:"seq"`
	Title string `json:"title"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

var gdb *gorm.DB

func init() {
	var err error
	gdb, err = gorm.Open("sqlite3", "marquee.db");
	if err != nil {
		panic("failed to connect database")
	}

	err = gdb.AutoMigrate(Marquee{}).Error
	if err != nil {
		panic(fmt.Sprintf("AutoMigration Failed: %s", err))
	}
}

// set User's table name to be `profiles`
func (Marquee) TableName() string {
	return "marquee"
}

func AddMarquee(m Marquee) {
	err := gdb.Create(&m).Error
	if err != nil {

	}
}

func GetSingleMarquee(seq int) (Marquee, error) {
	var m Marquee
	err := gdb.Where("seq = ?", seq).First(&m).Error
	if err != nil {
		fmt.Println(err)
	}
	return m, err
}

func GetMarquees() []Marquee {
	var marquees []Marquee
	err := gdb.Order("start_time desc").Find(&marquees).Error
	if err != nil {
		fmt.Println(err)
	}
	return marquees
}

func UpdMarquee(m Marquee) {
	err := gdb.Save(&m).Error
	if err != nil {
		fmt.Println(err)
	}
}

func DelMarquee(seq int) {
	err := gdb.Where("seq = ?", seq).Delete(Marquee{}).Error
	if err != nil {
		fmt.Println(err)
	}
}