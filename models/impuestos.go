package models
var (
	TipoImpuestoList map[string]*TipoImpuesto
)
func init() {
	TipoImpuestoList = make(map[string]*TipoImpuesto)
}
type Impuesto struct {
	Importe float64
}
type TipoImpuesto struct {
	ImpuestosList map[string]*Impuesto
}
func AddImpuesto(u TipoImpuesto, index string)  {
	TipoImpuestoList[index] = &u	
}
func DameImpuestoPorCodigo(index string) TipoImpuesto {
	return *TipoImpuestoList[index]
}
func GetAllImpuesto() map[string]*TipoImpuesto {
	return TipoImpuestoList
}
func ClearImpuesto() {
	TipoImpuestoList = make(map[string]*TipoImpuesto)
}