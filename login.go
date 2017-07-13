package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	//"fmt"
	"crypto/sha256"
	"encoding/base64"
	"mms/orm"
	"strings"
)

type UserLogin struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	session := sessions.Default(c)

	var form UserLogin
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusOK, gin.H{"result":"Form binding error!"})
	} else {
		user := orm.GetSingleUser(form.Username)

		passwordSha, err := sha([]byte(form.Password))
		if err != nil {
			panic("SHA hash failed.")
		}

		if result := strings.Compare(user.Password, passwordSha); result != 0 {
			c.JSON(http.StatusOK, gin.H{"result":"帳號或密碼輸入錯誤!"})
		} else {
			session.Set("username", form.Username)
			session.Save()
			c.Redirect(http.StatusSeeOther, "/")
		}
	}
}

func AddUser(c *gin.Context) {
	var userAdd UserLogin
	if err := c.Bind(&userAdd); err != nil {
		c.JSON(http.StatusOK, gin.H{"result":"Form binding error!"})
	} else {
		passwordSha, _ := sha([]byte(userAdd.Password))
		orm.AddUser(userAdd.Username, passwordSha)
		c.JSON(http.StatusOK, gin.H{"result":"OK"})
	}
}

func sha(secret []byte) (string, error) {
	h := sha256.New()

	_, err := h.Write(secret)
	if err != nil {
		return "", err
	}

	sha := base64.URLEncoding.EncodeToString(h.Sum(nil))
	return sha, nil
}
