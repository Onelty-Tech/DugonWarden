package main

import (
	"DugonWarden/src/cmd/handlers"
	initsrv "DugonWarden/src/cmd/init"

	"github.com/gin-gonic/gin"
)

func main() {
	robot := initsrv.InitRobot()
	r := gin.Default()
	r.GET("warden/lock/:module", handlers.LockHandler(robot))
	r.POST("warden/unlock/:module", handlers.UnlockModuleHandler(robot))
	r.GET("warden/show", func(ctx *gin.Context) {
		ctx.JSON(
			200,
			gin.H{
				"message": robot.LookupPort,
			},
		)
	})
	r.Run(":8080")
}
