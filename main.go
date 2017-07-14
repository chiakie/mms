package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"net/http"
	"fmt"
	"encoding/json"
	"mms/orm"
	"strconv"
	"strings"
	"mms/domain"
)

func CheckUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)

		if !skip(c.Request.URL.Path) {
			username := session.Get("username")
			if username == nil {
				c.Redirect(http.StatusSeeOther, "/login")
			}
		}
		/* ↑↑↑↑↑ Before request */
		c.Next()
		/* ↓↓↓↓↓ After request  */
	}
}

func skip(path string) bool {

	if strings.Contains(path, "/resources") {
		return true
	}

	if strings.Contains(path, "/login") {
		return true
	}

	return false;
}

func main() {
	router := gin.Default()
	// MW: enable cookie-based session
	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{MaxAge:300}) // 10 minutes to expire
	router.Use(sessions.Sessions("MarqueeSession", store))

	// MW: check user
	router.Use(CheckUser())

	// views and resources
	router.LoadHTMLGlob("templates/*")
	router.Static("/resources", "./resources")



	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	router.POST("/login", Login)
	router.POST("/user/add", AddUser)

	router.GET("/data", func(c *gin.Context) {
		data := orm.GetMarquees()
		c.JSON(http.StatusOK, data)
	})

	router.POST("/add", func(c *gin.Context){
		var mcGee domain.McGee
		err := c.Bind(&mcGee)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"result": "failed",
			})
			return
		}

		var marquee orm.Marquee
		marquee.Title = mcGee.Title
		marquee.StartTime = mcGee.StartTime
		marquee.EndTime = mcGee.EndTime
		orm.AddMarquee(marquee)

		blob, _ := json.Marshal(mcGee)
		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
			"data": string(blob),
		})
	})

	router.GET("/del/:seq", func(c *gin.Context) {
		seq, _ := strconv.Atoi(c.Param("seq"))
		orm.DelMarquee(seq)

		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})

	router.GET("/edit/:seq", func(c *gin.Context) {
		seq, _ := strconv.Atoi(c.Param("seq"))

		marquee, err := orm.GetSingleMarquee(seq)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "failed",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
			"data": marquee,
		})
	})

	router.POST("/edit/:seq", func(c *gin.Context) {
		seq, _ := strconv.Atoi(c.Param("seq"))

		var mcgee domain.McGee
		err := c.Bind(&mcgee)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "failed",
			})
			return
		}

		marqueeData := orm.Marquee{Seq:seq, Title:mcgee.Title, StartTime:mcgee.StartTime, EndTime:mcgee.EndTime}
		orm.UpdMarquee(marqueeData)

		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})

	router.Run(":8081")
}
