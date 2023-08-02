package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Allow all origins
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "OPTIONS"}
	r.Use(cors.New(corsConfig))

	r.GET("/:manifest/manifest.json", func(c *gin.Context) {

		manifestPath := "manifests/" + c.Param("manifest") + ".json"

		jsonFile, err := os.Open(manifestPath)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}

		defer jsonFile.Close()

		byteValue, _ := os.ReadFile(manifestPath)

		var manifest Manifest
		json.Unmarshal(byteValue, &manifest)

		c.JSON(http.StatusOK, manifest)
	})

	r.Run()
}
