package handlers

import (
	"DugonWarden/src/models/modelmodule"
	"DugonWarden/src/models/modelrobot"
	"fmt"

	"github.com/gin-gonic/gin"
)

type respbody struct {
	Pins []string `json:"Pins"`
}

func LockHandler(robot *modelrobot.Robot) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var resp = respbody{}
		moduleName := ctx.Param("module")
		if module, err := robot.LookupModule(moduleName); err == nil {
			if err := tryUnLock(module); err != nil {
				fmt.Println("esta lockeado al hacer un tryunlock")
				ctx.JSON(
					409,
					gin.H{
						"message": err.Error(),
					},
				)
				return
			}
			module.Lock()
			ctx.JSON(
				200,
				gin.H{
					"message": "el modulo existe y estaba libre.",
				},
			)
			return
		}
		if err := ctx.BindJSON(&resp); err != nil {
			ctx.JSON(
				404,
				gin.H{
					"message": err.Error(),
				},
			)
			return
		}
		robot.AddModule(moduleName, resp.Pins) //se agrega el modulo(por defecto al crearlo ya marca la flag "Available")
		ctx.JSON(
			200,
			gin.H{
				"message": "No existe tal modulo, se creo uno nuevo.",
			},
		)
	}
}

func tryUnLock(module *modelmodule.Module) error {
	if module.Available == false {
		return fmt.Errorf("warning: The module is currently in use.")
	}
	return nil
}
