package initsrv

import (
	"DugonWarden/src/models/modelmodule"
	"DugonWarden/src/models/modelrobot"
)

/*
Inicializa todos los parametros del robot.
*/
func InitRobot() *modelrobot.Robot {
	return &modelrobot.Robot{
		LookupPort:  make(map[string]*modelmodule.Module),
		Unavailable: make([]string, 0),
	}
}
