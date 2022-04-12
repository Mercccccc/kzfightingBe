package controller

import "github.com/gin-gonic/gin"

func Run() {
	g := gin.Default()
	g.Use(Cors())

	g.POST("/record", PostRecord)
	g.POST("/comment", PostComment)
	g.GET("/records/number", GetRecordNumber)
	g.GET("/records", GetRecordByPage)

	err := g.Run(":678")
	if err != nil {
		panic(err)
	}
	return
}
