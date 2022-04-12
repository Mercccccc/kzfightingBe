package controller

import "github.com/gin-gonic/gin"

func Run() {
	g := gin.Default()

	err := g.Run(":678")
	if err != nil {
		panic(err)
	}
	return
}
