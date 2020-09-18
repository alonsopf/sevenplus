package models
var (
	CandidatosTimbradoList map[string]*CandidatosTimbrado
)
func init() {
	CandidatosTimbradoList = make(map[string]*CandidatosTimbrado)
}
type CandidatosTimbrado struct {
	JRNAL_NO int 
	JRNAL_LINE int 
	AmountPrima float64
	RFC string 
	RazonSocial string 
	Correo string 
	NumCtaPago string
	MetodoDePago string 
	Nombre string
	PERIOD int
	Fecha string
	DESCRIPTN string
}

func AddCandidatosTimbrado(u CandidatosTimbrado, index string)  {
	CandidatosTimbradoList[index] = &u	
}
func GetAllCandidatosTimbrado() map[string]*CandidatosTimbrado {
	return CandidatosTimbradoList
}
func ClearCandidatosTimbrado() {
	CandidatosTimbradoList = make(map[string]*CandidatosTimbrado)	
}
