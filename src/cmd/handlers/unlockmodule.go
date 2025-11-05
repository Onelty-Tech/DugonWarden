package handlers

import (
	"DugonWarden/src/models/modelrobot"

	"github.com/gin-gonic/gin"
)

func UnlockModuleHandler(robot *modelrobot.Robot) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		moduleName := ctx.Param("module")
		if module, err := robot.LookupModule(moduleName); err == nil {
			module.UnlockModule(ctx)
			return
		}
		ctx.JSON(
			409,
			gin.H{
				"message": "Accion invalida, el modulo no esta cerrado",
			},
		)
		return
	}
}
