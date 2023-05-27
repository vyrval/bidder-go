package main

import (
	"encoding/json"
	"log"

	"ako.com/ad"
	"github.com/gin-gonic/gin"
)

func healthHandler(c *gin.Context) {
	c.AbortWithStatus(200)
}

func getAllAds(c *gin.Context) {
	c.JSON(200, ad.GetAllAds())
}

func getAdById(c *gin.Context) {
	id := c.Param("id")
	if ad, found := ad.GetAdById(id); found {
		c.JSON(200, ad)
	} else {
		c.AbortWithStatus(204)
	}
}

func createAd(c *gin.Context) {
	request := c.Request

	var newAd ad.Ad
	err := json.NewDecoder(request.Body).Decode(&newAd)
	if err != nil {
		log.Println(err)
		c.JSON(400, err.Error())
		return
	}
	c.JSON(201, *ad.CreateAd(newAd))
}

func getAllLeaves(c *gin.Context) {
	c.JSON(200, ad.GetTree())
}
func getLeafByKey(c *gin.Context) {
	key := c.Param("key")
	if leaf, found := ad.GetLeafByKey(key); found {
		c.JSON(200, leaf)
	} else {
		c.AbortWithStatus(204)
	}
}

func main() {
	router := gin.Default()

	router.GET("/health", healthHandler)
	router.GET("/ads", getAllAds)
	router.GET("/ads/:id", getAdById)

	router.GET("/tree", getAllLeaves)
	router.GET("/tree/:key", getLeafByKey)

	router.POST("/ads", createAd)

	router.Run("localhost:8080")
}
