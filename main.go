package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/status/:status", func(c *gin.Context) {
		status, _ := strconv.Atoi(c.Param("status"))
		c.JSON(status, gin.H{
			"time": time.Now().UnixNano(),
		})
	})
	r.GET("/wait/:seconds", func(c *gin.Context) {
		seconds, _ := strconv.Atoi(c.Param("seconds"))
		if seconds > 60 {
			seconds = 60
		}

		time.Sleep(time.Duration(seconds) * time.Second)
		c.JSON(http.StatusOK, gin.H{
			"time": time.Now().UnixNano(),
		})
	})
	r.Run()
}
