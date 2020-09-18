package models
var (
	TimbradoDiaList map[string]*TimbradoDia
)
func init() {
	TimbradoDiaList = make(map[string]*TimbradoDia)
}


type TimbradoDia struct {
	Nombre string
	IdBanco int
	Importe float64
	Descripcion string
	Timbrado int
	UUID string
}
func AddTimbradoDia(u TimbradoDia, index string)  {
	TimbradoDiaList[index] = &u	
}
func DameTimbradoDiaPorCodigo(index string) TimbradoDia {
	return *TimbradoDiaList[index]
}
func GetAllTimbradoDia() map[string]*TimbradoDia {
	return TimbradoDiaList
}
func ClearTimbradoDia() {
	TimbradoDiaList = make(map[string]*TimbradoDia)
}