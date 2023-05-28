package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("api/rate", getRate)
	router.GET("api/subscribe", getEmails)
	router.POST("api/subscribe", postEmail)
	router.POST("api/sendEmails", sendEmails)
	router.Run("localhost:8080")
}
