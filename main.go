package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"fmt"
	"encoding/json"
	"mms/orm"
	"mms/domain"
	"strconv"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/resources", gin.Dir("resources", false))

	router.GET("/", func(c *gin.Context) {
		data := orm.GetMarquees()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"marquees":data,
		})
	})

	router.POST("/add", func(c *gin.Context){
		var mcGee domain.McGee
		err := c.Bind(&mcGee)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"result": "Failed",
			})
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

		var mcGee domain.McGee
		err := c.Bind(&mcGee)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"result": "Failed",
			})
		}

		orm.UpdMarquee(seq, mcGee)

		c.JSON(http.StatusOK, gin.H{
			"result": "ok",
		})
	})

	router.Run(":8081")
}
