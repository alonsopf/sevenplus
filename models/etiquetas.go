package models
var (
	EtiquetasList map[string]*Etiquetas
)
func init() {
	EtiquetasList = make(map[string]*Etiquetas)
}
type Palabras struct {
	IdLigadoRFC string
	Palabra string
}
type Etiquetas struct {
	IdEtiqueta string
	Etiqueta string
	PalabrasList  map[string]*Palabras
}
func DameEtiquetas(index string) Etiquetas {
	return *EtiquetasList[index]
}
func AddEtiquetas(u Etiquetas, index string)  {
	EtiquetasList[index] = &u	
}
func GetAllEtiquetas() map[string]*Etiquetas {
	return EtiquetasList
}
func ClearEtiquetas() {
	EtiquetasList = make(map[string]*Etiquetas)	
}