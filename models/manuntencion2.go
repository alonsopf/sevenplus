package models
var (
	Manuntencion2List map[string]*Manuntencion2
)
func init() {
	Manuntencion2List = make(map[string]*Manuntencion2)
}

type Manuntencion2 struct {
	Nombre string
	Total float64
	ER string
	Periodo string
	TimeStamp string
	Timbrado int
	UUID string
}
func AddManuntencion2(u Manuntencion2, index string)  {
	Manuntencion2List[index] = &u	
}
func DameManuntencion2PorCodigo(index string) Manuntencion2 {
	return *Manuntencion2List[index]
}
func GetAllManuntencion2() map[string]*Manuntencion2 {
	return Manuntencion2List
}
func ClearManuntencion2() {
	Manuntencion2List = make(map[string]*Manuntencion2)
}