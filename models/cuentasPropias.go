package models
var (
	CuentasPropiasTransferenciaList map[string]*CuentasPropiasTransferencia
)
func init() {
	CuentasPropiasTransferenciaList = make(map[string]*CuentasPropiasTransferencia)
}

type CuentasPropiasTransferencia struct {
	Nombre string
	IdBanco int
	Importe float64
	Fecha string
	Descripcion string
	STATUS string
	Referencia string
	Saldo string
	Leyenda1 string
	Leyenda2 string
	Concentrado string
}
func AddCuentasPropiasTransferencia(u CuentasPropiasTransferencia, index string)  {
	CuentasPropiasTransferenciaList[index] = &u	
}
func DameCuentasPropiasTransferencia(index string) CuentasPropiasTransferencia {
	return *CuentasPropiasTransferenciaList[index]
}
func GetAllCuentasPropiasTransferencia() map[string]*CuentasPropiasTransferencia {
	return CuentasPropiasTransferenciaList
}
func ClearCuentasPropiasTransferencia() {
	CuentasPropiasTransferenciaList = make(map[string]*CuentasPropiasTransferencia)
}