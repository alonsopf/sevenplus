package models
var (
	ManuntencionList map[string]*Manuntencion
)
func init() {
	ManuntencionList = make(map[string]*Manuntencion)
}

type PorConceptos struct {
	SaldoGravado float64
	SaldoExento float64
	Tipo int
	TipoSAT string
	Concepto string
	GravadoExcento int
	IdRegistro int
}
type Manuntencion struct {
	ER string
	Correo string
	Nombre string
	Rfc string
	Curp string
	NumEmpleado string
	TipoContrato string
	TipoJornada string
	NumSeguridadSocial string
	FechaInicioRelLaboral string
	RiesgoPuesto string
	PeriodicidadPago string
	Departamento string
	TipoRegimen string
	ConceptosList map[string]*PorConceptos
}
func AddManuntencion(u Manuntencion, index string)  {
	ManuntencionList[index] = &u	
}
func DameManuntencionPorCodigo(index string) Manuntencion {
	return *ManuntencionList[index]
}
func GetAllManuntencion() map[string]*Manuntencion {
	return ManuntencionList
}
func ClearManuntencion() {
	ManuntencionList = make(map[string]*Manuntencion)
}