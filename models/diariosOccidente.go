package models
var (
	DiarioOccidenteList map[string]*DiarioOccidente
)
func init() {
	DiarioOccidenteList = make(map[string]*DiarioOccidente)
}
type DiarioOccidente struct {
	Diario int
	Linea int
	PERIOD int
	DESCRIPTN string
	ACNT_CODE string
	ANAL_T9 string
	ANAL_T99 string
}
func AddDiarioOccidente(u DiarioOccidente, index string)  {
	DiarioOccidenteList[index] = &u	
}
func GetAllDiarioOccidente() map[string]*DiarioOccidente {
	return DiarioOccidenteList
}
func ClearDiarioOccidente() {
	DiarioOccidenteList = make(map[string]*DiarioOccidente)
}