package models
var (
	CuentaBancariaList map[string]*CuentaBancaria
)
func init() {
	CuentaBancariaList = make(map[string]*CuentaBancaria)
}
type PorUUID struct {
	FOLIO_FISCAL string
	AMOUNT float64
	PDF string
	XML string
	Ruta string
}
type CuentaBancaria struct {
	IdBanco int
	Importe float64
	Descripcion string 
	Leyenda1 string 
	Leyenda2 string
	CuentaPropia int
	Referencia string
	Fecha string
	Timbrado int 
	UUIDList map[string]*PorUUID
}
func AddCuentaBancaria(u CuentaBancaria, index string)  {
	CuentaBancariaList[index] = &u	
}
func DameCuentaBancaria(index string) CuentaBancaria {
	return *CuentaBancariaList[index]
}
func GetAllCuentaBancaria() map[string]*CuentaBancaria {
	return CuentaBancariaList
}
func ClearCuentaBancaria() {
	CuentaBancariaList = make(map[string]*CuentaBancaria)
}