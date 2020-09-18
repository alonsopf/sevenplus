package models
var (
	NotificacionesList map[string]*Notificaciones
)
func init() {
	NotificacionesList = make(map[string]*Notificaciones)
}
type Notificaciones struct {
	Timestamp int 
	Accion string
	Titulo string 
	Descripcion string
}
func ClearNotificaciones() {
	NotificacionesList = make(map[string]*Notificaciones)
}

func AddNotificaciones(u Notificaciones, index string)  {
	NotificacionesList[index] = &u	
}
func GetAllNotificaciones() map[string]*Notificaciones {
	return NotificacionesList
}
