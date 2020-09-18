package models
var (
	NominasConceptoList map[string]*NominasConcepto
)
func init() {
	NominasConceptoList = make(map[string]*NominasConcepto)
}
type NominasConcepto struct {
	IdRegistro int
	ACNT_CODE string
	Tipo int
	TipoSAT string
	Concepto string
	GravadoExcento int
}
func ClearNominasConcepto() {
	NominasConceptoList = make(map[string]*NominasConcepto)
}

func AddNominasConcepto(u NominasConcepto, index string)  {
	NominasConceptoList[index] = &u	
}
func GetAllNominasConcepto() map[string]*NominasConcepto {
	return NominasConceptoList
}
