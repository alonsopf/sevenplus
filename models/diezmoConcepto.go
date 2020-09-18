package models
var (
	DiezmoConceptoList map[string]*DiezmoConcepto
)
func init() {
	DiezmoConceptoList = make(map[string]*DiezmoConcepto)
}
type DiezmoConcepto struct {
	Concepto string
	IdDiezmoConcepto int
	Activo int
}
func ClearDiezmoConcepto() {
	DiezmoConceptoList = make(map[string]*DiezmoConcepto)
}

func AddDiezmoConcepto(u DiezmoConcepto, index string)  {
	DiezmoConceptoList[index] = &u	
}
func GetAllDiezmoConcepto() map[string]*DiezmoConcepto {
	return DiezmoConceptoList
}
