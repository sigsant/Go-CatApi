package facts

import (
	"fmt"
)

func init() {
	fmt.Println("\n\tINFO: Paquete 'fact' inicializado correctamente...")
}

//Facts devuelve los valores del Json deseados.
type Facts struct {
	Animal string `json:"type"`
	ID     string `json:"_id"`
	Fact   string `json:"text"`
}
