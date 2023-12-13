package main

import (
	"be/model"
	"be/router"
	"be/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.New()
	model.Init()
	fmt.Println(utils.GenerateHash("301029"))
	router.Router(e)
	e.Run(":443")
}
