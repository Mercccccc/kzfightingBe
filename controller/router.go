package controller

import "github.com/gin-gonic/gin"

func Run() {
	g := gin.Default()

	g.POST("/record", PostRecord)
	g.POST("/comment", PostComment)

	err := g.Run(":678")
	if err != nil {
		panic(err)
	}
	return
}
