package models
var (
	IngresoEgresoList map[string]*IngresoEgreso
)
func init() {
	IngresoEgresoList = make(map[string]*IngresoEgreso)
}

type IngresoEgreso struct {
	Cuenta string
	Nombre string
	Ingreso float64
	Egreso float64
}
func DameIngresoEgreso(index string) IngresoEgreso {
	return *IngresoEgresoList[index]
}
func AddIngresoEgreso(u IngresoEgreso, index string)  {
	IngresoEgresoList[index] = &u	
}
func GetAllIngresoEgreso() map[string]*IngresoEgreso {
	return IngresoEgresoList
}
func ClearIngresoEgreso() {
	IngresoEgresoList = make(map[string]*IngresoEgreso)	
}