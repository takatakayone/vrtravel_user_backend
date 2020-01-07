package app

import "github.com/gin-gonic/gin"

var (
	router = gin.Default()
)

func StartApplication() {
	urlMappings()

	router.Run(":8080")
}
