package models
var (
	NombreCuentasList map[string]*NombreCuentas
)
func init() {
	NombreCuentasList = make(map[string]*NombreCuentas)
}

type NombreCuentas struct {
	Nombre string
	Cuenta string
	ACNT_CODE string
}
func AddNombreCuentas(u NombreCuentas, index string)  {
	NombreCuentasList[index] = &u	
}
func DameNombreCuentas(index string) NombreCuentas {
	return *NombreCuentasList[index]
}
func GetAllNombreCuentas() map[string]*NombreCuentas {
	return NombreCuentasList
}
func ClearNombreCuentas() {
	NombreCuentasList = make(map[string]*NombreCuentas)
}