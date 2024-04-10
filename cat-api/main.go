package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type cat struct {
	Id       int    `json:"id"`
	Birthday string `json:"birthday"`
	Name     string `json:"name"`
	Breed    string `json:"breed"`
	IsHungry bool   `json:"isHungry"`
}

var cats = []cat{
	{Id: 1, Birthday: "12/17", Name: "Alex", Breed: "Maine Coon", IsHungry: false},
	{Id: 2, Birthday: "4/26", Name: "Andrew", Breed: "Siamese", IsHungry: true},
	{Id: 3, Birthday: "5/23", Name: "Laurie", Breed: "Scottish Fold", IsHungry: false},
}

func getAllCats(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, cats)
}

func getCatById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input for param \"id\""})
		return
	}

	for _, a := range cats {
		if a.Id == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "cat not found"})
}

func postCats(c *gin.Context) {
	var newCat cat

	if err := c.BindJSON(&newCat); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid input for body \"cat\""})
		return
	}

	for _, a := range cats {
		if a.Name == newCat.Name {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "a cat with that name already exists"})
			return
		}
	}

	cats = append(cats, newCat)
	c.IndentedJSON(http.StatusCreated, newCat)
}

func main() {
	router := gin.Default()

	router.GET("/cats", getAllCats)
	router.GET("/cat/:id", getCatById)
	router.POST("/cats", postCats)

	router.Run("localhost:3000")
}
