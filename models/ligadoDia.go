package models
var (
	LigadoDiaList map[string]*LigadoDia
)
func init() {
	LigadoDiaList = make(map[string]*LigadoDia)
}

type Ligado struct {
	FOLIO_FISCAL string
	AMOUNT float64
	PDF string
	XML string
	Ruta string
	IdEtiqueta int
}
type LigadoDia struct {
	Nombre string
	IdBanco int
	Importe float64
	Descripcion string
	LigadoList map[string]*Ligado
}
func AddLigadoDia(u LigadoDia, index string)  {
	LigadoDiaList[index] = &u	
}
func DameLigadoDiaPorCodigo(index string) LigadoDia {
	return *LigadoDiaList[index]
}
func GetAllLigadoDia() map[string]*LigadoDia {
	return LigadoDiaList
}
func ClearLigadoDia() {
	LigadoDiaList = make(map[string]*LigadoDia)
}