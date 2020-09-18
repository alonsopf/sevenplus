package models
var (
	UsuariosIglesiasList map[string]*UsuariosIglesias
)
func init() {
	UsuariosIglesiasList = make(map[string]*UsuariosIglesias)
}
type UsuariosIglesias struct {
	Usuario string
	TipoDeUsuario int
	IdUsuario int
	Dimension string
	Nombre string
	Iglesia string
	Activo int
}
func ClearUsuariosIglesias() {
	UsuariosIglesiasList = make(map[string]*UsuariosIglesias)
}

func AddUsuariosIglesias(u UsuariosIglesias, index string)  {
	UsuariosIglesiasList[index] = &u	
}
func GetAllUsuariosIglesias() map[string]*UsuariosIglesias {
	return UsuariosIglesiasList
}
