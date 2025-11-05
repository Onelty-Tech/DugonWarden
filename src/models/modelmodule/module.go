package modelmodule

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Module struct {
	Name      string    `json:"Name"`
	Key       uuid.UUID `json:"Key"`
	Port      []string  `json:"Port"`
	Available bool      `json:"Available"`
}

type respUnlock struct {
	Key uuid.UUID `json:"Key"`
}

/*
Coloca candado al modulo.
*/
func (m *Module) Lock() {
	m.Available = false
}

func (m *Module) UnlockModule(ctx *gin.Context) {
	resp := respUnlock{}

	if err := ctx.BindJSON(&resp); err != nil {
		ctx.JSON(
			404,
			gin.H{
				"message": err.Error(),
			},
		)
		return
	}
	if resp.Key != m.Key {
		ctx.JSON(
			404,
			gin.H{
				"message": "invalid key.",
			},
		)
		return
	}
	ctx.JSON(
		200,
		gin.H{
			"message": "sucessfull, el modulo se a abierto.",
		},
	)
}
