package models
var (
	SabanaList map[string]*Sabana
)
func init() {
	SabanaList = make(map[string]*Sabana)
}
type PorConcepto struct {
	Saldo float64	
	NombreCuenta string
}
type Sabana struct {
	Nombre string
	PorConcepto map[string]*PorConcepto
}
func AddSabana(u Sabana, index string)  {
	SabanaList[index] = &u	
}
func GetAllSabana() map[string]*Sabana {
	return SabanaList
}
func ClearSabana() {
	SabanaList = make(map[string]*Sabana)
}
func DameSabanaPorCodigo(index string) Sabana {
	return *SabanaList[index]
}