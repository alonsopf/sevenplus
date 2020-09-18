package models
var (
	DiferenciasList map[string]*Diferencias
)
func init() {
	DiferenciasList = make(map[string]*Diferencias)
}
type Diferencias struct {
	Diario int
	Linea int
	AMOUNT float64
	OTH_AMOUNT float64
	TRANS_DATETIME string
	PERIOD int
	DESCRIPTN string
	TREFERENCE string
	ACNT_CODE string
}
func AddDiferencia(u Diferencias, index string)  {
	DiferenciasList[index] = &u	
}
func GetAllDiferencias() map[string]*Diferencias {
	return DiferenciasList
}
func ClearDiferencia() {
	DiferenciasList = make(map[string]*Diferencias)
}