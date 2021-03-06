package orm

import (
	"fmt"
)

type User struct {
	Username string `gorm:"not null;unique" json:"username"`
	Password string `gorm:"not null;" json:"password"`
}

// set User's table name to be `marquee`
func (User) TableName() string {
	return "user"
}

func GetSingleUser(name string) (User, error) {
	var user User
	err := gdb.Where("username = ?", name).First(&user).Error
	if err != nil {
		fmt.Println(fmt.Sprintf("Username \"%s\" is not valid.", name))
	}

	return user, err
}

func AddUser(name string, password string) error {
	var user User
	user.Username = name
	user.Password = password

	err := gdb.Create(&user).Error
	if err != nil {
		fmt.Println("Error:", err)
	}

	return err
}

