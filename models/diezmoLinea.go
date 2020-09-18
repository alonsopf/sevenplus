package models
var (
	DiezmoConceptoLineaList map[string]*DiezmoConceptoLinea
)
func init() {
	DiezmoConceptoLineaList = make(map[string]*DiezmoConceptoLinea)
}
type DiezmoConceptoLinea struct {
	IdDiezmoLinea int 
	IdDiezmoConcepto int
	Orden int 
	Nombre string 
	Cuenta string 
	EsSuma string 
	Aparece int 
	D_C string
	ANAL_T0 string
	ANAL_T1 string
	ANAL_T2 string
	ANAL_T3 string
	ANAL_T4 string
	ANAL_T5 string
	ANAL_T6 string
	ANAL_T7 string
	ANAL_T8 string
	ANAL_T9 string
	DeboFacturar int 
	Rfc string
	Tipo int
	DESCRIPTN string
	TREFERENCE string
}
func ClearDiezmoConceptoLinea() {
	DiezmoConceptoLineaList = make(map[string]*DiezmoConceptoLinea)
}

func AddDiezmoConceptoLinea(u DiezmoConceptoLinea, index string)  {
	DiezmoConceptoLineaList[index] = &u	
}
func GetAllDiezmoConceptoLinea() map[string]*DiezmoConceptoLinea {
	return DiezmoConceptoLineaList
}
