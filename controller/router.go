package controller

import "github.com/gin-gonic/gin"

func Run() {
	g := gin.Default()

	g.POST("/record", PostRecord)

	err := g.Run(":678")
	if err != nil {
		panic(err)
	}
	return
}
