package models
var (
	LineaDetalleList map[string]*LineaDetalle
)
func init() {
	LineaDetalleList = make(map[string]*LineaDetalle)
}

type LineaDetalle struct {
	Subdetalle int
	Nombre string
	AMOUNT float64
}
func DameLineaDetalle(index string) LineaDetalle {
	return *LineaDetalleList[index]
}
func AddLineaDetalle(u LineaDetalle, index string)  {
	LineaDetalleList[index] = &u	
}
func GetAllLineaDetalle() map[string]*LineaDetalle {
	return LineaDetalleList
}
func ClearLineaDetalle() {
	LineaDetalleList = make(map[string]*LineaDetalle)	
}
