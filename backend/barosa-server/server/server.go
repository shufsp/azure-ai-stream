package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"image"
	"strconv"
	"barosa.fun/azure-ai-stream-backend/auth"
	"barosa.fun/azure-ai-stream-backend/command"
	"barosa.fun/azure-ai-stream-backend/compression"
	"github.com/gen2brain/avif"
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

func RequestCorsMiddleware() gin.HandlerFunc {
	return auth.CorsServerMiddleware
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
	window := strings.TrimSpace(c.Query("window"))
	if len(window) == 0 {	
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "window query parameter is required",
		})
		return
	}

	windowSearchMethod := strings.TrimSpace(c.Query("method"))
	if len(windowSearchMethod) == 0 {
		windowSearchMethod = "class"
	}

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

	// get screenshot of window
	screenshotFile, err := command.CommandRunBarosaScreenshot(window, windowSearchMethod)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("window screenshot failed: %v", err),
		})
		return
	}

	// avif/lanzcos
	avifQuality, err := strconv.Atoi(c.Query("avifQuality"))
	if err != nil {
		avifQuality = 30
	}
	avifAlphaQuality, err := strconv.Atoi(c.Query("avifAlphaQuality"))
	if err != nil {
		avifAlphaQuality = 10
	}
	avifSpeed, err := strconv.Atoi(c.Query("avifSpeed"))
	if err != nil {
		avifSpeed = 10
	}

	avifOutputFilename := fmt.Sprintf("%s_avif", screenshotFile)
	_, err = compression.AvifCompress(screenshotFile, avifOutputFilename, avif.Options{
		Quality: avifQuality,
		QualityAlpha: avifAlphaQuality,
		Speed: avifSpeed,
		ChromaSubsampling: image.YCbCrSubsampleRatio420,
	}, 20)
	if err != nil {	
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("failed to compress avif: %v", err),
		})
		return
	}

	azureResponse, err := command.CommandRunBarosaAzure(avifOutputFilename, features)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("barosa azure call failed: %v", err),
		})
		return
	}
	
	var azureJson map[string]interface{}
	err = json.Unmarshal([]byte(azureResponse), &azureJson)
	if err != nil {	
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("unable to parse azure response to json: %v (azure raw response: %s)", err, azureResponse),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"azureResponse": azureJson,
	})
}

func Init() {
	log.Println("Starting barosa backend")
	r := gin.Default()
	r.SetTrustedProxies(nil)
	r.Use(RequestCorsMiddleware())
	r.Use(RequestAuthorizeMiddleware())
	r.GET("/ping", RequestPing)
	r.GET("/image-features", RequestImageFeatures)
	r.Run()
}
