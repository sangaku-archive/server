package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/:manifest/manifest.json", func(c *gin.Context) {

		manifestPath := "manifests/" + c.Param("manifest") + ".json"

		jsonFile, err := os.Open(manifestPath)
		if err != nil {
			c.JSON(http.StatusNotFound, err)
		}

		fmt.Println("Successfully Opened users.json")
		defer jsonFile.Close()

		byteValue, _ := os.ReadFile(manifestPath)

		var manifest Manifest

		json.Unmarshal(byteValue, &manifest)

		c.JSON(http.StatusOK, manifest)
	})

	r.Run()
}
