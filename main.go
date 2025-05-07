package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type item struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Email string `json:"email"`
}

var items = []item{
	{Id: 1, Name: "mehafooz-1", Email: "mehafooz1@gmail.com"},
	{Id: 2, Name: "mehafooz-2", Email: "mehafooz2@gmail.com"},
}

func main() {
	fmt.Println("hello worldd!!!")

	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello test\n")
		ctx.JSON(http.StatusOK, gin.H{"Welcome to local hoster 8080": "ok"})

	})

	router.POST("/item", addItem)

	router.GET("/item", getItems)

	router.DELETE("/item/:id", deleteItems)

	router.Run(":8080")
}

func addItem(ctx *gin.Context) {
	var itm item

	err := ctx.ShouldBindBodyWithJSON(&itm)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Error while parsing the item!")
		return
	}

	items = append(items, itm)

	ctx.JSON(http.StatusAccepted, gin.H{"data": items})

}

func getItems(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"All the itemss are ": items})
}

func deleteItems(ctx *gin.Context) {
	parm := ctx.Param("id")

	parmId, err := strconv.Atoi(parm)

	if err != nil {
		ctx.String(http.StatusBadRequest, "Error while parsing the id!")
		return
	}

	for i, itm := range items {
		if itm.Id == parmId {
			items = append(items[:i], items[i+1:]...)
			ctx.JSON(http.StatusAccepted, gin.H{"new items after deleting": items})
		}
	}
}
