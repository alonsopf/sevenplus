package models
var (
	CedulaCuentasBancosList map[string]*CedulaCuentasBancos
	CuentasPropiasList map[string]*CuentasPropias
)
func init() {
	CedulaCuentasBancosList = make(map[string]*CedulaCuentasBancos)
	CuentasPropiasList = make(map[string]*CuentasPropias)
}
type PorPeriodoBancos struct {
	SaldoInicial float64
	Depositos float64
	Retiros float64
	SaldoFinal float64
}
type CedulaCuentasBancos struct {
	Nombre string
	Codigo string
	PeriodosList map[string]*PorPeriodoBancos
}

type CuentasPropias struct {
	Fecha string
	SalidaCuenta string
	SalidaImporte float64
	EntradaCuenta string
	EntradaImporte float64
	Observaciones string
}
func AddCuentasPropias(u CuentasPropias, index string)  {
	CuentasPropiasList[index] = &u	
}
func GetAllCuentasPropias() map[string]*CuentasPropias {
	return CuentasPropiasList
}
func ClearCuentasPropias() {
	CuentasPropiasList = make(map[string]*CuentasPropias)
}





func AddCedulaCuentasBancos(u CedulaCuentasBancos, index string)  {
	CedulaCuentasBancosList[index] = &u	
}
func DameCedulaCuentasBancosPorCodigo(index string) CedulaCuentasBancos {
	return *CedulaCuentasBancosList[index]
}
func GetAllCedulaCuentasBancos() map[string]*CedulaCuentasBancos {
	return CedulaCuentasBancosList
}
func ClearCedulaCuentasBancos() {
	CedulaCuentasBancosList = make(map[string]*CedulaCuentasBancos)
}