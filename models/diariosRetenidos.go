package models
var (
	DiarioRetenidoList map[string]*DiarioRetenido
)
func init() {
	DiarioRetenidoList = make(map[string]*DiarioRetenido)
}
type DiarioRetenido struct {
	Timestamp int
	Cuantas int
	Usuario string
	TipoDiario string
}
func ClearDiarioRetenidos() {
	DiarioRetenidoList = make(map[string]*DiarioRetenido)
}

func AddDiarioRetenido(u DiarioRetenido, index string)  {
	DiarioRetenidoList[index] = &u	
}
func GetAllDiarioRetenido() map[string]*DiarioRetenido {
	return DiarioRetenidoList
}
