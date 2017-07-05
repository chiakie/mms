package main

import (
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"fmt"
	"encoding/json"
	"marquee/orm"
)

type Marquee struct {
	Title     string `json:"title" binding:"required"`
	StartTime string `json:"start_time" binding:"required"`
	EndTime   string `json:"end_time"  binding:"required"`
}

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
		var mcGee Marquee
		err := c.Bind(&mcGee)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{
				"result": "Failed",
			})
		}

		blob, _ := json.Marshal(mcGee)

		c.JSON(http.StatusOK, gin.H{
			"result": "OK",
			"data": string(blob),
		})
	})

	router.Run(":8081")
}
