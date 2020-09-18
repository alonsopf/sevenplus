package models
var (
	BalanzaList map[string]*Balanza
)
func init() {
	BalanzaList = make(map[string]*Balanza)
}
type Balanza struct {
	Inicial string
	Creditos string
	Debitos string
	Final string
}
func AddBalanza(u Balanza, index string)  {
	BalanzaList[index] = &u	
}
func GetAllBalanza() map[string]*Balanza {
	return BalanzaList
}
func ClearBalanza() {
	BalanzaList = make(map[string]*Balanza)
}