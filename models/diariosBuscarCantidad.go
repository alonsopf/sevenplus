package models
var (
	DiarioBuscarCantidadList map[string]*DiarioBuscarCantidad
)
func init() {
	DiarioBuscarCantidadList = make(map[string]*DiarioBuscarCantidad)
}
type DiarioBuscarCantidad struct {
	Diario int
	Linea int
	AMOUNT float64
	D_C string
	TRANS_DATETIME string
	PERIOD int
	DESCRIPTN string
	TREFERENCE string
}
func AddDiarioBuscarCantidad(u DiarioBuscarCantidad, index string)  {
	DiarioBuscarCantidadList[index] = &u	
}
func GetAllDiarioBuscarCantidad() map[string]*DiarioBuscarCantidad {
	return DiarioBuscarCantidadList
}
func ClearDiarioBuscarCantidad() {
	DiarioBuscarCantidadList = make(map[string]*DiarioBuscarCantidad)
}