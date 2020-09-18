package models
var (
	ClientesList map[string]*Clientes
)
func init() {
	ClientesList = make(map[string]*Clientes)
}
type NumCtaPago struct {
	IdNumCtaPago int
	NumCtaPago string
}
type Clientes struct {
	Rfc string
	RazonSocial string
	Correo string
	IdCliente int
	NumCtaPagoList  map[string]*NumCtaPago
}
func DameClientePorIdCLiente(index string) Clientes {
	return *ClientesList[index]
}
func AddClientes(u Clientes, index string)  {
	ClientesList[index] = &u	
}
func GetAllClientes() map[string]*Clientes {
	return ClientesList
}
func ClearClientes() {
	ClientesList = make(map[string]*Clientes)	
}
