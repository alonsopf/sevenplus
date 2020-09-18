package models
var (
	DistritoDList map[string]*DistritoD
)
func init() {
	DistritoDList = make(map[string]*DistritoD)
}
type PorCuentaD struct {
	DESCR string
	PeriodoList map[string]*PorPeriodoD
}
type PorPeriodoD struct {
	Saldo float64
}
type PorIglesiaD struct {
	Nombre string
	CuentaList map[string]*PorCuentaD
}
type DistritoD struct {
	Nombre string
	IglesiaList map[string]*PorIglesiaD
}
func AddDistritoD(u DistritoD, index string)  {
	DistritoDList[index] = &u	
}
func DameDistritoDPorCodigo(index string) DistritoD {
	return *DistritoDList[index]
}
func GetAllDistritoD() map[string]*DistritoD {
	return DistritoDList
}
func ClearistritoD() {
	DistritoDList = make(map[string]*DistritoD)
}