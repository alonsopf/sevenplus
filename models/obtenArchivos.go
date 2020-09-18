package models
var (
	ArchivosList map[string]*Archivos
)
func init() {
	ArchivosList = make(map[string]*Archivos)
}

type Archivos struct {
	Nombre string
	Completo string
}
func AddArchivos(u Archivos, index string)  {
	ArchivosList[index] = &u	
}
func DameArchivos(index string) Archivos {
	return *ArchivosList[index]
}
func GetAllArchivos() map[string]*Archivos {
	return ArchivosList
}
func ClearArchivos() {
	ArchivosList = make(map[string]*Archivos)
}