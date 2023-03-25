package p1

import (
	"github.com/gin-gonic/gin"
)

// Helloworld godoc
//
//	@Summary		send hello world
//	@Description	send hello world
//	@Success		200
//	@Router			/p1/hello [get]
func Helloworld(g *gin.Context) {
	g.JSON(200, "hello world")
}

// ping godoc
//
//	@Summary		send status 200
//	@Description	check server is running
//	@Success		200
//	@Router			/p1/ping [get]
func PingExample(g *gin.Context) {
	g.Status(200)
}
