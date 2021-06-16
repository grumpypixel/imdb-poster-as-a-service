package main

import (
	"net/http"

	"github.com/buger/jsonparser"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/grumpypixel/imdb-poster-go"
	"github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()
	router.Use(static.Serve("/", static.LocalFile("./assets", false)))
	router.POST("/poster", posterHandler)
	err := router.Run()
	if err != nil {
		logrus.Fatal(err)
	}
}

func posterHandler(c *gin.Context) {
	raw, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	movie, err := jsonparser.GetString(raw, "movie")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	all, err := jsonparser.GetBoolean(raw, "all")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	imdb := imdb.NewIMDB(all)
	posters := imdb.FetchPoster(movie)

	c.JSON(http.StatusOK, gin.H{
		"posters": posters,
	})
}
