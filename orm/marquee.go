package orm

import (
	"fmt"
)

type Marquee struct {
	Seq int `gorm:"primary_key" json:"seq"`
	Title string `json:"title"`
	StartTime string `json:"start_time"`
	EndTime string `json:"end_time"`
}

// set User's table name to be `marquee`
func (Marquee) TableName() string {
	return "marquee"
}

func AddMarquee(m Marquee) {
	err := gdb.Create(&m).Error
	if err != nil {
		fmt.Println(err)
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