package facts

// func init() {
// 	fmt.Println("\n\tDEBUG: Paquete 'fact' inicializado correctamente...")
// }

//Facts devuelve los valores del Json deseados.
type Facts struct {
	Animal string `json:"type"`
	ID     string `json:"_id"`
	Fact   string `json:"text"`
}
