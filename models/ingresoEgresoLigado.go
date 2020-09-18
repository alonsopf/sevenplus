package models
var (
	IngresoEgresoLigadoList map[string]*IngresoEgresoLigado
)
func init() {
	IngresoEgresoLigadoList = make(map[string]*IngresoEgresoLigado)
}

type IngresoEgresoLigado struct {
	Cuenta string
	Nombre string
	Ingreso float64
	Egreso float64
	LigadoIngreso float64
	LigadoEgreso float64
}
func DameIngresoEgresoLigado(index string) IngresoEgresoLigado {
	return *IngresoEgresoLigadoList[index]
}
func AddIngresoEgresoLigado(u IngresoEgresoLigado, index string)  {
	IngresoEgresoLigadoList[index] = &u	
}
func GetAllIngresoEgresoLigado() map[string]*IngresoEgresoLigado {
	return IngresoEgresoLigadoList
}
func ClearIngresoEgresoLigado() {
	IngresoEgresoLigadoList = make(map[string]*IngresoEgresoLigado)	
}