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
		user, _ := orm.GetSingleUser(form.Username)

		passwordSha, err := sha([]byte(form.Password))
		if err != nil {
			panic("SHA hash failed.")
		}

		if result := strings.Compare(user.Password, passwordSha); result != 0 {
			c.JSON(http.StatusOK, gin.H{"result":"failed", "message":"帳號或密碼輸入錯誤!"})
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
		c.JSON(http.StatusOK, gin.H{"result":"failed", "message":"Form binding error!"})
	} else {
		passwordSha, _ := sha([]byte(userAdd.Password))
		err = orm.AddUser(userAdd.Username, passwordSha)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE") {
				c.JSON(http.StatusOK, gin.H{"result":"failed", "message":userAdd.Username + "此帳號已存在!"})
			} else {
				c.JSON(http.StatusOK, gin.H{"result":"failed", "message":err.Error()})
			}
		} else {
			c.JSON(http.StatusOK, gin.H{"result":"ok"})
		}
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

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("username")
	session.Save()

	c.Redirect(http.StatusSeeOther, "/login")
}
