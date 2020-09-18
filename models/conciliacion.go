package models
var (
	ConciliacionList map[string]*Conciliacion
)
func init() {
	ConciliacionList = make(map[string]*Conciliacion)
}

type Conciliacion struct {
	Diario int
	Linea int 
	Cantidad float64
	D_C string
	Fecha string
	DescripcionConta string
	DescripcionBanco string
	IdBanco int
}
func DameConciliacion(index string) Conciliacion {
	return *ConciliacionList[index]
}
func AddConciliacion(u Conciliacion, index string)  {
	ConciliacionList[index] = &u	
}
func GetAllConciliacion() map[string]*Conciliacion {
	return ConciliacionList
}
func ClearConciliacion() {
	ConciliacionList = make(map[string]*Conciliacion)	
}