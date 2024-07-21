package server

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"barosa.fun/azure-ai-stream-backend/auth"
	"github.com/gin-gonic/gin"
)

var validFeatures = map[string]bool{
    "read":		true,
    "caption":		true,
    "denseCaptions":	true,
    "smartCrops":	true,
    "objects":		true,
    "tags":		true,
    "people":		true,
}

func RequestAuthorizeMiddleware() gin.HandlerFunc {
	return auth.AuthCheckBearerTokenServerMiddleware 
}

func RequestPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func RequestImageFeatures(c *gin.Context) {
	features := strings.TrimSpace(c.Query("features"))
	if len(features) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "features query parameter is required",
		})
		return
	}

	featuresSlice := strings.Split(features, ",") 
	for i := range featuresSlice {
		if !validFeatures[featuresSlice[i]] {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("provided feature '%s' is not valid", featuresSlice[i]),
			})
			return
		}
	}
	fmt.Printf("%v", featuresSlice)
}

func Init() {
	log.Println("Starting barosa backend")
	r := gin.Default()
	r.Use(RequestAuthorizeMiddleware())
	r.GET("/ping", RequestPing)
	r.GET("/image-features", RequestImageFeatures)
	r.Run()
}
