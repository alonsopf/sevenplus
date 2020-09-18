package models
var (
	VerLigadoList map[string]*VerLigado
)
func init() {
	VerLigadoList = make(map[string]*VerLigado)
}

type VerLigado struct {
	NombreCuenta string
	FechaFactura string
	Rfc string
	Razon string
	TotalFactura float64
	Ligado float64
	ImporteBanco float64
	BUNIT string
	FechaBanco string
	DescripcionBanco string
}
func AddVerLigado(u VerLigado, index string)  {
	VerLigadoList[index] = &u	
}
func DameVerLigado(index string) VerLigado {
	return *VerLigadoList[index]
}
func GetAllVerLigado() map[string]*VerLigado {
	return VerLigadoList
}
func ClearVerLigado() {
	VerLigadoList = make(map[string]*VerLigado)
}