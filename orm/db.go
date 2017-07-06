package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"mms/domain"
)

type Marquee struct {
	Seq uint
	Title string
	StartTime string
	EndTime string
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

func AddMarquee(m Marquee) {
	err := gdb.Create(&m).Error
	if err != nil {

	}
}

func GetSingleMarquee(seq int) (Marquee, error) {
	var m Marquee
	err := gdb.Where("seq = ?", seq).First(&m).Error
	if err != nil {

	}
	return m, err;
}

func GetMarquees() []Marquee {
	var marquees []Marquee
	err := gdb.Order("start_time desc").Find(&marquees).Error
	if err != nil {

	}
	return marquees
}

func UpdMarquee(seq int, m interface{}) {
	mcgee := m.(domain.McGee)
	err := gdb.Model(Marquee{}).Where("seq = ?", seq).Updates(map[string]interface{}{"title": mcgee.Title, "start_time":mcgee.StartTime, "end_time":mcgee.EndTime}).Error
	if err != nil {

	}
}

func DelMarquee(seq int) {
	err := gdb.Where("seq = ?", seq).Delete(Marquee{}).Error
	if err != nil {

	}
}