package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/go-playground/validator/v10"
)

type Filter struct {
	AdType string `json:"type" validate:"required,type"`
	H      int    `json:"h" validate:"required,h"`
	W      int    `json:"w" validate:"required,:w"`
}

type Ad struct {
	ID      string  `json:id`
	Name    string  `json:name`
	Filters Filter  `json:filters`
	Price   float32 `json:price`
	Key     string
}

var ads = make(map[string]Ad)
var sortedAds = make(map[string](map[string]*Ad))
var validate *validator.Validate

func healthHandler(c *gin.Context) {
	c.AbortWithStatus(200)
}

func getAllAds(c *gin.Context) {
	c.JSON(200, ads)
}

func getAdById(c *gin.Context) {
	id := c.Param("id")
	if ad, found := ads[id]; found {
		c.JSON(200, ad)
	} else {
		c.AbortWithStatus(204)
	}
}

func createAd(c *gin.Context) {
	request := c.Request

	var newAd Ad
	json.NewDecoder(request.Body).Decode(&newAd)
	if err := validate.Struct(newAd); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}
		c.JSON(400, err)
		return
	}


	if len(newAd.ID) == 0 {
		newAd.ID = uuid.NewString()
	}
	newAd.Key = fmt.Sprintf("%s:%d:%d",newAd.Filters.AdType, newAd.Filters.H, newAd.Filters.W)
	ads[newAd.ID] = newAd
	addToSortedAds(&newAd)
	c.JSON(201, newAd)
}

func addToSortedAds(ad *Ad) {
	if leaf := sortedAds[ad.Key]; leaf == nil {
		sortedAds[ad.Key] = make(map[string]*Ad)
	}
	sortedAds[ad.Key][ad.ID] = ad
	fmt.Sprintln(ad)
}

func getAllLeaves(c * gin.Context) {
	c.JSON(200, sortedAds)
}
func getLeafByKey(c * gin.Context) {
	key := c.Param("key")
	if leaf, found := sortedAds[key]; found {
		c.JSON(200, leaf)
	} else {
		c.AbortWithStatus(204)
	}
}

func wait(c * gin.Context) {
	// open output file
	fo, err := os.Create("output.txt")
	if err != nil {
			panic(err)
	}
	// close fo on exit and check for its returned error
	defer func() {
			if err := fo.Close(); err != nil {
					panic(err)
			}
	}()

	for i := 0; i < 9000; i++ {
		// write a chunk
		if _, err := fo.WriteString(fmt.Sprintf("%d", i)); err != nil {
				panic(err)
		}
	}
	c.AbortWithStatus(200)
}

func main() {
	validate = validator.New()

	router := gin.Default()

	router.GET("/health", healthHandler)
	router.GET("/ads", getAllAds)
	router.GET("/ads/:id", getAdById)

	router.GET("/tree", getAllLeaves)
	router.GET("/tree/:key", getLeafByKey)

	router.GET("/wait", wait)

	router.POST("/ads", createAd)

	router.Run("localhost:8080")
}
