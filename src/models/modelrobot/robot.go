package modelrobot

import (
	"DugonWarden/src/models/modelmodule"
	"fmt"

	"github.com/google/uuid"
)

type Robot struct {
	LookupPort  map[string]*modelmodule.Module
	Unavailable []string
}

/*
funcion que agrega el modulo al mapa del robot.
*/
func (r *Robot) AddModule(name string, ports []string) {
	for _, port := range ports {
		for _, usedport := range r.Unavailable {
			if port == usedport {
				return
			}
		}
	}
	r.LookupPort[name] = &modelmodule.Module{
		Name:      name,
		Key:       uuid.New(),
		Port:      ports,
		Available: false, //ocupado por defecto siempre
	}
	r.Unavailable = append(r.Unavailable, ports...)
}

/*
Echa un vistazo los datos de los modulos y devuelve si existe.
*/
func (r *Robot) LookupModule(name string) (*modelmodule.Module, error) {
	if module, exists := r.LookupPort[name]; exists {
		return module, nil
	}
	return &modelmodule.Module{}, fmt.Errorf("el modulo no existe en el robot.")
}
