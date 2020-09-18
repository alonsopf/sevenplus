package models
var (
	ConciliacionDiariosList map[string]*ConciliacionDiarios
)
func init() {
	ConciliacionDiariosList = make(map[string]*ConciliacionDiarios)
}

type ConciliacionDiarios struct {
	Diario int
	Linea int 
	Cantidad float64
	Fecha string
	Referencia string
	DescripcionConta string
	Periodo int
	DiarioC int
	LineaC int 
	PeriodoC int
}
func DameConciliacionDiarios(index string) ConciliacionDiarios {
	return *ConciliacionDiariosList[index]
}
func AddConciliacionDiarios(u ConciliacionDiarios, index string)  {
	ConciliacionDiariosList[index] = &u	
}
func GetAllConciliacionDiarios() map[string]*ConciliacionDiarios {
	return ConciliacionDiariosList
}
func ClearConciliacionDiarios() {
	ConciliacionDiariosList = make(map[string]*ConciliacionDiarios)	
}